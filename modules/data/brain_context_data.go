package data

/*******************
* Import
*******************/
import (
	"chatia/modules/interfaces"
)

/*******************
* S_BrainContext
*******************/
type S_BrainContext struct {
	learn func(brainContext interfaces.I_BrainContext, data []byte)
	exec  func(brainContext interfaces.I_BrainContext, command string) string

	// Data
	firstSynapseID uint32
	brainConfig    interfaces.I_BrainConfig
	name           string

	// file
	dataOffset int64

	// Debug only
	dumpMemory func(brainContext interfaces.I_BrainContext)
}

/*******************
*  Functions for the interface I_BrainContext
*******************/
func (brainContext *S_BrainContext) Initialize(brainConfig interfaces.I_BrainConfig, name string, firstSynapseID uint32) {
	brainContext.name = name
	brainContext.brainConfig = brainConfig
	brainContext.firstSynapseID = firstSynapseID
}

func (brainContext *S_BrainContext) GetName() string {
	return brainContext.name
}

func (brainContext *S_BrainContext) GetFirstSynapseID() uint32 {
	return brainContext.firstSynapseID
}

func (brainContext *S_BrainContext) GetFirstSynapse() interfaces.I_Synapse {
	return brainContext.brainConfig.GetSynapsesGroupManagement().GetSynapseFromID(brainContext.firstSynapseID)
}

func (brainContext *S_BrainContext) SetFirstSynapse(firstSynapse interfaces.I_Synapse) {
	brainContext.firstSynapseID = firstSynapse.GetID()
	brainContext.brainConfig.GetBrainContextManagement().UpdateToFile(brainContext)
}

func (brainContext *S_BrainContext) GetFirstCellID() uint32 {
	cell := brainContext.GetFirstCell()
	if cell == nil {
		return 0
	}
	return cell.GetID()
}

func (brainContext *S_BrainContext) GetFirstCell() interfaces.I_Cell {
	synapse := brainContext.GetFirstSynapse()
	if synapse == nil {
		return nil
	}
	cell := synapse.GetCell()
	if cell == nil {
		return nil
	}
	return cell
}

func (brainContext *S_BrainContext) SetFirstCell(firstCell interfaces.I_Cell) {
	synapse := brainContext.GetFirstSynapse()
	if synapse == nil {
		synapse = CreateSynapse(brainContext.brainConfig, nil, firstCell)
		brainContext.SetFirstSynapse(synapse)
		return
	}
	print("brain_context_data/SetFirstCell : Need to update data")
}

func (brainContext *S_BrainContext) SetDumpMemoryFunction(dumpMemory func(brainContext interfaces.I_BrainContext)) {
	brainContext.dumpMemory = dumpMemory
}

func (brainContext *S_BrainContext) SetLearnFunction(learn func(brainContext interfaces.I_BrainContext, data []byte)) {
	brainContext.learn = learn
}

func (brainContext *S_BrainContext) SetExecFunction(exec func(brainContext interfaces.I_BrainContext, command string) string) {
	brainContext.exec = exec
}

func (brainContext *S_BrainContext) CallLearnFunction(data []byte) {
	if brainContext.learn != nil {
		brainContext.learn(brainContext, data)
	}
}

func (brainContext *S_BrainContext) CallExecFunction(command string) string {
	if brainContext.exec != nil {
		return brainContext.exec(brainContext, command)
	}

	return ""
}

func (brainContext *S_BrainContext) CallDumpMemoryFunction() {
	if brainContext.dumpMemory != nil {
		brainContext.dumpMemory(brainContext)
	}
}

func (brainContext *S_BrainContext) GetBrainConfig() interfaces.I_BrainConfig {
	return brainContext.brainConfig
}

func (brainContext *S_BrainContext) SetFileOffset(offset int64) {
	brainContext.dataOffset = offset
}

func (brainContext *S_BrainContext) GetFileOffset() int64 {
	return brainContext.dataOffset
}
