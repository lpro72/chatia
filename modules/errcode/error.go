package errcode

/*******************
* Import
*******************/
import (
	"fmt"
	"runtime"
	"runtime/debug"
)

/*******************
* ANSI escape code color
*******************/
const RED = "\033[33m"
const YELLOW = "\033[33m"
const BLUE = "\033[34m"
const RESET_ALL = "\033[0m"

/*******************
* Error code
*******************/
// Severity
const ERROR_FATAL = 500000
const ERROR_CRITICAL int = 400000
const ERROR int = 300000
const WARNING int = 200000
const INFO int = 100000
const SUCCESS int = 0
const NO_ERROR int = 0

// Object
const ERROR_PROG int = 0
const ERROR_BRAIN int = 10000
const ERROR_MSG int = 20000
const ERROR_DEBUG int = 30000
const ERROR_COMMAND int = 40000
const ERROR_SERVER int = 50000
const ERROR_CLIENT int = 60000
const ERROR_CELL int = 70000
const ERROR_CONFIG int = 80000

// Error type
const ERROR_NOT_FOUND = 1
const ERROR_EMPTY = 2
const ERROR_INVALID_DATA = 3
const ERROR_NOT_SET = 4
const ERROR_NO_NETWORK = 5
const ERROR_NOT_CONNECT = 6
const ERROR_OPEN = 7
const ERROR_CLOSE = 8
const ERROR_CREATE = 9
const ERROR_WRITE = 10
const ERROR_READ = 11
const ERROR_MEMORY_ALLOC = 12

// Warning code
const WARNING_BRAIN_EMPTY int = WARNING + ERROR_BRAIN + ERROR_EMPTY
const WARNING_DEBUG_NOT_SET int = WARNING + ERROR_DEBUG + ERROR_NOT_SET
const WARNING_COMMAND_NOT_FOUND int = WARNING + ERROR_COMMAND + ERROR_NOT_FOUND
const WARNING_SERVER_NOT_CONNECT int = WARNING + ERROR_SERVER + ERROR_NOT_CONNECT
const WARNING_CELL_INVALID_DATA int = WARNING + ERROR_CELL + ERROR_INVALID_DATA
const WARNING_CELL_NOT_SET int = WARNING + ERROR_CELL + ERROR_NOT_SET

// Error code
const ERROR_MSG_NOT_FOUND int = ERROR + ERROR_MSG + ERROR_NOT_FOUND
const ERROR_SERVER_WRITE int = ERROR + ERROR_SERVER + ERROR_WRITE
const ERROR_SERVER_READ int = ERROR + ERROR_SERVER + ERROR_READ
const ERROR_CLIENT_WRITE int = ERROR + ERROR_CLIENT + ERROR_WRITE
const ERROR_CLIENT_READ int = ERROR + ERROR_CLIENT + ERROR_READ

// Error critical code
const ERROR_CRITICAL_BRAIN_NOT_FOUND int = ERROR_CRITICAL + ERROR_BRAIN + ERROR_NOT_FOUND
const ERROR_CRITICAL_COMMAND_INVALID_DATA int = ERROR_CRITICAL + ERROR_COMMAND + ERROR_INVALID_DATA
const ERROR_CRITICAL_CONFIG_INVALID_DATA int = ERROR_CRITICAL + ERROR_CONFIG + ERROR_INVALID_DATA

// Error fatal code
const ERROR_FATAL_PROG_NOT_FOUND int = ERROR_FATAL + ERROR_PROG + ERROR_NOT_FOUND
const ERROR_FATAL_PROG_MEMORY_ALLOC int = ERROR_FATAL + ERROR_PROG + ERROR_MEMORY_ALLOC
const ERROR_FATAL_BRAIN_INVALID int = ERROR_FATAL + ERROR_BRAIN + ERROR_INVALID_DATA
const ERROR_FATAL_SERVER_NO_NETWORK int = ERROR_FATAL + ERROR_SERVER + ERROR_NO_NETWORK
const ERROR_FATAL_CLIENT_INVALID_DATA int = ERROR_FATAL + ERROR_CLIENT + ERROR_INVALID_DATA
const ERROR_FATAL_CLIENT_NOT_CONNECT int = ERROR_FATAL + ERROR_CLIENT + ERROR_NOT_CONNECT
const ERROR_FATAL_CONFIG_CREATE int = ERROR_FATAL + ERROR_CONFIG + ERROR_CREATE
const ERROR_FATAL_CONFIG_WRITE int = ERROR_FATAL + ERROR_CONFIG + ERROR_WRITE
const ERROR_FATAL_CONFIG_OPEN int = ERROR_FATAL + ERROR_CONFIG + ERROR_OPEN
const ERROR_FATAL_CONFIG_READ int = ERROR_FATAL + ERROR_CONFIG + ERROR_READ

/*******************
* Error string
*******************/
var errorString = map[int]string{
	// Program
	ERROR_FATAL_PROG_NOT_FOUND:    "The program path cannot be found",
	ERROR_FATAL_PROG_MEMORY_ALLOC: "Cannot allocate memory for the program",

	// Brain
	WARNING_BRAIN_EMPTY:            "The brain '%s' is empty.",
	ERROR_FATAL_BRAIN_INVALID:      "The brain is not consistent.",
	ERROR_CRITICAL_BRAIN_NOT_FOUND: "The brain '%s' context is not found.",

	// Debug
	WARNING_DEBUG_NOT_SET: "The debug tags not set.",

	// Command Line
	WARNING_COMMAND_NOT_FOUND:           "The command '%s' is not found.",
	ERROR_CRITICAL_COMMAND_INVALID_DATA: "The command line have incompatible arguments. (%s, %s).",

	// Server
	WARNING_SERVER_NOT_CONNECT:    "The connection to the daemon was fail.\n%s",
	ERROR_SERVER_WRITE:            "Error when writing data to the client.\n%s",
	ERROR_SERVER_READ:             "Error when reading data from the client.\n%s",
	ERROR_FATAL_SERVER_NO_NETWORK: "Cannot open network connection (%s).\n%s",

	// Client
	ERROR_CLIENT_WRITE:              "Error when writing data to the server.\n%s",
	ERROR_CLIENT_READ:               "Error when reading data from the server.\n%s",
	ERROR_FATAL_CLIENT_INVALID_DATA: "Cannot resolve TCP connection to the server (%s).\n%s",
	ERROR_FATAL_CLIENT_NOT_CONNECT:  "Cannot connect to the server.\n%s",

	// Cell
	WARNING_CELL_INVALID_DATA: "Invalid cell, return and empty cell instead",
	WARNING_CELL_NOT_SET:      "Nil cell",

	// Configuration
	ERROR_CRITICAL_CONFIG_INVALID_DATA: "The configuration file have invalid data",
	ERROR_FATAL_CONFIG_OPEN:            "Cannot open the configuration file",
	ERROR_FATAL_CONFIG_CREATE:          "Cannot create the configuration file",
	ERROR_FATAL_CONFIG_WRITE:           "Cannot write to the configuration file",
	ERROR_FATAL_CONFIG_READ:            "Cannot read the configuration file",

	// Message
	ERROR_MSG_NOT_FOUND: "The error message is not found."}

/*******************
* GetStrFromErrorCode
*******************/
// Get the error message from the error code
func GetStrFromErrorCode(errorCode int, extraVar ...any) string {
	msg, exists := errorString[errorCode]
	if !exists {
		return errorString[ERROR_MSG_NOT_FOUND]
	}

	if len(extraVar) > 0 {
		msg = fmt.Sprintf(msg, extraVar...)
	}

	return msg
}

/*******************
* GetStrFromErrorCodeColor
*******************/
// Get the error message of the error code with ANSI escape code for color.
func GetStrFromErrorCodeColor(errorCode int, extraVar ...any) string {
	color := ""

	if errorCode >= ERROR_CRITICAL {
		color = RED
	} else if errorCode >= ERROR {
		color = RED
	} else if errorCode >= WARNING {
		color = YELLOW
	} else if errorCode >= INFO {
		color = BLUE
	}

	return fmt.Sprintf("%s%s%s", color, GetStrFromErrorCode(errorCode, extraVar...), RESET_ALL)
}

/*******************
* PrintMsgFromErrorCode
*******************/
func PrintMsgFromErrorCode(errorCode int, extraVar ...any) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	fmt.Printf("%s (%d) : %s\n", file, line, GetStrFromErrorCodeColor(errorCode, extraVar...))
}

/*******************
* Error_PrintMsgFromErrorCode
*******************/
// Print the error message of the error code
func PrintCallStack() {
	debug.PrintStack()
}
