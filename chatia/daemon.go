package main

/*******************
* Import
*******************/
import (
    "net"
    "fmt"
    "strings"
    "chatia/modules/error"
)

/*******************
* Constante
*******************/
const (
    HOST = "localhost"
    PORT = "9001"
    TYPE = "tcp"
)

/*******************
* ExecAsDaemon
*******************/
func ExecAsDaemon(context *ContextStruct) {
    listen, err := net.Listen(TYPE, HOST + ":" + PORT)
    if err != nil {
        error.PrintMsgFromErrorCode(error.ERROR_FATAL_SERVER_NO_NETWORK, HOST + ":" + PORT, err.Error())
        return
    }
    defer listen.Close()

    fmt.Println("Server has started on PORT " + PORT)

    for {
        conn, err := listen.Accept()
        if err != nil {
            error.PrintMsgFromErrorCode(error.WARNING_SERVER_NOT_CONNECT, err.Error())
        }
        go handleConnection(conn)
    }

}

/*******************
* handleConnection
*******************/
func handleConnection(conn net.Conn) {
    received := make([]byte, 256);

    println("Client connected")

    for true {
        _, err := conn.Read(received)
        command := strings.TrimRight(string(received), "\x00")
        if err != nil {
            error.PrintMsgFromErrorCode(error.ERROR_SERVER_READ, err.Error())
            break
        }
    
        fmt.Println(command)
    
        _, err = conn.Write([]byte("test from server"))
        if err != nil {
            error.PrintMsgFromErrorCode(error.ERROR_SERVER_WRITE, err.Error())
            break
        }
        if command == "exit\n" {
            break
        }
    }
    
    conn.Close()
    
    println("Client disconnected")

}
