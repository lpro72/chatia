package interfaces

/*******************
* Interface
*******************/
type I_SynapsesGroup interface {
	AppendSynapseToGroup(synapse I_Synapse)
	GetSynapsesCount() uint32
	GetSynapseFromID(synapseID uint32) I_Synapse

	// File Management
	I_File

	// Memory access
	I_Lock
}
