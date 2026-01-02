package interfaces

/*******************
* Interfaces I_BrainContextManagement
*******************/
type I_BrainContextManagement interface {
	AddNewBrainContext(name string, context I_BrainContext)
	GetBrainContext(name string) I_BrainContext
	CreateNewBrainContext(brain I_BrainConfig, name string, firstSynapseID uint32) I_BrainContext
	UpdateToFile(context I_BrainContext)

	// File Management
	I_File

	// Memory access
	I_Lock
}
