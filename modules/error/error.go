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

// Error type
const ERROR_NOT_FOUND                       = 1
const ERROR_EMPTY                           = 2
const ERROR_INVALID_DATA                    = 3
const ERROR_NOT_SET                         = 4

// Warning message
const WARNING_BRAIN_EMPTY int               = WARNING + ERROR_BRAIN + ERROR_EMPTY
const WARNING_DEBUG_NOT_SET int             = WARNING + ERROR_DEBUG + ERROR_NOT_SET
const WARNING_COMMAND_NOT_FOUND int         = WARNING + ERROR_COMMAND + ERROR_NOT_FOUND

// Error message
const ERROR_MSG_NOT_FOUND int               = ERROR + ERROR_MSG + ERROR_NOT_FOUND
const ERROR_CRITICAL_BRAIN_NOT_FOUND int    = ERROR_CRITICAL + ERROR_BRAIN + ERROR_NOT_FOUND
const ERROR_FATAL_BRAIN_INVALID int         = ERROR_FATAL + ERROR_BRAIN + ERROR_INVALID_DATA

/*******************
* Error string
*******************/
var errorString = map[int]string{
    WARNING_BRAIN_EMPTY:                    "The brain '%s' is empty",
    WARNING_DEBUG_NOT_SET:                  "The debug tags not set",
    WARNING_COMMAND_NOT_FOUND:              "The command '%s' is not found",
    ERROR_MSG_NOT_FOUND:                    "The error message is not found",
    ERROR_CRITICAL_BRAIN_NOT_FOUND:         "The brain '%s' context is not found",
    ERROR_FATAL_BRAIN_INVALID:              "The brain is not consistent"}

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
 
