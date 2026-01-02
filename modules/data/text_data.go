package data

/*******************
* Import
*******************/
import (
	"fmt"

	"chatia/modules/interfaces"
)

/*******************
* Types
*******************/
type S_TextData struct {
	Count         int
	LetterSynapse interfaces.I_Synapse
	WordSynapse   interfaces.I_Synapse
}

/*******************
* Interface I_CellData
*******************/
func (textData *S_TextData) DumpCell(currentCell interfaces.I_Cell, indentation []byte) {

	println("text_data/DumpCell")
}

func (textData *S_TextData) GetSerializedData() []byte {
	println("text_data/GetSerializedData")
	return []byte(fmt.Sprintf("%d", textData.Count))
}
