//go:build !NoText

package brain

/*******************
* Import
*******************/
import (
	"chatia/modules/brain/cell"
	"chatia/modules/errcode"
	"fmt"
	"unicode"
)

/*******************
* ExecText
*******************/
// Learn some texte
func ExecText(command string, extraVar ...any) string {
	switch command {
	case "GetRandomWordFromWordCell":
		return GetRandomWordFromWordCell()
	case "GetRandomWordFromLetterCell":
		return GetRandomWordFromLetterCell()
	}

	errcode.PrintMsgFromErrorCode(errcode.WARNING_COMMAND_NOT_FOUND, command)

	return ""
}

/*******************
* LearnTextFromBrain
*******************/
func LearnTextFromBrain(data []byte, firstCell cell.I_CellManagement) {
	LearnText(string(data), firstCell)
}

/*******************
* LearnText
*******************/
// Learn some texte
func LearnText(text string, textCell cell.I_CellManagement) {
	var currentCell cell.I_CellManagement = nil
	var firstCell cell.I_CellManagement = nil
	var letterData *cell.S_LetterCellData = nil
	var wordCell cell.I_CellManagement = nil
	var wordData *cell.S_WordCellData = nil
	textData := cell.GetDataFromTextCell(textCell)

	// initialisation
	if textData.WordCell == nil {
		textData.WordCell = cell.CreateCell(nil, nil, 0)
	}
	// special case of a new brain
	if textData.LetterCell == nil {
		textData.LetterCell = cell.CreateCell(nil, nil, 0)
	}

	for _, r := range text {
		// New word
		if !unicode.IsLetter(r) {
			// Mark as end of a word
			if currentCell != nil {
				letterData = cell.GetDataFromLetterCell(currentCell)
				if letterData.WordCellID == 0 {
					wordCell = cell.CreateWordCell(textData.WordCell, firstCell, currentCell)
					letterData.WordCellID = wordCell.GetID()
					if textData.WordCell == nil {
						textData.WordCell = wordCell
					}
				} else {
					wordCell = cell.GetCellFromGroup(letterData.WordCellID)
				}
				wordData = cell.GetDataFromWordCell(wordCell)
				wordData.Count += 1
			}
			currentCell = nil
			firstCell = nil
			continue
		}

		// New word
		if currentCell == nil {
			textData.Count += 1
			currentCell = textData.LetterCell
		}

		childCell := cell.SearchForLetterCell(unicode.ToLower(r), currentCell)
		letterData = cell.GetDataFromLetterCell(childCell)
		letterData.Count += 1
		if firstCell == nil {
			firstCell = childCell
		}

		currentCell = childCell
	}

	// Last word
	if currentCell != nil {
		letterData = cell.GetDataFromLetterCell(currentCell)
		if letterData.WordCellID == 0 {
			wordCell := cell.CreateWordCell(textData.WordCell, firstCell, currentCell)
			letterData.WordCellID = wordCell.GetID()
			if textData.WordCell == nil {
				textData.WordCell = wordCell
			}
		} else {
			wordCell = cell.GetCellFromGroup(letterData.WordCellID)
		}
		wordData = cell.GetDataFromWordCell(wordCell)
		wordData.Count += 1
	}

}

/*******************
* IsEmptyBrainForText
*******************/
func IsEmptyBrainForText() bool {
	brain := GetBrainContext("Text")
	textData := cell.GetDataFromTextCell(brain.GetFirstCell())
	if textData == nil || textData.LetterCell == nil || textData.WordCell == nil || textData.LetterCell.GetFirstChildCell() == nil || textData.WordCell.GetFirstChildCell() == nil {
		return true
	}
	return false
}

/*******************
* GetRandomWordFromWordCell
*******************/
func GetRandomWordFromWordCell() string {
	if IsEmptyBrainForText() {
		errcode.PrintMsgFromErrorCode(errcode.WARNING_BRAIN_EMPTY, "Text")
		return ""
	}

	brain := GetBrainContext("Text")
	textData := cell.GetDataFromTextCell(brain.GetFirstCell())
	return cell.GetRandowWordFromWordCells(textData.WordCell.GetFirstChildCell(), textData.Count)
}

/*******************
* GetRandomWordFromLetterCell
*******************/
func GetRandomWordFromLetterCell() string {
	if IsEmptyBrainForText() {
		errcode.PrintMsgFromErrorCode(errcode.WARNING_BRAIN_EMPTY, "Text")
		return ""
	}

	brain := GetBrainContext("Text")
	textData := cell.GetDataFromTextCell(brain.GetFirstCell())
	return string(cell.GetRandowWordFromLetterCell(textData.LetterCell, textData.Count))
}

/*******************
* DumpMemoryText
*******************/
func DumpMemoryText() {
	if IsEmptyBrainForText() {
		errcode.PrintMsgFromErrorCode(errcode.WARNING_BRAIN_EMPTY, "Text")
		return
	}

	brain := GetBrainContext("Text")
	textData := cell.GetDataFromTextCell(brain.GetFirstCell())
	fmt.Printf("count : %d\n", textData.Count)
	for childCell := textData.LetterCell.GetFirstChildCell(); childCell != nil; childCell = childCell.GetNextCell() {
		childData := childCell.GetData()
		if childData == nil {
			// Return and empty cell
			errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA, "text_debug.go")
			errcode.PrintCallStack()
			continue
		}
		childData.DumpCell(childCell, []byte{})
	}
}

/*******************
* UnittestText
*******************/
func UnittestText() {
	fmt.Println("Learning....")
	LearnFromString("This is a test", "Text")
	LearnFromString("Ceci est un test", "Text")
	LearnFromString("The brain work", "Text")
	LearnFromString("Test with the word these", "Text")

	fmt.Println("Dump memory :")
	DumpMemory("Text")

	fmt.Println("Get random words :")
	for i := 0; i < 10; i++ {
		fmt.Printf("\nTest %d\n-------\n", i)
		fmt.Printf("From letter cell : %s\n", Exec("GetRandomWordFromLetterCell", "Text"))
		fmt.Printf("From word cell : %s\n", Exec("GetRandomWordFromWordCell", "Text"))
	}
}
