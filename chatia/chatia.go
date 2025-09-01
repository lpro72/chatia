package main

/*******************
* Import
*******************/
import (
        "os"
        "fmt"
        "chatia/modules/error"
        "chatia/modules/brain"
)

/*******************
* main
*******************/
func main() {
    brain.DumpMemory("Text")
    brain.LearnFromString("This is a test", "Text")
    brain.LearnFromString("Ceci est un test", "Text")
    brain.LearnFromString("The brain work", "Text")
    brain.LearnFromString("Test with the word these", "Text")
    brain.DumpMemory("Text")
    for i := 0; i < 10; i++ {
        fmt.Printf("Test %d\n-------\n", i)
        fmt.Println(brain.Exec("GetRandomWordFromLetterCell", "Text"))
        fmt.Println(brain.Exec("GetRandomWordFromWordCell", "Text"))
    }
    os.Exit(error.SUCCESS)
}

