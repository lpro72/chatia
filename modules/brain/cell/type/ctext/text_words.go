package ctext

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
* CreateWordCellFromSerializeData
*******************/
func CreateWordCellFromSerializeData(dataSerialized []byte) interfaces.I_CellData {
	wordData := new(data.S_WordCellData)
	return wordData
}

/*******************
* WordCell_Create
*******************/
func WordCell_Create(brainConfig interfaces.I_BrainConfig, parentSynapse interfaces.I_Synapse, FirstLetterSynapse interfaces.I_Synapse, lastLetterSynapse interfaces.I_Synapse) interfaces.I_Cell {
	newWordData := new(data.S_WordCellData)
	newCell := data.CreateCell(brainConfig, newWordData, g_WordCellType)
	newWordData.LastLetterSynapse = lastLetterSynapse
	newWordData.FirstLetterSynapse = FirstLetterSynapse
	newWordData.Word = LetterCell_GetWordFromLastSynapse(lastLetterSynapse)
	return (newCell)
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
