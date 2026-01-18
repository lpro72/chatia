package interfaces

/*******************
* Interface
*******************/
type I_CellsGroupManagement interface {
	GetCellGroupsCount() uint32
	GetCellsCount() uint32
	GetNextCellID() uint32
	AppendCellToGroup(cell I_Cell) uint32
	GetCellFromID(cellID uint32) I_Cell

	// File Management
	I_File

	// Memory access
	I_Lock
}
