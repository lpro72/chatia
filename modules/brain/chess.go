//go:build !NoChess

package brain

/*******************
* Import
*******************/
import (
	"chatia/modules/brain/cell"
	"fmt"
)

/*******************
* Types
*******************/
type ChessStruct struct {
}

/*******************
* LearnChess
*******************/
func LearnChess(data []byte, firstCell cell.I_CellManagement) {
	fmt.Println("***** Enter LearnChess *****")

	fmt.Println("***** Exit LearnChess *****")
}

/*******************
* ChessFactory
*******************/
func ChessFactory(brain I_Brain) {
	fmt.Println("***** Enter ChessFactory *****")

	//brain.SetFirstCell(new(ChessStruct))
	brain.SetLearnFunction(LearnChess)
	brain.SetDumpMemoryFunction(DumpMemoryChess)
	brain.SetExecFunction(ExecChess)
	brain.SetUnittestFunction(UnittestChess)

	fmt.Println("***** Exit ChessFactory *****")
}

/*******************
* init
*******************/
func init() {
	fmt.Println("***** Enter init *****")

	registerBrainContext("Chess", ChessFactory)

	fmt.Println("***** Exit init *****")
}

/*******************
* ExecText
*******************/
// Learn some texte
func ExecChess(command string, extraVar ...any) string {
	fmt.Println("***** Enter ExecChess *****")

	fmt.Println(command)
	for index, value := range extraVar {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
	fmt.Println("***** Exit ExecChess *****")

	return ""
}

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
