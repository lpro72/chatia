package main

/*******************
* Import
*******************/
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"chatia/modules/errcode"
)

/*******************
* ExecShell
*******************/
func ExecShell() int {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CLIENT_INVALID_DATA, HOST+":"+PORT, err.Error())
		return errcode.ERROR_FATAL_CLIENT_INVALID_DATA
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CLIENT_NOT_CONNECT, err.Error())
		return errcode.ERROR_FATAL_CLIENT_NOT_CONNECT
	}
	defer conn.Close()

	readerStdin := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	for {
		fmt.Print("Enter command: ")
		command, err := readerStdin.ReadString('\n')
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_CLIENT_READ, err.Error())
			return errcode.ERROR_CLIENT_READ
		}
		command = strings.TrimSpace(command)

		_, err = conn.Write([]byte(command + "\n"))
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_CLIENT_WRITE, err.Error())
			return errcode.ERROR_CLIENT_WRITE
		}

		if command == "exit" || command == "stop" {
			break
		}

		response, err := connReader.ReadString('\n')
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_CLIENT_READ, err.Error())
			return errcode.ERROR_CLIENT_READ
		}

		fmt.Println(strings.TrimRight(response, "\n"))
	}

	return errcode.SUCCESS
}
