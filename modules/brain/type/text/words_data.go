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
func (wordData *S_WordCellData) DumpData() {
	fmt.Printf("Word Count: %d, FirstLetterSynapseID: %d, LastLetterSynapseID: %d, Word: %s\n", wordData.Count, wordData.FirstLetterSynapseID, wordData.LastLetterSynapseID, wordData.Word)
}

func (wordData *S_WordCellData) GetSerializedData() []byte {
	formatted := fmt.Sprintf("%d-%d-%d-%s", wordData.FirstLetterSynapseID, wordData.LastLetterSynapseID, wordData.Count, wordData.Word)
	return []byte(formatted)
}

/*******************
* CreateWordCellFromSerializeData
*******************/
func CreateWordCellFromSerializeData(dataSerialized []byte) interfaces.I_CellData {
	wordData := new(S_WordCellData)
	reader := bytes.NewReader(dataSerialized)
	_, err := fmt.Fscanf(reader, "%d-%d-%d-%s", &wordData.FirstLetterSynapseID, &wordData.LastLetterSynapseID, &wordData.Count, &wordData.Word)
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
	newWordData.LastLetterSynapseID = lastLetterSynapse.GetID()
	newWordData.FirstLetterSynapseID = FirstLetterSynapse.GetID()
	newWordData.Word = LetterCell_GetWordFromLastSynapse(lastLetterSynapse)
	newCell := data.CreateCell(brainConfig, newWordData, g_WordCellType)
	return (newCell)
}
