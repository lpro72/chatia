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
type S_LetterCellData struct {
	Count      int
	Letter     rune
	WordCellID int
}

/*******************
* Globals Varables
*******************/
var g_LetterCellType int = 0

/*******************
* init
*******************/
func init() {
	g_LetterCellType = RegisterCellType("Letter")
}

/*******************
* DumpCell
*******************/
func (letterData *S_LetterCellData) DumpCell(currentCell I_CellManagement, indentation []byte) {
	if letterData.WordCellID != 0 {
		fmt.Printf("---------------\n")
		//        fmt.Printf("%sCurrent Cell : %v\n", indentation, currentCell)
		//        fmt.Printf("%sLetter Data : %v\n", indentation, letterData)
		wordCell := GetCellFromGroup(letterData.WordCellID)
		wordCellData := GetDataFromWordCell(wordCell)
		wordCellData.DumpCell(wordCell, indentation)
	} else {
		fmt.Printf("---------------\n")
		//        fmt.Printf("%sCurrent Cell : %v\n", indentation, currentCell)
		//        fmt.Printf("%sLetter Data : %v\n", indentation, letterData)
		fmt.Printf("%sLetter : %c, count : %d\n", indentation, letterData.Letter, letterData.Count)
	}
	for childCell := currentCell.GetFirstChildCell(); childCell != nil; childCell = childCell.GetNextCell() {
		dataCell := childCell.GetData()
		if dataCell == nil {
			// Return and empty cell
			errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA, "text_letters_data_debug.go")
			continue
		}
		dataCell.DumpCell(childCell, append(indentation, []byte{' '}...))
	}
}

/*******************
* CreateLetterCell
*******************/
func CreateLetterCell(letter rune, parentCell I_CellManagement) I_CellManagement {
	// New cell must be created
	newLetterData := new(S_LetterCellData)
	newCell := CreateCell(parentCell, newLetterData, g_LetterCellType)
	newLetterData.Letter = letter
	return (newCell)
}

/*******************
* GetDataFromLetterCell
*******************/
func GetDataFromLetterCell(currentCell I_CellManagement) *S_LetterCellData {
	if currentCell == nil {
		// Return and empty cell
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_NOT_SET)
		errcode.PrintCallStack()
		return new(S_LetterCellData)
	}

	// Must be a letter cell
	letterData, ok := currentCell.GetData().(*S_LetterCellData)
	if !ok {
		// Return and empty cell
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA)
		fmt.Printf("%v\n", currentCell.GetData())
		errcode.PrintCallStack()
		return new(S_LetterCellData)
	}

	return (letterData)
}
