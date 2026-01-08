package text

/*******************
* Import
*******************/
import (
	"chatia/modules/data"
	"chatia/modules/interfaces"
)

// 	"chatia/modules/errcode"
// 	"chatia/modules/interfaces"
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
	println("letters_data/GetSerializedData")
	return []byte("")
}

/*******************
* CreateLetterCellFromSerializeData
*******************/
func CreateLetterCellFromSerializeData(dataSerialized []byte) interfaces.I_CellData {
	letterData := new(S_LetterCellData)
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
