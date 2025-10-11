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

	"chatia/modules/brain"
	"chatia/modules/errcode"
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
* ExecAsDaemon
*******************/
func ExecAsDaemon() int {
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
	var brainContext brain.I_Brain = nil
	var brainNameUsed string = "None"
	var returnMsg string
	fmt.Println("Client connected")

	for {
		returnMsg = "Command not found"

		received, err := connReader.ReadString('\n')
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_SERVER_READ, err.Error())
			break
		}

		command := strings.TrimSpace(received)
		fmt.Println(command)

		if command == "stop" {
			setDaemonState(DAEMON_STOP)
			setDaemonListener(nil)
			break
		} else if command == "exit" {
			break
		} else if newBrainNameUsed, ok := strings.CutPrefix(command, "use "); ok {
			switch newBrainNameUsed {
			case "new brain":
				count := atomic.LoadInt32(&activeConnections)
				returnMsg = fmt.Sprintf("Cannot create a new brain while there are %d active connections", count)
				if count < 2 {
					brain.UseTemporaryBrain()
					returnMsg = "New brain created and used"
					brainNameUsed = ""
					brainContext = nil
				}
			case "main brain":
				brain.UseMainBrain()
				returnMsg = "Main brain used"
				brainNameUsed = ""
				brainContext = nil
			default:
				newBrainContext := brain.GetBrainContext(newBrainNameUsed)
				if newBrainContext == nil {
					returnMsg = "Invalid brain name"
				} else {
					brainNameUsed = newBrainNameUsed
					brainContext = newBrainContext
					returnMsg = "Done"
				}
			}
		} else if brainContext == nil {
			returnMsg = "You need to use a Brain"
		} else if strings.HasPrefix(command, "unittest") {
			brainContext.CallUnittestFunction()
			returnMsg = "Done"
		} else if strings.HasPrefix(command, "dumpmemory") {
			brainContext.CallDumpMemoryFunction()
			returnMsg = "Done"
		}

		_, err = fmt.Fprintf(conn, "(%s) : %s\n", brainNameUsed, returnMsg)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_SERVER_WRITE, err.Error())
			break
		}
	}

	fmt.Println("Client disconnected")

}
