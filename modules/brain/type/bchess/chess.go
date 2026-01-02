package bchess

/*******************
* Import
*******************/
import (
	"fmt"

	"chatia/modules/data"
	"chatia/modules/interfaces"
)

/*******************
* Types
*******************/
type ChessStruct struct {
}

/*******************
* ChessFactory
*******************/
func ChessFactory(brainContext interfaces.I_BrainContext) {
	brainContext.SetDumpMemoryFunction(DumpMemoryChess)
	brainContext.SetLearnFunction(LearnChess)
	brainContext.SetExecFunction(ExecChess)
}

/*******************
* ChessBrainContext_Register
*******************/
func ChessBrainContext_Register() {
	data.BrainContextManagement_RegisterNewContext("Chess", ChessFactory)
}

/*******************
* DumpMemoryChess
*******************/
func DumpMemoryChess(brainContext interfaces.I_BrainContext) {
	fmt.Println("DumpMemoryChess")
}

/*******************
* LearnChess
*******************/
func LearnChess(brainContext interfaces.I_BrainContext, data []byte) {
	fmt.Println("LearnChess")
}

/*******************
* ExecText
*******************/
func ExecChess(brainContext interfaces.I_BrainContext, command string) string {
	fmt.Println(command)

	return ""
}
