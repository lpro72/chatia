//go:build !NoText

package brain

/*******************
* Import
*******************/
import (
	"chatia/modules/brain/cell"
)

/*******************
* TextFactory
*******************/
func TextFactory(brain I_Brain) {
	brain.SetFirstCell(cell.CreateTextCell())
	brain.SetLearnFunction(LearnTextFromBrain)
	brain.SetDumpMemoryFunction(DumpMemoryText)
	brain.SetExecFunction(ExecText)
	brain.SetUnittestFunction(UnittestText)
}

/*******************
* init
*******************/
func init() {
	registerBrainContext("Text", TextFactory)
}
