package main

/*******************
* Import
*******************/
import (
    "os"
    "net"
    "fmt"
    "bufio"
    "chatia/modules/error"
)

/*******************
* ExecShell
*******************/
func ExecShell(context *ContextStruct) {
    received := make([]byte, 100);

    tcpServer, err := net.ResolveTCPAddr(TYPE, HOST + ":" + PORT)
    if err != nil {
        error.PrintMsgFromErrorCode(error.ERROR_FATAL_CLIENT_INVALID_DATA, HOST + ":" + PORT, err.Error())
    }
    
    conn, err := net.DialTCP(TYPE, nil, tcpServer)
    if err != nil {
        error.PrintMsgFromErrorCode(error.ERROR_FATAL_CLIENT_NOT_CONNECT, err.Error())
    }

    readerStdin := bufio.NewReader(os.Stdin)

    for true {
        fmt.Print("Enter command: ")
        command, _ := readerStdin.ReadString('\n')
        
        _, err = conn.Write([]byte(command))
        if err != nil {
            error.PrintMsgFromErrorCode(error.ERROR_CLIENT_WRITE, err.Error())
        }

    
        _, err = conn.Read(received)
        if err != nil {
            error.PrintMsgFromErrorCode(error.ERROR_CLIENT_READ, err.Error())
        }
    
        fmt.Println(string(received))
        
        if command == "exit\n" {
            break
        }
    }
    
    conn.Close()

}


