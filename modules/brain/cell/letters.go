//go:build Text

package cell

/*******************
* Import
*******************/
import (
    "os"
    "math/rand"
    "fmt"
    "chatia/modules/error"
)

/*******************
* Types
*******************/
type LetterStruct struct {
    Letter rune
    WordCell *CellStruct
}

/*******************
* GetLetterCell
*******************/
func GetLetterCell(currentCell *CellStruct) *LetterStruct {
    // Must be a letter cell
    letterCell, ok := currentCell.Cell.(*LetterStruct)
    if !ok {
        error.PrintMsgFromErrorCode(error.ERROR_FATAL_BRAIN_INVALID)
        os.Exit(error.ERROR_FATAL_BRAIN_INVALID)
    }
    
    return(letterCell)
}

/*******************
* DumpCell
*******************/
func (letterCell LetterStruct) DumpCell(currentCell *CellStruct, indentation []byte) {
    wordCell := GetWordCell(letterCell.WordCell)
    if wordCell != nil {
        wordCell.DumpCell(currentCell, indentation)
    } else {
        fmt.Printf("%sLetter : %c, count : %d\n", indentation, letterCell.Letter, currentCell.Count)
    }
    for childCell := currentCell.ChildCell; childCell != nil; childCell = childCell.NextCell {
        childCell.Cell.DumpCell(childCell, append(indentation, []byte{' '}...))
    }
}

/*******************
* GetWordFromLastCell
*******************/
func GetWordFromLastCell(currentCell *CellStruct) string {
    var word []rune;
    letterCell := GetLetterCell(currentCell)
    word = append(word, letterCell.Letter)
    for parentCell := currentCell.ParentCell; parentCell != nil; parentCell = parentCell.ParentCell {
        letterCell = GetLetterCell(parentCell)
        word = append([]rune{letterCell.Letter}, word...)
    }
    return(string(word))
}

/*******************
* SearchForLetterCell
*******************/
func SearchForLetterCell(letter rune, currentCell *CellStruct) *CellStruct {
    var lastCell *CellStruct = currentCell
    
    for currentCell != nil {
        letterCell := GetLetterCell(currentCell)
        if letterCell.Letter == letter {
            return currentCell
        }
        lastCell = currentCell
        currentCell = currentCell.NextCell
    }

    // New cell must be created
    newCell := new(CellStruct)
    newletterCell := new(LetterStruct)
    newCell.Cell = newletterCell
    if lastCell != nil {
        lastCell.NextCell = newCell
    }
    newletterCell.Letter = letter
    return(newCell)
}

/*******************
* GetRandowLetterCell
*******************/
func GetRandowLetterCell(currentCell *CellStruct) []rune {
    if currentCell.ChildCell != nil {
        nextCellID := rand.Intn(currentCell.Count) + 1
        for childCell := currentCell.ChildCell; childCell != nil; childCell = childCell.NextCell {
            letterCount := childCell.Count
            wordCount := 0

            letterCell := GetLetterCell(childCell)
    
            if letterCell.WordCell != nil {
                wordCount = letterCell.WordCell.Count
            }
            if nextCellID <= letterCount + wordCount {
                if nextCellID <= wordCount {
                    return []rune{letterCell.Letter}
                }
                childLetters := GetRandowLetterCell(childCell)
                return append([]rune{letterCell.Letter}, childLetters...)
            }
            nextCellID -= letterCount + wordCount
        }
    }
    return nil
}

