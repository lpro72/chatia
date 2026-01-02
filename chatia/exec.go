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
* execFile
*******************/
func execFile(filePath string) int {
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

	// Read the file to execute
	file, err := os.Open(filePath)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_FILE_OPEN, filePath, err.Error())
		return errcode.ERROR_FATAL_FILE_OPEN
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)
	connReader := bufio.NewReader(conn)
	for {
		// Read a line from the file
		line, err := fileReader.ReadString('\n')
		if err != nil {
			// End of file
			if err.Error() == "EOF" {
				break
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_FILE_READ, filePath, err.Error())
			return errcode.ERROR_FATAL_FILE_READ
		}
		command := strings.TrimSpace(line)
		if command == "" || strings.HasPrefix(command, "#") {
			continue
		}

		fmt.Printf(">> %s\n", command)

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
