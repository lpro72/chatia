package text

/*******************
* Import
*******************/
import (
	"encoding/binary"

	"chatia/modules/data"
	"chatia/modules/errcode"
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
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(textData.Count))
	return buf
}

/*******************
* CreateTextCellFromSerializeData
*******************/
func CreateTextCellFromSerializeData(dataSerialized []byte) interfaces.I_CellData {
	errcode.PrintCallStack()
	if len(dataSerialized) < 4 {
		return nil
	}
	count := binary.BigEndian.Uint32(dataSerialized[0:4])
	textData := new(S_TextData)
	textData.Count = int(count)
	return textData
}

/*******************
* TextCell_Create
*******************/
func TextCell_Create(brainContext interfaces.I_BrainContext) interfaces.I_Cell {
	newTextData := new(S_TextData)
	newCell := data.CreateCell(brainContext.GetBrainConfig(), newTextData, g_TextCellType)
	return (newCell)
}
