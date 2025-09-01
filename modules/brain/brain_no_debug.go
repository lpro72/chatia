//go:build !debug

package brain

/*******************
* Import
*******************/
import (
    "chatia/modules/error"
)

/*******************
* DumpMemory
*******************/
func DumpMemory(name string) {
    error.PrintMsgFromErrorCode(error.WARNING_DEBUG_NOT_SET)
}


