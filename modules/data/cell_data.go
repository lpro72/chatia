package data

/*******************
* Import
*******************/
import (
	"chatia/modules/interfaces"
)

/*******************
* Types
*******************/
type S_Cell struct {
	id          uint32
	cellType    uint32
	cellData    interfaces.I_CellData
	brainConfig interfaces.I_BrainConfig
}

/*******************
* Interface I_Cell
*******************/
func (currentCell *S_Cell) GetID() uint32 {
	return currentCell.id
}

func (currentCell *S_Cell) GetCellType() uint32 {
	return currentCell.cellType
}

func (currentCell *S_Cell) GetSerializedData() []byte {
	if currentCell.cellData == nil {
		return []byte{}
	}

	return currentCell.cellData.GetSerializedData()
}

func (currentCell *S_Cell) GetData() interfaces.I_CellData {
	return currentCell.cellData
}

func (currentCell *S_Cell) DumpCell() {
	cellData := currentCell.GetData()
	if cellData != nil {
		cellData.DumpData()
	}
}

/*******************
* CreateCell
*******************/
func CreateCell(brainConfig interfaces.I_BrainConfig, currentData interfaces.I_CellData, cellType uint32) interfaces.I_Cell {
	newCell := new(S_Cell)
	newCell.cellType = cellType
	newCell.cellData = currentData
	newCell.brainConfig = brainConfig
	newCell.id = brainConfig.GetCellsGroupManagament().GetNextCellID()
	brainConfig.GetCellsGroupManagament().AppendCellToGroup(newCell)

	return (newCell)
}
