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
	Count           int
	LetterSynapseID uint32
	WordSynapseID   uint32
	FormatSynapseID uint32
}

/*******************
* Interface I_CellData
*******************/
func (textData *S_TextData) DumpCell(currentCell interfaces.I_Cell, indentation []byte) {

	println("text_data/DumpCell")
}

func (textData *S_TextData) GetSerializedData() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint32(buf[0:4], uint32(textData.Count))
	binary.BigEndian.PutUint32(buf[4:8], uint32(textData.LetterSynapseID))
	binary.BigEndian.PutUint32(buf[8:12], uint32(textData.WordSynapseID))
	binary.BigEndian.PutUint32(buf[12:16], uint32(textData.FormatSynapseID))
	return buf
}

/*******************
* CreateTextCellFromSerializeData
*******************/
func CreateTextCellFromSerializeData(dataSerialized []byte) interfaces.I_CellData {
	if len(dataSerialized) < 16 {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_CELL_READ)
		return nil
	}
	count := binary.BigEndian.Uint32(dataSerialized[0:4])
	letterSynapseID := binary.BigEndian.Uint32(dataSerialized[4:8])
	wordSynapseID := binary.BigEndian.Uint32(dataSerialized[8:12])
	formatSynapseID := binary.BigEndian.Uint32(dataSerialized[12:16])
	textData := new(S_TextData)
	textData.Count = int(count)
	textData.LetterSynapseID = letterSynapseID
	textData.WordSynapseID = wordSynapseID
	textData.FormatSynapseID = formatSynapseID
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
