package text

/*******************
* Import
*******************/
import (
	//	"math/rand"

	"chatia/modules/data"
	"chatia/modules/interfaces"
	// "chatia/modules/templates"
)

/*******************
* Globals Varables
*******************/
var g_WordCellType uint32 = 0

/*******************
* WordCell_Register
*******************/
func WordCell_Register() {
	g_WordCellType = data.CellType_RegisterNewType("Word", CreateWordCellFromSerializeData)
}

/*******************
* WordCell_GetRandowWord
*******************/
func WordCell_GetRandowWord(currentCell interfaces.I_Cell, count int) string {
	println("text_words/WordCell_GetRandowWord")
	// 	for nextCellID := rand.Intn(count) + 1; currentCell != nil; currentCell = currentCell.GetNextCell() {
	// 		wordData := templates.GetDataFromCell[*data.S_WordCellData](currentCell)
	// 		if nextCellID <= wordData.Count {
	// 			return wordData.Word
	// 		}
	// 		nextCellID -= wordData.Count
	// 	}
	return ""
}
