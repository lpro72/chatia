package text

/*******************
* Import
*******************/
import (
	"encoding/binary"

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
func (letterData *S_LetterCellData) DumpCell(currentCell interfaces.I_Cell, indentation []byte) {
	println("letters_data/DumpCell")
	//	if letterData.WordCellID != 0 {
	//		fmt.Printf("---------------\n")
	//		//        fmt.Printf("%sCurrent Cell : %v\n", indentation, currentCell)
	//		//        fmt.Printf("%sLetter Data : %v\n", indentation, letterData)
	//		wordCell := currentCell.GetBrain().GetCellFromID(letterData.WordCellID)
	//		wordCellData := templates.GetDataFromCell[*S_WordCellData](wordCell)
	//		wordCellData.DumpCell(wordCell, indentation)
	//	} else {
	//
	//		fmt.Printf("---------------\n")
	//		//        fmt.Printf("%sCurrent Cell : %v\n", indentation, currentCell)
	//		//        fmt.Printf("%sLetter Data : %v\n", indentation, letterData)
	//		fmt.Printf("%sLetter : %c, count : %d\n", indentation, letterData.Letter, letterData.Count)
	//	}
	//
	//	for childCell := currentCell.GetFirstChildCell(); childCell != nil; childCell = childCell.GetNextCell() {
	//		dataCell := childCell.GetData()
	//		if dataCell == nil {
	//			// Return and empty cell
	//			errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA, "text_letters_data_debug.go")
	//			continue
	//		}
	//		dataCell.DumpCell(childCell, append(indentation, []byte{' '}...))
	//	}
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
