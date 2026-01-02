package main

/*******************
* Import
*******************/
import (
	"chatia/modules/brain"

	"chatia/modules/brain/type/bchess"
	"chatia/modules/brain/type/btext"
	"chatia/modules/data"
)

/*******************
* initAll
*******************/
func initAll() {
	data.BrainConfig_Init()
	brain.BrainManagement_Register()
	btext.TextBrainContext_Register()
	bchess.ChessBrainContext_Register()
	data.BrainConfig_Register()
}

/*******************
* closeAll
*******************/
func closeAll() {
	data.BrainConfig_Close()
	data.CellType_Close()
}
