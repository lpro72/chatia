//go:build Chess && !debug

package brain

/*******************
* Import
*******************/
import (
    "chatia/modules/error"
)

/*******************
* DumpMemoryChess
*******************/
func DumpMemoryChess() {
    error.PrintMsgFromErrorCode(error.WARNING_DEBUG_NOT_SET)
}

/*******************
* UnittestChess
*******************/
func UnittestChess() {
}


