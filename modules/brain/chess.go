//go:build Chess

package brain

/*******************
* Import
*******************/
import (
    "chatia/modules/brain/cell"
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
    LearnText(string(data), firstCell.(*TextStruct))
}

/*******************
* ChessFactory
*******************/
func ChessFactory(brain *BrainStruct)  {
    brain.FirstCell = new(ChessStruct)
    brain.Learn = LearnChess
    brain.DumpMemory = DumpMemoryChess
    brain.Exec = ExecChess
}

/*******************
* init
*******************/
func init() {
    AddBrainFactory("Chess", TextFactory)
    CreateBrainContext("Text")
}

/*******************
* ExecText
*******************/
// Learn some texte
func ExecChess(command string, extraVar ...any) string {
    return ""
}


