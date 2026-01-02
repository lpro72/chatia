package data

/*******************
* Import
*******************/
import (
	// 	"fmt"

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

// /*******************
// * Interface I_CellData
// *******************/
// func (wordData *S_WordCellData) DumpCell(currentCell interfaces.I_CellManagement, indentation []byte) {
// 	letterData := templates.GetDataFromCell[*S_LetterCellData](wordData.LastLetterCell)
// 	fmt.Printf("%sLetter : %c, count : %d, word : %s, Word Count : %d\n", indentation, letterData.Letter, letterData.Count, wordData.Word, wordData.Count)
// }

func (wordData *S_WordCellData) GetSerializedData() []byte {
	return []byte("")
}
