package text

/*******************
* Import
*******************/
import (
	"bytes"
	"fmt"

	"chatia/modules/data"
	"chatia/modules/errcode"
	"chatia/modules/interfaces"
	// "chatia/modules/templates"
)

/*******************
* Types
*******************/
type S_WordCellData struct {
	Count                int
	FirstLetterSynapseID uint32
	LastLetterSynapseID  uint32
	Word                 string
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
	return fmt.Appendf(nil, "%s-%d-%d-%d", wordData.Word, wordData.FirstLetterSynapseID, wordData.LastLetterSynapseID, wordData.Count)
}

/*******************
* CreateWordCellFromSerializeData
*******************/
func CreateWordCellFromSerializeData(dataSerialized []byte) interfaces.I_CellData {
	wordData := new(S_WordCellData)
	reader := bytes.NewReader(dataSerialized)
	_, err := fmt.Fscanf(reader, "%s-%d-%d-%d", &wordData.Word, &wordData.FirstLetterSynapseID, &wordData.LastLetterSynapseID, &wordData.Count)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_CELL_READ)
	}
	return wordData
}

/*******************
* WordCell_Create
*******************/
func WordCell_Create(brainConfig interfaces.I_BrainConfig, FirstLetterSynapse interfaces.I_Synapse, lastLetterSynapse interfaces.I_Synapse) interfaces.I_Cell {
	newWordData := new(S_WordCellData)
	newCell := data.CreateCell(brainConfig, newWordData, g_WordCellType)
	newWordData.LastLetterSynapseID = lastLetterSynapse.GetID()
	newWordData.FirstLetterSynapseID = FirstLetterSynapse.GetID()
	newWordData.Word = LetterCell_GetWordFromLastSynapse(lastLetterSynapse)
	return (newCell)
}
