//go:build Text

package brain

/*******************
* Import
*******************/
import (
    "unicode"
    "math/rand"
    "time"
    "chatia/modules/error"
    "chatia/modules/brain/cell"
)

/*******************
* Types
*******************/
type TextStruct struct {
    letterCell *cell.CellStruct
    wordCell *cell.CellStruct
}

/*******************
* LearnTextFromBrain
*******************/
func LearnTextFromBrain(data []byte, firstCell any) {
    LearnText(string(data), firstCell.(*TextStruct))
}

/*******************
* TextFactory
*******************/
func TextFactory(brain *BrainStruct)  {
    brain.FirstCell = new(TextStruct)
    brain.Learn = LearnTextFromBrain
    brain.DumpMemory = DumpMemoryText
    brain.Exec = ExecText
}

/*******************
* init
*******************/
func init() {
    AddBrainFactory("Text", TextFactory)
    CreateBrainContext("Text")
}

/*******************
* ExecText
*******************/
// Learn some texte
func ExecText(command string, extraVar ...any) string {
    if command == "GetRandomWordFromWordCell" {
        return GetRandomWordFromWordCell()
    } else if command == "GetRandomWordFromLetterCell" {
        return GetRandomWordFromLetterCell()
    }
    
    error.PrintMsgFromErrorCode(error.WARNING_COMMAND_NOT_FOUND, command)
    
    return ""
}

/*******************
* LearnText
*******************/
// Learn some texte
func LearnText(text string, textCell *TextStruct) {
    var currentCell *cell.CellStruct = nil
    var firstCell *cell.CellStruct = nil
    for _, r := range text {
        // New word
        if ! unicode.IsLetter(r) {
            // Mark as end of a word
            if currentCell != nil {
                letterCell := cell.GetLetterCell(currentCell)
                wordCell := cell.GetWordCell(letterCell.WordCell)
                if wordCell == nil {
                    letterCell.WordCell = new(cell.CellStruct)
                    letterCell.WordCell.Cell = new(cell.WordStruct)
                    if textCell.wordCell == nil {
                        textCell.wordCell = letterCell.WordCell
                    } else {
                        letterCell.WordCell.NextCell = textCell.wordCell
                        textCell.wordCell = letterCell.WordCell
                    }
                    wordCell = cell.GetWordCell(letterCell.WordCell)
                    wordCell.LastLetterCell = currentCell
                    wordCell.FirstLetterCell = firstCell
                    wordCell.Word = cell.GetWordFromLastCell(currentCell)
                }
                letterCell.WordCell.Count += 1
            }
            currentCell = nil
            firstCell = nil
            continue
        }

        // New word
        if currentCell == nil {
            currentCell = textCell.letterCell
            if currentCell != nil {
                currentCell.Count += 1
            }
        }

        // special case of a new brain
        if currentCell == nil {
            currentCell = new(cell.CellStruct)
            currentCell.Cell = new(cell.LetterStruct)
            textCell.letterCell = currentCell
            currentCell.Count += 1
        }
        
        childCell := cell.SearchForLetterCell(unicode.ToLower(r), currentCell.ChildCell)
        if childCell.ParentCell == nil {
            childCell.ParentCell = currentCell
        }
        childCell.Count += 1
        if firstCell == nil {
            firstCell = childCell
        }
        if currentCell.ChildCell == nil {
            currentCell.ChildCell = childCell
        }

        currentCell = childCell
    }

    // Last word
    if currentCell != nil {
        letterCell := cell.GetLetterCell(currentCell)
        wordCell := cell.GetWordCell(letterCell.WordCell)
        if wordCell == nil {
            letterCell.WordCell = new(cell.CellStruct)
            wordCell = new(cell.WordStruct)
            letterCell.WordCell.Cell = wordCell
            wordCell.LastLetterCell = currentCell
            wordCell.FirstLetterCell = firstCell
            wordCell.Word = cell.GetWordFromLastCell(currentCell)
        } 
        letterCell.WordCell.Count += 1
    }
    
}

/*******************
* IsEmptyBrainForText
*******************/
func IsEmptyBrainForText() bool {
    brain := GetBrainContext("Text")
    textCell := brain.FirstCell.(*TextStruct)
    if textCell == nil || textCell.letterCell == nil || textCell.letterCell.ChildCell == nil {
        return true
    }
    return false
}

/*******************
* GetRandomWordFromWordCell
*******************/
func GetRandomWordFromWordCell() string {
    if IsEmptyBrainForText() {
        error.PrintMsgFromErrorCode(error.WARNING_BRAIN_EMPTY, "Text")
        return ""
    }
    
    brain := GetBrainContext("Text")
    textCell := brain.FirstCell.(*TextStruct)
    rand.Seed(time.Now().UnixNano())
    return cell.GetRandowWordCell(textCell.wordCell, textCell.letterCell.Count)
}

/*******************
* GetRandomWordFromLetterCell
*******************/
func GetRandomWordFromLetterCell() string {
    if IsEmptyBrainForText() {
        error.PrintMsgFromErrorCode(error.WARNING_BRAIN_EMPTY, "Text")
        return ""
    }

    brain := GetBrainContext("Text")
    textCell := brain.FirstCell.(*TextStruct)
    rand.Seed(time.Now().UnixNano())
    return string(cell.GetRandowLetterCell(textCell.letterCell))
}


