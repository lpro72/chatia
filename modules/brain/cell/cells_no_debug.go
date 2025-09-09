//go:build !debug

package cell

/*******************
* Import
*******************/
import (
        "chatia/modules/error"
)

/*******************
* DumpCell
*******************/
func DumpCell(currentCell *CellStruct, indentation []byte) {
    error.PrintMsgFromErrorCode(error.WARNING_DEBUG_NOT_SET)
}


