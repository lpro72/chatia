package text

/*******************
* Import
*******************/
import (
	"chatia/modules/data"
	"chatia/modules/interfaces"
)

/*******************
* Globals Varables
*******************/
var g_TextCellType uint32 = 0

/*******************
* TextCell_Register
*******************/
func TextCell_Register() {
	g_TextCellType = data.CellType_RegisterNewType("Text", CreateTextCellFromSerializeData)
}

/*******************
* textBrainContext_Factory
*******************/
func textBrainContext_Factory(brainContext interfaces.I_BrainContext) {
	if brainContext.GetFirstCellID() == 0 {
		brainContext.SetFirstCell(TextCell_Create(brainContext))
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
	TextCell_Register()
	LetterCell_Register()
	WordCell_Register()

}
