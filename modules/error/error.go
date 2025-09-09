package error

/*******************
* Import
*******************/
import (
        "fmt"
)

/*******************
* ANSI escape code color
*******************/
const RED                                   = "\033[33m"
const YELLOW                                = "\033[33m"
const BLUE                                  = "\033[34m"
const RESET_ALL                             = "\033[0m"

/*******************
* Error code
*******************/
// Severity
const ERROR_FATAL                           = 500000
const ERROR_CRITICAL int                    = 400000
const ERROR int                             = 300000
const WARNING int                           = 200000
const INFO int                              = 100000
const SUCCESS int                           = 0
const NO_ERROR int                          = 0

// Object
const ERROR_BRAIN int                       = 10000
const ERROR_MSG int                         = 20000
const ERROR_DEBUG int                       = 30000
const ERROR_COMMAND int                     = 40000
const ERROR_SERVER int                      = 50000
const ERROR_CLIENT int                      = 60000

// Error type
const ERROR_NOT_FOUND                       = 1
const ERROR_EMPTY                           = 2
const ERROR_INVALID_DATA                    = 3
const ERROR_NOT_SET                         = 4
const ERROR_NO_NETWORK                      = 5
const ERROR_NOT_CONNECT                     = 6
const ERROR_WRITE                           = 7
const ERROR_READ                            = 8

// Warning code
const WARNING_BRAIN_EMPTY int               = WARNING + ERROR_BRAIN + ERROR_EMPTY
const WARNING_DEBUG_NOT_SET int             = WARNING + ERROR_DEBUG + ERROR_NOT_SET
const WARNING_COMMAND_NOT_FOUND int         = WARNING + ERROR_COMMAND + ERROR_NOT_FOUND
const WARNING_SERVER_NOT_CONNECT int        = WARNING + ERROR_SERVER + ERROR_NOT_CONNECT

// Error code
const ERROR_MSG_NOT_FOUND int               = ERROR + ERROR_MSG + ERROR_NOT_FOUND
const ERROR_SERVER_WRITE int                = ERROR + ERROR_SERVER + ERROR_WRITE
const ERROR_SERVER_READ int                 = ERROR + ERROR_SERVER + ERROR_READ
const ERROR_CLIENT_WRITE int                = ERROR + ERROR_CLIENT + ERROR_WRITE
const ERROR_CLIENT_READ int                 = ERROR + ERROR_CLIENT + ERROR_READ

// Error critical code
const ERROR_CRITICAL_BRAIN_NOT_FOUND int    = ERROR_CRITICAL + ERROR_BRAIN + ERROR_NOT_FOUND

// Error fatal code
const ERROR_FATAL_BRAIN_INVALID int         = ERROR_FATAL + ERROR_BRAIN + ERROR_INVALID_DATA
const ERROR_FATAL_SERVER_NO_NETWORK int     = ERROR_FATAL + ERROR_SERVER + ERROR_NO_NETWORK
const ERROR_FATAL_CLIENT_INVALID_DATA int   = ERROR_FATAL + ERROR_CLIENT + ERROR_INVALID_DATA
const ERROR_FATAL_CLIENT_NOT_CONNECT int    = ERROR_FATAL + ERROR_CLIENT + ERROR_NOT_CONNECT

/*******************
* Error string
*******************/
var errorString = map[int]string{
    // Brain
    WARNING_BRAIN_EMPTY:                    "The brain '%s' is empty.",
    ERROR_FATAL_BRAIN_INVALID:              "The brain is not consistent.",
    ERROR_CRITICAL_BRAIN_NOT_FOUND:         "The brain '%s' context is not found.",

    // Debug
    WARNING_DEBUG_NOT_SET:                  "The debug tags not set.",

    // Command Line
    WARNING_COMMAND_NOT_FOUND:              "The command '%s' is not found.",

    // Server
    WARNING_SERVER_NOT_CONNECT:             "The connection to the daemon was fail.\n%s",
    ERROR_SERVER_WRITE:                     "Error when writing data to the client.\n%s",
    ERROR_SERVER_READ:                      "Error when reading data from the client.\n%s",
    ERROR_FATAL_SERVER_NO_NETWORK:          "Cannot open network connection (%s).\n%s",

    // Client
    ERROR_CLIENT_WRITE:                     "Error when writing data to the server.\n%s",
    ERROR_CLIENT_READ:                      "Error when reading data from the server.\n%s",
    ERROR_FATAL_CLIENT_INVALID_DATA:        "Cannot resolve TCP vonnection to the server (%s).\n%s",
    ERROR_FATAL_CLIENT_NOT_CONNECT:         "Cannot connect to the server.\n%s",

    // Message
    ERROR_MSG_NOT_FOUND:                    "The error message is not found."}

/*******************
* GetStrFromErrorCode
*******************/
// Get the error message from the error code
func GetStrFromErrorCode(errorCode int, extraVar ...any) string {
    msg, exists := errorString[errorCode]
    if ! exists {
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
* Error_PrintMsgFromErrorCode
*******************/
// Print the error message of the error code
func PrintMsgFromErrorCode(errorCode int, extraVar ...any) {
    fmt.Println(GetStrFromErrorCodeColor(errorCode, extraVar...))
 }
 
