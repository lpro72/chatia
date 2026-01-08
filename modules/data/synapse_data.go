package data

/*******************
* Import
*******************/
import (
	"chatia/modules/interfaces"
)

/*******************
* Types
*******************/
type S_Synapse struct {
	brainConfig interfaces.I_BrainConfig

	synapseID uint32
	cellID    uint32
	score     uint32

	nextSynapseID      uint32
	previousSynapseID  uint32
	childSynapseIDList []uint32
	parentSynapseID    uint32
}

/*******************
* Interface I_Synapse
*******************/
func (currentSynapse *S_Synapse) GetID() uint32 {
	return currentSynapse.synapseID
}

func (currentSynapse *S_Synapse) GetParent() interfaces.I_Synapse {
	return currentSynapse.brainConfig.GetSynapsesGroupManagement().GetSynapseFromID(currentSynapse.parentSynapseID)
}

func (currentSynapse *S_Synapse) GetNext() interfaces.I_Synapse {
	return currentSynapse.brainConfig.GetSynapsesGroupManagement().GetSynapseFromID(currentSynapse.nextSynapseID)
}

func (currentSynapse *S_Synapse) GetPrevious() interfaces.I_Synapse {
	return currentSynapse.brainConfig.GetSynapsesGroupManagement().GetSynapseFromID(currentSynapse.previousSynapseID)
}

func (currentSynapse *S_Synapse) GetFirstChild() interfaces.I_Synapse {
	if currentSynapse.childSynapseIDList == nil {
		return nil
	}
	return currentSynapse.brainConfig.GetSynapsesGroupManagement().GetSynapseFromID(currentSynapse.childSynapseIDList[0])
}

func (currentSynapse *S_Synapse) addChildSynapse(newSynapse *S_Synapse) {
	var childSynapse interfaces.I_Synapse = nil
	if currentSynapse.childSynapseIDList == nil {
		currentSynapse.childSynapseIDList = make([]uint32, 0, 10)
	} else {
		id := len(currentSynapse.childSynapseIDList) - 1
		childSynapse = currentSynapse.brainConfig.GetSynapsesGroupManagement().GetSynapseFromID(currentSynapse.childSynapseIDList[id])
	}
	if childSynapse != nil {
		if concreteSynapse, ok := childSynapse.(*S_Synapse); ok {
			concreteSynapse.nextSynapseID = newSynapse.GetID()
		}
		newSynapse.previousSynapseID = childSynapse.GetID()
	}
	currentSynapse.childSynapseIDList = append(currentSynapse.childSynapseIDList, newSynapse.GetID())
	newSynapse.parentSynapseID = currentSynapse.GetID()
}

func (currentSynapse *S_Synapse) GetCell() interfaces.I_Cell {
	return currentSynapse.brainConfig.GetCellsGroupManagament().GetCellFromID(currentSynapse.cellID)
}

func (currentSynapse *S_Synapse) DumpCell(indentation []byte) {
	println("synapse_data/DumpCell")
	// 	cellData := currentSynapse.GetData()
	// 	if cellData != nil {
	// 		cellData.DumpCell(currentSynapse, indentation)
	// 	}
}

/*******************
* CreateSynapse
*******************/
func CreateSynapse(brainConfig interfaces.I_BrainConfig, parentSynapse interfaces.I_Synapse, cell interfaces.I_Cell) interfaces.I_Synapse {
	newSynapse := new(S_Synapse)
	newSynapse.brainConfig = brainConfig
	newSynapse.synapseID = brainConfig.GetSynapsesGroupManagement().GetNextSynapseID()

	if parentSynapse != nil {
		if concreteSynapse, ok := parentSynapse.(*S_Synapse); ok {
			concreteSynapse.addChildSynapse(newSynapse)
		}
	}

	brainConfig.GetSynapsesGroupManagement().AppendSynapseToGroup(newSynapse)

	return (newSynapse)
}
