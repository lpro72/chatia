//go:build Text && debug

package brain

/*******************
* Import
*******************/
import (
    "chatia/modules/error"
    "fmt"
    "chatia/modules/brain/cell"
)

/*******************
* DumpMemoryText
*******************/
func DumpMemoryText() {
    if IsEmptyBrainForText() {
        error.PrintMsgFromErrorCode(error.WARNING_BRAIN_EMPTY, "Text")
        return
    }
    
    brain := GetBrainContext("Text")
    textCell := brain.FirstCell.(*TextStruct)
    fmt.Printf("count : %d\n", textCell.letterCell.Count)
    for childCell := textCell.letterCell.ChildCell; childCell != nil; childCell = childCell.NextCell {
        cell.DumpCell(childCell, []byte{})
    }
}


