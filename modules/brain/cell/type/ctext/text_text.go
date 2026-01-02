package ctext

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
* CreateTextCellFromSerializeData
*******************/
func CreateTextCellFromSerializeData(dataSerialized []byte) interfaces.I_CellData {
	textData := new(data.S_TextData)
	return textData
}

/*******************
* TextCell_Create
*******************/
func TextCell_Create(brainContext interfaces.I_BrainContext) interfaces.I_Cell {
	newTextData := new(data.S_TextData)
	newCell := data.CreateCell(brainContext.GetBrainConfig(), newTextData, g_TextCellType)
	return (newCell)
}
