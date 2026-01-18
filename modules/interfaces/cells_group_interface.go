package interfaces

/*******************
* Interface
*******************/
type I_CellsGroup interface {
	AppendCellToGroup(cell I_Cell)
	GetCellsCount() uint32
	GetCellFromID(cellID uint32) I_Cell

	// File Management
	I_File

	// Memory access
	I_Lock
}
