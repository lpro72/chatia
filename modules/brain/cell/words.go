//go:build Text

package cell

/*******************
* Import
*******************/
import (
    "math/rand"
    "fmt"
    "chatia/modules/error"
    "os"
)

/*******************
* Types
*******************/
type WordStruct struct {
    FirstLetterCell *CellStruct
    LastLetterCell *CellStruct
    Word string
}

/*******************
* DumpCell
*******************/
func (wordCell WordStruct) DumpCell(currentCell *CellStruct, indentation []byte) {
    letterCell := GetLetterCell(wordCell.LastLetterCell)
    fmt.Printf("%sLetter : %c, count : %d, word : %s, Word Count : %d\n", indentation, letterCell.Letter, wordCell.LastLetterCell.Count, wordCell.Word, letterCell.WordCell.Count)
}

/*******************
* GetWordCell
*******************/
func GetWordCell(currentCell *CellStruct) *WordStruct {
    // Must be a letter cell
    if currentCell == nil {
        return nil
    }
        
    wordCell, ok := currentCell.Cell.(*WordStruct)
    if !ok {
        error.PrintMsgFromErrorCode(error.ERROR_FATAL_BRAIN_INVALID)
        os.Exit(error.ERROR_FATAL_BRAIN_INVALID)
    }

    return(wordCell)
}

/*******************
* GetRandowWordCell
*******************/
// Get a random cell
func GetRandowWordCell(currentCell *CellStruct, count int) string {
    for nextCellID := rand.Intn(count) + 1; currentCell != nil; currentCell = currentCell.NextCell {
        if nextCellID <= currentCell.Count {
            return GetWordCell(currentCell).Word
        }
        nextCellID -= currentCell.Count
    }
    return ""
}

