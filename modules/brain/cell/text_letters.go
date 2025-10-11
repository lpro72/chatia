//go:build !NoText

package cell

/*******************
* Import
*******************/
import (
	"math/rand"
)

/*******************
* GetWordFromLastCell
*******************/
func GetWordFromLastCell(lastCell I_CellManagement) string {
	var word []rune
	letterData := GetDataFromLetterCell(lastCell)
	word = append(word, letterData.Letter)
	for parentCell := lastCell.GetParentCell(); parentCell != nil; parentCell = parentCell.GetParentCell() {
		// The last parent cell is not a valid letter cell
		if parentCell.GetParentCell() == nil {
			break
		}
		letterData = GetDataFromLetterCell(parentCell)
		word = append([]rune{letterData.Letter}, word...)
	}
	return (string(word))
}

/*******************
* SearchForLetterCell
*******************/
func SearchForLetterCell(letter rune, parentCell I_CellManagement) I_CellManagement {
	var currentCell I_CellManagement = parentCell.GetFirstChildCell()

	for currentCell != nil {
		letterData := GetDataFromLetterCell(currentCell)
		if letterData.Letter == letter {
			return currentCell
		}
		currentCell = currentCell.GetNextCell()
	}

	// New cell must be created
	return (CreateLetterCell(letter, parentCell))
}

/*******************
* GetRandowWordFromLetterCell
*******************/
func GetRandowWordFromLetterCell(currentCell I_CellManagement, count int) []rune {
	if currentCell.GetFirstChildCell() != nil {
		nextCellID := rand.Intn(count) + 1
		for childCell := currentCell.GetFirstChildCell(); childCell != nil; childCell = childCell.GetNextCell() {
			letterData := GetDataFromLetterCell(childCell)
			letterCount := letterData.Count
			wordCount := 0

			if letterData.WordCellID != 0 {
				wordCell := GetCellFromGroup(letterData.WordCellID)
				wordCellData := GetDataFromWordCell(wordCell)
				wordCount = wordCellData.Count
			}

			if nextCellID <= letterCount+wordCount {
				if nextCellID <= wordCount {
					return []rune{letterData.Letter}
				}
				childLetters := GetRandowWordFromLetterCell(childCell, letterData.Count)
				return append([]rune{letterData.Letter}, childLetters...)
			}
			nextCellID -= letterCount + wordCount
		}
	}
	return make([]rune, 1)
}
