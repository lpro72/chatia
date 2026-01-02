package main

/*******************
* Import
*******************/
import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"sync/atomic"

	"chatia/modules/data"
	"chatia/modules/errcode"
	"chatia/modules/interfaces"
)

/*******************
* Constante
*******************/
const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"

	DAEMON_STOP  = 0
	DAEMON_START = 1
)

/*******************
* Globals variables
*******************/
var daemonState int32 = DAEMON_STOP
var daemonListener net.Listener
var daemonListenerMutex sync.Mutex
var connectionList sync.WaitGroup
var activeConnections int32 = 0

/*******************
* setDaemonState
*******************/
func setDaemonState(state int) {
	atomic.StoreInt32(&daemonState, int32(state))
}

/*******************
* isDaemonStarted
*******************/
func isDaemonStarted() bool {
	return atomic.LoadInt32(&daemonState) == DAEMON_START
}

/*******************
* setDaemonListener
*******************/
func setDaemonListener(listener net.Listener) {
	daemonListenerMutex.Lock()
	defer daemonListenerMutex.Unlock()
	if listener == nil && daemonListener != nil {
		daemonListener.Close()
	}
	daemonListener = listener
}

/*******************
* getDaemonListener
*******************/
func getDaemonListener() net.Listener {
	daemonListenerMutex.Lock()
	defer daemonListenerMutex.Unlock()
	return daemonListener
}

/*******************
* execAsDaemon
*******************/
func execAsDaemon() int {
	l, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_SERVER_NO_NETWORK, HOST+":"+PORT, err.Error())
		return errcode.ERROR_FATAL_SERVER_NO_NETWORK
	}
	setDaemonListener(l)

	fmt.Println("Server has started on PORT " + PORT)

	setDaemonState(DAEMON_START)
	for isDaemonStarted() {
		l := getDaemonListener()
		if l == nil {
			continue
		}

		conn, err := l.Accept()
		if err != nil {
			if isDaemonStarted() {
				errcode.PrintMsgFromErrorCode(errcode.WARNING_SERVER_NOT_CONNECT, err.Error())
			}
			continue
		}
		connectionList.Add(1)
		go func(c net.Conn) {
			defer connectionList.Done()
			defer atomic.AddInt32(&activeConnections, -1)
			atomic.AddInt32(&activeConnections, 1)
			handleConnection(c)
		}(conn)
	}

	connectionList.Wait()
	return errcode.SUCCESS
}

/*******************
* handleConnection
*******************/
func handleConnection(conn net.Conn) {
	if conn == nil {
		return
	}
	defer conn.Close()

	connReader := bufio.NewReader(conn)
	var brainConfig interfaces.I_BrainConfig = data.UseMainBrain()
	var brainContext interfaces.I_BrainContext = nil
	var brainNameUsed string = "None"
	var returnMsg string
	fmt.Println("Client connected")

	waitForConnection := true
	for waitForConnection {
		returnMsg = "Command not found"

		received, err := connReader.ReadString('\n')
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_SERVER_READ, err.Error())
			break
		}

		command := strings.TrimSpace(received)
		parts := strings.SplitN(command, " ", 2)
		cmd := parts[0]
		args := ""
		if len(parts) >= 2 {
			args = parts[1]
		}

		switch cmd {
		case "stop":
			setDaemonState(DAEMON_STOP)
			setDaemonListener(nil)
			waitForConnection = false
			returnMsg = "Server is stopping"
		case "exit":
			waitForConnection = false
			returnMsg = "Goodbye"
		case "dumpmemory":
			if brainContext == nil {
				returnMsg = "You need to use a Brain"
			} else {
				brainContext.CallDumpMemoryFunction()
				returnMsg = "Done"
			}
		case "learn":
			if brainContext == nil {
				returnMsg = "You need to use a Brain"
			} else {
				brainContext.CallLearnFunction([]byte(args))
			}
			returnMsg = "Done"
		case "exec":
			if brainContext == nil {
				returnMsg = "You need to use a Brain"
			} else {
				brainContext.CallExecFunction(args)
			}
			returnMsg = "Done"
		case "use":
			newBrainNameUsed := strings.TrimSpace(args)
			switch newBrainNameUsed {
			case "new brain":
				count := atomic.LoadInt32(&activeConnections)
				returnMsg = fmt.Sprintf("Cannot create a new brain while there are %d active connections", count)
				if count < 2 {
					brainConfig = data.UseTemporaryBrain()
					returnMsg = "New brain create and ready to be use"
					brainNameUsed = ""
					brainContext = nil
				}
			case "main brain":
				brainConfig = data.UseMainBrain()
				returnMsg = "Main brain used"
				brainNameUsed = ""
				brainContext = nil
			default:
				newBrainContext := brainConfig.GetBrainContextManagement().GetBrainContext(newBrainNameUsed)
				if newBrainContext == nil {
					returnMsg = fmt.Sprintf("%s is an invalid brain name", newBrainNameUsed)
				} else {
					brainNameUsed = newBrainNameUsed
					brainContext = newBrainContext
					returnMsg = "Done"
				}
			}
		default:
			if brainContext == nil {
				returnMsg = "You need to use a Brain"
			}
		}

		_, err = fmt.Fprintf(conn, "(%s) : %s\n", brainNameUsed, returnMsg)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_SERVER_WRITE, err.Error())
			break
		}
	}

	fmt.Println("Client disconnected")
}
