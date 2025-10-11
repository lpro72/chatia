//go:build !NoText

package cell

/*******************
* Import
*******************/
import (
	"chatia/modules/errcode"
	"fmt"
)

/*******************
* Types
*******************/
type S_TextData struct {
	Count      int
	LetterCell I_CellManagement
	WordCell   I_CellManagement
}

/*******************
* Globals Varables
*******************/
var g_TextCellType int = 0

/*******************
* init
*******************/
func init() {
	g_TextCellType = RegisterCellType("Text")
}

/*******************
* DumpCell
*******************/
func (textData *S_TextData) DumpCell(currentCell I_CellManagement, indentation []byte) {
}

/*******************
* CreateLetterCell
*******************/
func CreateTextCell() I_CellManagement {
	// New cell must be created
	newTextData := new(S_TextData)
	newCell := CreateCell(nil, newTextData, g_TextCellType)
	return (newCell)
}

/*******************
* GetDataFromLetterCell
*******************/
func GetDataFromTextCell(currentCell I_CellManagement) *S_TextData {
	if currentCell == nil {
		// Return and empty cell
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_NOT_SET)
		errcode.PrintCallStack()
		return new(S_TextData)
	}

	// Must be a letter cell
	textData, ok := currentCell.GetData().(*S_TextData)
	if !ok {
		// Return and empty cell
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA)
		fmt.Printf("%v\n", currentCell.GetData())
		errcode.PrintCallStack()
		return new(S_TextData)
	}

	return (textData)
}
