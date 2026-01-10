package text

/*******************
* Import
*******************/
import (
	// 	"math/rand"

	"chatia/modules/data"
	"chatia/modules/interfaces"
	"chatia/modules/templates"
)

/*******************
* Globals Varables
*******************/
var g_LetterCellType uint32 = 0

/*******************
* LetterCell_Register
*******************/
func LetterCell_Register() {
	g_LetterCellType = data.CellType_RegisterNewType("Letter", CreateLetterCellFromSerializeData)
}

/*******************
* LetterCell_GetWordFromLastSynapse
*******************/
func LetterCell_GetWordFromLastSynapse(lastSynapse interfaces.I_Synapse) string {
	var word []rune
	letterData := templates.GetDataFromCell[*S_LetterCellData](lastSynapse.GetCell())
	word = append(word, letterData.Letter)
	for parentSynapse := lastSynapse.GetParent(); parentSynapse != nil; parentSynapse = parentSynapse.GetParent() {
		// The last parent cell is not a valid letter cell
		if parentSynapse.GetParent() == nil {
			break
		}
		letterData = templates.GetDataFromCell[*S_LetterCellData](parentSynapse.GetCell())
		word = append([]rune{letterData.Letter}, word...)
	}
	return (string(word))
}

/*******************
* LetterCell_Search
*******************/
func LetterCell_Search(brainConfig interfaces.I_BrainConfig, letter rune, parentSynapse interfaces.I_Synapse) (interfaces.I_Synapse, interfaces.I_Cell) {

	currentSynapse := parentSynapse.GetFirstChild()

	for currentSynapse != nil {
		currentCell := currentSynapse.GetCell()
		letterData := templates.GetDataFromCell[*S_LetterCellData](currentCell)
		if letterData.Letter == letter {
			return currentSynapse, currentCell
		}
		currentSynapse = currentSynapse.GetNext()
	}

	// New cell must be created
	currentCell := LetterCell_Create(brainConfig, letter)
	println("Created new letter cell for letter:", string(letter))
	return data.CreateSynapse(brainConfig, parentSynapse, currentCell, 26), currentCell
}

/*******************
* LetterCell_GetRandowWord
*******************/
func LetterCell_GetRandowWord(currentCell interfaces.I_Cell, count int) []rune {
	println("text_letters/LetterCell_GetRandowWord")
	// 	if currentCell.GetFirstChildCell() != nil {
	// 		nextCellID := rand.Intn(count) + 1
	// 		for childCell := currentCell.GetFirstChildCell(); childCell != nil; childCell = childCell.GetNextCell() {
	// 			letterData := templates.GetDataFromCell[*data.S_LetterCellData](childCell)
	// 			letterCount := letterData.Count
	// 			wordCount := 0

	// 			if letterData.WordCellID != 0 {
	// 				wordCell := currentCell.GetBrain().GetCellFromID(letterData.WordCellID)
	// 				wordCellData := templates.GetDataFromCell[*data.S_WordCellData](wordCell)
	// 				wordCount = wordCellData.Count
	// 			}

	// 			if nextCellID <= letterCount+wordCount {
	// 				if nextCellID <= wordCount {
	// 					return []rune{letterData.Letter}
	// 				}
	// 				childLetters := LetterCell_GetRandowWord(childCell, letterData.Count)
	// 				return append([]rune{letterData.Letter}, childLetters...)
	// 			}
	// 			nextCellID -= letterCount + wordCount
	// 		}
	// 	}
	return make([]rune, 1)
}
