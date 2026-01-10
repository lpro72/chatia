package text

/*******************
* Import
*******************/
import (
	"fmt"
	"unicode"

	"chatia/modules/data"
	"chatia/modules/errcode"
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
	cellsGroupManagement := brainConfig.GetCellsGroupManagament()
	synapsesGroupManagement := brainConfig.GetSynapsesGroupManagement()
	textCell := brainContext.GetFirstCell()
	var currentCell interfaces.I_Cell = nil
	var currentSynapse interfaces.I_Synapse = nil
	var firstSynapse interfaces.I_Synapse = nil
	var letterData *S_LetterCellData = nil
	var wordCell interfaces.I_Cell = nil

	textData := templates.GetDataFromCell[*S_TextData](textCell)
	wordSynapse := synapsesGroupManagement.GetSynapseFromID(textData.WordSynapseID)
	letterSynapse := synapsesGroupManagement.GetSynapseFromID(textData.LetterSynapseID)
	// initialisation
	if wordSynapse == nil {
		wordSynapse = data.CreateSynapse(brainConfig, nil, nil, 1)
		textData.WordSynapseID = wordSynapse.GetID()
	}
	if letterSynapse == nil {
		letterSynapse = data.CreateSynapse(brainConfig, nil, nil, 26)
		textData.LetterSynapseID = letterSynapse.GetID()
	}

	for _, r := range text {
		// New word
		if !unicode.IsLetter(r) {
			// Mark as end of a word
			if currentCell != nil {
				letterData = templates.GetDataFromCell[*S_LetterCellData](currentCell)
				if letterData.WordCellID == 0 {
					wordCell = WordCell_Create(brainConfig, firstSynapse, currentSynapse)
					letterData.WordCellID = wordCell.GetID()
					data.CreateSynapse(brainConfig, wordSynapse, wordCell, 1)
				} else {
					wordCell = cellsGroupManagement.GetCellFromID(letterData.WordCellID)
				}
				wordData := templates.GetDataFromCell[*S_WordCellData](wordCell)
				wordData.Count += 1
			}
			currentCell = nil
			firstSynapse = nil
			continue
		}

		// New word
		if currentSynapse == nil {
			textData.Count += 1
			currentSynapse = letterSynapse
		}

		currentSynapse, currentCell := LetterCell_Search(brainConfig, unicode.ToLower(r), currentSynapse)
		letterData = templates.GetDataFromCell[*S_LetterCellData](currentCell)
		letterData.Count += 1
		if firstSynapse == nil {
			firstSynapse = currentSynapse
		}
	}

	// Last word
	if currentCell != nil {
		letterData = templates.GetDataFromCell[*S_LetterCellData](currentCell)
		if letterData.WordCellID == 0 {
			wordCell = WordCell_Create(brainConfig, firstSynapse, currentSynapse)
			letterData.WordCellID = wordCell.GetID()
			data.CreateSynapse(brainConfig, wordSynapse, wordCell, 1)
		} else {
			wordCell = brainConfig.GetCellsGroupManagament().GetCellFromID(letterData.WordCellID)
		}
		wordData := templates.GetDataFromCell[*S_WordCellData](wordCell)
		wordData.Count += 1
	}
}

/*******************
* IsEmptyBrainForText
*******************/
func IsEmptyBrainForText(brainContext interfaces.I_BrainContext) (*S_TextData, bool) {
	textData := templates.GetDataFromCell[*S_TextData](brainContext.GetFirstCell())
	if textData == nil || textData.LetterSynapseID == 0 || textData.WordSynapseID == 0 {
		return nil, true
	}
	return textData, false
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
	textData, empty := IsEmptyBrainForText(brainContext)
	if empty {
		errcode.PrintMsgFromErrorCode(errcode.WARNING_BRAIN_EMPTY, "Text")
		return
	}

	fmt.Printf("count : %d\n", textData.Count)
	if textData.LetterSynapseID != 0 {
		letterSynapse := brainContext.GetBrainConfig().GetSynapsesGroupManagement().GetSynapseFromID(textData.LetterSynapseID)
		for childSynapse := letterSynapse.GetNext(); childSynapse != nil; childSynapse = childSynapse.GetNext() {
			childData := templates.GetDataFromCell[*S_LetterCellData](childSynapse.GetCell())
			if childData == nil {
				fmt.Printf("Empty cell\n")
				// Return and empty cell
				errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA, "text_function.go")
				continue
			}
			fmt.Printf("Letter: %c, Count: %d, WordCellID: %d\n", childData.Letter, childData.Count, childData.WordCellID)
			childData.DumpCell(childSynapse.GetCell(), []byte{})
		}
	}

	if textData.WordSynapseID != 0 {
		wordSynapse := brainContext.GetBrainConfig().GetSynapsesGroupManagement().GetSynapseFromID(textData.WordSynapseID)
		for childSynapse := wordSynapse.GetNext(); childSynapse != nil; childSynapse = childSynapse.GetNext() {
			childData := templates.GetDataFromCell[*S_WordCellData](childSynapse.GetCell())
			if childData == nil {
				fmt.Printf("Empty cell\n")
				// Return and empty cell
				errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_INVALID_DATA, "text_function.go")
				continue
			}
			fmt.Printf("Word Count: %d\n", childData.Count)
			childData.DumpCell(childSynapse.GetCell(), []byte{})
		}
	}
}
