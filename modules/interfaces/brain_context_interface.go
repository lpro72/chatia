package interfaces

/*******************
* Interfaces I_BrainContext
*******************/
type I_BrainContext interface {
	Initialize(brainConfig I_BrainConfig, name string, firstCellID uint32)
	GetName() string
	GetFirstSynapseID() uint32
	GetFirstSynapse() I_Synapse
	SetFirstSynapse(firstSynapse I_Synapse)
	GetFirstCellID() uint32
	GetFirstCell() I_Cell
	SetFirstCell(firstCell I_Cell, maxChildListSize uint32)

	SetLearnFunction(learn func(brainContext I_BrainContext, data []byte))
	SetExecFunction(exec func(brainContext I_BrainContext, command string) string)
	SetDumpMemoryFunction(dumpMemory func(brainContext I_BrainContext))
	CallLearnFunction(data []byte)
	CallExecFunction(command string) string
	CallDumpMemoryFunction()

	GetBrainConfig() I_BrainConfig

	SetFileOffset(offset int64)
	GetFileOffset() int64
}
