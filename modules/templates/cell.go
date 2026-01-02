package templates

/*******************
* Import
*******************/
import (
	"chatia/modules/errcode"
	"chatia/modules/interfaces"
)

/*******************
* GetDataFromCell
*******************/
func GetDataFromCell[T any](currentCell interfaces.I_Cell) T {
	var zero T
	if currentCell == nil {
		// Return and empty cell
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_NOT_SET)
		errcode.PrintCallStack()
		return zero
	}

	// Must be a letter cell
	cellData, ok := currentCell.GetData().(T)
	if !ok {
		// Return and empty cell
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA)
		errcode.PrintCallStack()
		return zero
	}

	return (cellData)
}
