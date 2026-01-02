package interfaces

/*******************
* Interface
*******************/
type I_Cell interface {
	GetData() I_CellData
	GetSerializedData() []byte
	// GetBrain() I_BrainConfig
	GetID() uint32
	GetCellType() uint32
	// DumpCell(indentation []byte)
}
