//go:build Text && !debug

package brain

/*******************
* Import
*******************/
import (
    "chatia/modules/error"
)

/*******************
* DumpMemoryText
*******************/
func DumpMemoryText() {
    error.PrintMsgFromErrorCode(error.WARNING_DEBUG_NOT_SET)
}


