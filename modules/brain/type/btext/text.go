package btext

/*******************
* Import
*******************/
import (
	"chatia/modules/brain/cell/type/ctext"
	"chatia/modules/data"
	"chatia/modules/interfaces"
)

/*******************
* textBrainContext_Factory
*******************/
func textBrainContext_Factory(brainContext interfaces.I_BrainContext) {
	if brainContext.GetFirstCellID() == 0 {
		brainContext.SetFirstCell(ctext.TextCell_Create(brainContext))
	}
	brainContext.SetLearnFunction(LearnTextFromBrain)
	brainContext.SetDumpMemoryFunction(DumpMemoryText)
	brainContext.SetExecFunction(ExecText)
}

/*******************
* TextBrainContext_Register
*******************/
func TextBrainContext_Register() {
	data.BrainContextManagement_RegisterNewContext("Text", textBrainContext_Factory)
	ctext.TextCell_Register()
	ctext.LetterCell_Register()
	ctext.WordCell_Register()

}
