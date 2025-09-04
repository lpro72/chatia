//go:build Chess && debug

package brain

/*******************
* Import
*******************/
import (
    "fmt"
)

/*******************
* DumpMemoryChess
*******************/
func DumpMemoryChess() {
}

/*******************
* UnittestChess
*******************/
func UnittestChess() {
    LearnFromString("1.e4 e5 2.Nf3 Nc6", "Chess")
    DumpMemory("Chess")
    for i := 0; i < 10; i++ {
        fmt.Printf("Test %d\n-------\n", i)
        Exec("test", "Chess", "var 1", 42, "var 3")
    }

}


