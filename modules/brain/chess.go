//go:build Chess

package brain

/*******************
* Import
*******************/
import (
    "fmt"
//    "chatia/modules/brain/cell"
)

/*******************
* Types
*******************/
type ChessStruct struct {

}

/*******************
* LearnChess
*******************/
func LearnChess(data []byte, firstCell any) {
    fmt.Println("***** Enter LearnChess *****")

    fmt.Println("***** Exit LearnChess *****")
}

/*******************
* ChessFactory
*******************/
func ChessFactory(brain *BrainStruct)  {
    fmt.Println("***** Enter ChessFactory *****")

    brain.FirstCell = new(ChessStruct)
    brain.Learn = LearnChess
    brain.DumpMemory = DumpMemoryChess
    brain.Exec = ExecChess

    fmt.Println("***** Exit ChessFactory *****")
}

/*******************
* init
*******************/
func init() {
    fmt.Println("***** Enter init *****")

    AddBrainFactory("Chess", ChessFactory)
    CreateBrainContext("Chess")

    fmt.Println("***** Exit init *****")
}

/*******************
* ExecText
*******************/
// Learn some texte
func ExecChess(command string, extraVar ...any) string {
    fmt.Println("***** Enter ExecChess *****")

    fmt.Println("***** Exit ExecChess *****")

    return ""
}


