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
    fmt.Println("***** Enter main *****")
//    brain.LearnFromString("This is a test", "Text")
//    brain.LearnFromString("Ceci est un test", "Text")
//    brain.LearnFromString("The brain work", "Text")
//    brain.LearnFromString("Test with the word these", "Text")
//    brain.DumpMemory("Text")

    brain.LearnFromString("1.e4 e5 2.Nf3 Nc6", "Chess")
    brain.DumpMemory("Chess")
    for i := 0; i < 10; i++ {
        fmt.Printf("Test %d\n-------\n", i)
//        fmt.Println(brain.Exec("GetRandomWordFromLetterCell", "Text"))
//        fmt.Println(brain.Exec("GetRandomWordFromWordCell", "Text"))

        brain.Exec("test", "Chess")
    }

    fmt.Println("***** Exit main *****")

    os.Exit(error.SUCCESS)
}

