package text

/*******************
* Import
*******************/
import (
	// 	"fmt"

	"chatia/modules/data"
	"chatia/modules/interfaces"
	// "chatia/modules/templates"
)

/*******************
* Types
*******************/
type S_WordCellData struct {
	Count              int
	FirstLetterSynapse interfaces.I_Synapse
	LastLetterSynapse  interfaces.I_Synapse
	Word               string
}

/*******************
* Interface I_CellData
*******************/
func (wordData *S_WordCellData) DumpCell(currentCell interfaces.I_Cell, indentation []byte) {
	println("text_words/DumpCell")
	// letterData := templates.GetDataFromCell[*S_LetterCellData](wordData.LastLetterCell)
	// fmt.Printf("%sLetter : %c, count : %d, word : %s, Word Count : %d\n", indentation, letterData.Letter, letterData.Count, wordData.Word, wordData.Count)
}

func (wordData *S_WordCellData) GetSerializedData() []byte {
	return []byte("")
}

/*******************
* CreateWordCellFromSerializeData
*******************/
func CreateWordCellFromSerializeData(dataSerialized []byte) interfaces.I_CellData {
	wordData := new(S_WordCellData)
	return wordData
}

/*******************
* WordCell_Create
*******************/
func WordCell_Create(brainConfig interfaces.I_BrainConfig, parentSynapse interfaces.I_Synapse, FirstLetterSynapse interfaces.I_Synapse, lastLetterSynapse interfaces.I_Synapse) interfaces.I_Cell {
	newWordData := new(S_WordCellData)
	newCell := data.CreateCell(brainConfig, newWordData, g_WordCellType)
	newWordData.LastLetterSynapse = lastLetterSynapse
	newWordData.FirstLetterSynapse = FirstLetterSynapse
	newWordData.Word = LetterCell_GetWordFromLastSynapse(lastLetterSynapse)
	return (newCell)
}
