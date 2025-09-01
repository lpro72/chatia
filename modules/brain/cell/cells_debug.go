//go:build debug

package cell

/*******************
* DumpCell
*******************/
func DumpCell(currentCell *CellStruct, indentation []byte) {
    if currentCell.Cell != nil {
        currentCell.Cell.DumpCell(currentCell, indentation)
    }
}


