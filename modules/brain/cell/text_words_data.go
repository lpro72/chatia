//go:build Text

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
type S_WordCellData struct {
	Count           int
	FirstLetterCell I_CellManagement
	LastLetterCell  I_CellManagement
	Word            string
}

/*******************
* Globals Varables
*******************/
var g_WordCellType int = 0

/*******************
* init
*******************/
func init() {
	g_WordCellType = RegisterCellType("Word")
}

/*******************
* DumpCell
*******************/
func (wordData *S_WordCellData) DumpCell(currentCell I_CellManagement, indentation []byte) {
	letterData := GetDataFromLetterCell(wordData.LastLetterCell)
	//    fmt.Printf("%sCurrent Cell : %p, %v\n", indentation, &currentCell, currentCell)
	//    fmt.Printf("%sWord Data : %p, %v\n", indentation, &wordData, wordData)
	fmt.Printf("%sLetter : %c, count : %d, word : %s, Word Count : %d\n", indentation, letterData.Letter, letterData.Count, wordData.Word, wordData.Count)
}

/*******************
* CreateWordCell
*******************/
func CreateWordCell(parentCell I_CellManagement, firstLetterCell I_CellManagement, lastLetterCell I_CellManagement) I_CellManagement {
	newWordData := new(S_WordCellData)
	newCell := CreateCell(parentCell, newWordData, g_WordCellType)
	newWordData.LastLetterCell = lastLetterCell
	newWordData.FirstLetterCell = firstLetterCell
	newWordData.Word = GetWordFromLastCell(lastLetterCell)
	return (newCell)
}

/*******************
* GetDataFromWordCell
*******************/
func GetDataFromWordCell(currentCell I_CellManagement) *S_WordCellData {
	if currentCell == nil {
		// Return and empty word cell
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_NOT_SET)
		errcode.PrintCallStack()
		return new(S_WordCellData)
	}

	// Must be a word cell
	wordData, ok := currentCell.GetData().(*S_WordCellData)
	if !ok {
		// Return and empty word cell
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA)
		fmt.Printf("%v\n", currentCell.GetData())
		errcode.PrintCallStack()
		return new(S_WordCellData)
	}

	return (wordData)
}
