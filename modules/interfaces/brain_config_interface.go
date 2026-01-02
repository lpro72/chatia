package interfaces

/*******************
* Imports
*******************/

/*******************
* Interfaces
*******************/
type I_BrainConfig interface {
	CallFactoryContext()

	// Getters / Setters
	SetMainDirectory(directory string)
	GetMainDirectory() string
	SetSaveDirectory(directory string)
	GetSaveDirectory() string

	// Brain management
	GetBrainContextManagement() I_BrainContextManagement
	GetCellsGroupManagament() I_CellsGroupManagement
	GetSynapsesGroupManagement() I_SynapsesGroupManagement

	// Locking
	I_Lock

	// File management
	I_File
}
