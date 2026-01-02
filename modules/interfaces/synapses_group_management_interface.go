package interfaces

/*******************
* Interface
*******************/
type I_SynapsesGroupManagement interface {
	GetSynapsesGroupsCount() int
	GetSynapsesCount(synapseID int) int
	GetNextSynapseID() uint32
	AppendSynapseToGroup(synapse I_Synapse) uint32
	GetSynapseFromID(synapseID uint32) I_Synapse

	// File Management
	I_File

	// Memory access
	I_Lock
}
