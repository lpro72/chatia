package text

/*******************
* Import
*******************/
import (
	"encoding/binary"
	"fmt"

	"chatia/modules/data"
	"chatia/modules/interfaces"
)

// 	"chatia/modules/errcode"
// 	"chatia/modules/templates"

/*******************
* S_LetterCellData
*******************/
type S_LetterCellData struct {
	Count      int
	Letter     rune
	WordCellID uint32
}

/*******************
* Interface I_CellData
*******************/
func (letterData *S_LetterCellData) DumpData() {
	fmt.Printf("Letter: %c, Count: %d, WordCellID: %d\n", letterData.Letter, letterData.Count, letterData.WordCellID)
}

func (letterData *S_LetterCellData) GetSerializedData() []byte {
	buf := make([]byte, 12)

	binary.LittleEndian.PutUint32(buf[0:4], uint32(letterData.Count))
	binary.LittleEndian.PutUint32(buf[4:8], uint32(letterData.Letter))
	binary.LittleEndian.PutUint32(buf[8:12], letterData.WordCellID)

	return buf
}

/*******************
* CreateLetterCellFromSerializeData
*******************/
func CreateLetterCellFromSerializeData(dataSerialized []byte) interfaces.I_CellData {
	letterData := new(S_LetterCellData)
	letterData.Count = int(binary.LittleEndian.Uint32(dataSerialized[0:4]))
	letterData.Letter = rune(binary.LittleEndian.Uint32(dataSerialized[4:8]))
	letterData.WordCellID = binary.LittleEndian.Uint32(dataSerialized[8:12])
	return letterData
}

/*******************
* LetterCell_Create
*******************/
func LetterCell_Create(brainConfig interfaces.I_BrainConfig, letter rune) interfaces.I_Cell {
	// New cell must be created
	newLetterData := new(S_LetterCellData)
	newCell := data.CreateCell(brainConfig, newLetterData, g_LetterCellType)
	newLetterData.Letter = letter
	return (newCell)
}
