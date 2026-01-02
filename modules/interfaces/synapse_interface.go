package interfaces

/*******************
* Interface
*******************/
type I_Synapse interface {
	GetID() uint32
	GetParent() I_Synapse
	GetNext() I_Synapse
	GetPrevious() I_Synapse
	GetFirstChild() I_Synapse
	GetCell() I_Cell
}
