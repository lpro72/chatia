//go:build !NoText

package cell

/*******************
* Import
*******************/
import (
	"math/rand"
)

/*******************
* GetRandowWordFromWordCells
*******************/
func GetRandowWordFromWordCells(currentCell I_CellManagement, count int) string {
	for nextCellID := rand.Intn(count) + 1; currentCell != nil; currentCell = currentCell.GetNextCell() {
		wordData := GetDataFromWordCell(currentCell)
		if nextCellID <= wordData.Count {
			return wordData.Word
		}
		nextCellID -= wordData.Count
	}
	return ""
}
