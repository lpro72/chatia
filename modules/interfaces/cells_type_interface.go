package interfaces

type I_CellTypeManagement interface {
	GetCellTypeID(name string) uint32
	GetCellNextTypeID() uint32
	GetCellTypeName(id uint32) string
	CreateCellDataFromSerializedData(typeID uint32, serializedData []byte) I_CellData
	AddCellType(name string, CreateCellFromSerializeData func([]byte) I_CellData) uint32

	// File Management
	I_File

	// Memory access
	I_Lock
}
