//go:build !debug

package cell

/*******************
* DumpCell
*******************/
func DumpCell(currentCell *CellStruct, indentation []byte) {
    error.PrintMsgFromErrorCode(error.WARNING_DEBUG_NOT_SET)
}


