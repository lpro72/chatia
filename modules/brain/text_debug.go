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

/*******************
* UnittestText
*******************/
func UnittestText() {
    LearnFromString("This is a test", "Text")
    LearnFromString("Ceci est un test", "Text")
    LearnFromString("The brain work", "Text")
    LearnFromString("Test with the word these", "Text")
    DumpMemory("Text")

    for i := 0; i < 10; i++ {
        fmt.Printf("Test %d\n-------\n", i)
        fmt.Println(Exec("GetRandomWordFromLetterCell", "Text"))
        fmt.Println(Exec("GetRandomWordFromWordCell", "Text"))
    }
}


