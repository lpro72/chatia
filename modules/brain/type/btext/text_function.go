package btext

/*******************
* Import
*******************/
import (
	"unicode"

	"chatia/modules/brain/cell/type/ctext"
	"chatia/modules/data"

	// "chatia/modules/errcode"
	"chatia/modules/interfaces"
	"chatia/modules/templates"
)

/*******************
* ExecText
*******************/
func ExecText(brainContext interfaces.I_BrainContext, command string) string {
	println("text_function/ExecText")
	// 	switch command {
	// 	case "GetRandomWordFromWordCell":
	// 		return GetRandomWordFromWordCell(brainContext)
	// 	case "GetRandomWordFromLetterCell":
	// 		return GetRandomWordFromLetterCell(brainContext)
	// 	}

	// 	errcode.PrintMsgFromErrorCode(errcode.WARNING_COMMAND_NOT_FOUND, command)

	return ""
}

/*******************
* LearnTextFromBrain
*******************/
func LearnTextFromBrain(brainContext interfaces.I_BrainContext, data []byte) {
	LearnText(brainContext, string(data))
}

/*******************
* LearnText
*******************/
func LearnText(brainContext interfaces.I_BrainContext, text string) {
	brainConfig := brainContext.GetBrainConfig()
	textCell := brainContext.GetFirstCell()
	var currentCell interfaces.I_Cell = nil
	var currentSynapse interfaces.I_Synapse = nil
	var firstSynapse interfaces.I_Synapse = nil
	var letterData *data.S_LetterCellData = nil
	var wordCell interfaces.I_Cell = nil
	textData := templates.GetDataFromCell[*data.S_TextData](textCell)

	// initialisation
	if textData.WordSynapse == nil {
		textData.WordSynapse = data.CreateSynapse(brainConfig, nil, nil)
	}
	if textData.LetterSynapse == nil {
		textData.LetterSynapse = data.CreateSynapse(brainConfig, nil, nil)
	}

	for _, r := range text {
		// New word
		if !unicode.IsLetter(r) {
			// Mark as end of a word
			if currentCell != nil {
				letterData = templates.GetDataFromCell[*data.S_LetterCellData](currentCell)
				if letterData.WordCellID == 0 {
					wordCell = ctext.WordCell_Create(brainConfig, textData.WordSynapse, firstSynapse, currentSynapse)
					letterData.WordCellID = wordCell.GetID()
					data.CreateSynapse(brainConfig, textData.WordSynapse, wordCell)
				} else {
					wordCell = brainConfig.GetCellsGroupManagament().GetCellFromID(letterData.WordCellID)
				}
				wordData := templates.GetDataFromCell[*data.S_WordCellData](wordCell)
				wordData.Count += 1
			}
			currentCell = nil
			firstSynapse = nil
			continue
		}

		// New word
		if currentSynapse == nil {
			textData.Count += 1
			currentSynapse = textData.LetterSynapse
		}

		currentSynapse, currentCell := ctext.LetterCell_Search(brainConfig, unicode.ToLower(r), currentSynapse)
		letterData = templates.GetDataFromCell[*data.S_LetterCellData](currentCell)
		letterData.Count += 1
		if firstSynapse == nil {
			firstSynapse = currentSynapse
		}
	}

	// Last word
	if currentCell != nil {
		letterData = templates.GetDataFromCell[*data.S_LetterCellData](currentCell)
		if letterData.WordCellID == 0 {
			wordCell = ctext.WordCell_Create(brainConfig, textData.WordSynapse, firstSynapse, currentSynapse)
			letterData.WordCellID = wordCell.GetID()
			data.CreateSynapse(brainConfig, textData.WordSynapse, wordCell)
		} else {
			wordCell = brainConfig.GetCellsGroupManagament().GetCellFromID(letterData.WordCellID)
		}
		wordData := templates.GetDataFromCell[*data.S_WordCellData](wordCell)
		wordData.Count += 1
	}
}

/*******************
* IsEmptyBrainForText
*******************/
func IsEmptyBrainForText(brainContext interfaces.I_BrainContext) bool {
	println("text_function/IsEmptyBrainForText")
	// 	textData := templates.GetDataFromCell[*data.S_TextData](brainContext.GetFirstCell())
	// 	if textData == nil || textData.LetterCell == nil || textData.WordCell == nil || textData.LetterCell.GetFirstChildCell() == nil || textData.WordCell.GetFirstChildCell() == nil {
	// 		return true
	// 	}
	return false
}

/*******************
* GetRandomWordFromWordCell
*******************/
func GetRandomWordFromWordCell(brainContext interfaces.I_BrainContext) string {
	println("text_function/GetRandomWordFromWordCell")
	// 	if IsEmptyBrainForText(brainContext) {
	// 		errcode.PrintMsgFromErrorCode(errcode.WARNING_BRAIN_EMPTY, "Text")
	// 		return ""
	// 	}

	// 	textData := templates.GetDataFromCell[*data.S_TextData](brainContext.GetFirstCell())
	// 	return ctext.WordCell_GetRandowWord(textData.WordCell.GetFirstChildCell(), textData.Count)
	return ""
}

/*******************
* GetRandomWordFromLetterCell
*******************/
func GetRandomWordFromLetterCell(brainContext interfaces.I_BrainContext) string {
	println("text_function/GetRandomWordFromLetterCell")
	// if IsEmptyBrainForText(brainContext) {
	// 		errcode.PrintMsgFromErrorCode(errcode.WARNING_BRAIN_EMPTY, "Text")
	//  	return ""
	// }

	// 	textData := templates.GetDataFromCell[*data.S_TextData](brainContext.GetFirstCell())
	// 	return string(ctext.LetterCell_GetRandowWord(textData.LetterCell, textData.Count))
	return ""
}

/*******************
* DumpMemoryText
*******************/
func DumpMemoryText(brainContext interfaces.I_BrainContext) {
	println("text_function/DumpMemoryText")
	// 	if IsEmptyBrainForText(brainContext) {
	// 		errcode.PrintMsgFromErrorCode(errcode.WARNING_BRAIN_EMPTY, "Text")
	// 		return
	// 	}

	// textData := templates.GetDataFromCell[*data.S_TextData](brainContext.GetFirstCell())
	// fmt.Printf("count : %d\n", textData.Count)
	//
	//	for childCell := textData.LetterCell.GetFirstChildCell(); childCell != nil; childCell = childCell.GetNextCell() {
	//		childData := childCell.GetData()
	//		if childData == nil {
	//			// Return and empty cell
	//			errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA, "text_debug.go")
	//			errcode.PrintCallStack()
	//			continue
	//		}
	//		childData.DumpCell(childCell, []byte{})
	//	}
}
