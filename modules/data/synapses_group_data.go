package data

/*******************
* Import
*******************/
import (
	"fmt"
	"io"
	"os"

	"chatia/modules/errcode"
	"chatia/modules/interfaces"
	"chatia/modules/utils"
)

/*******************
* S_SynapsesGroup
*******************/
type S_SynapsesGroup struct {
	brainConfig interfaces.I_BrainConfig

	SynapsegroupID    uint32
	SynapseCount      uint32
	SynapseList       []interfaces.I_Synapse
	SynapseExtendedID uint32

	fileHandle *os.File
	fileName   string
	loaded     bool
	dataOffset int64

	MemoryAccess interfaces.I_Lock
}

/*******************
* Internal functions
*******************/
func (synapsesGroup *S_SynapsesGroup) Initialize(brainConfig interfaces.I_BrainConfig, synapsesGroupeID uint32) {
	synapsesGroup.MemoryAccess = &utils.S_Lock{}
	synapsesGroup.SynapseList = make([]interfaces.I_Synapse, 0)
	synapsesGroup.SynapsegroupID = synapsesGroupeID
	synapsesGroup.brainConfig = brainConfig
}

func (synapsesGroup *S_SynapsesGroup) LoadSynapsesGroupFromFile() {
	if synapsesGroup.loaded {
		return
	}
	synapsesGroup.loaded = true
	synapsesGroup.fileName = fmt.Sprintf("synapses_group_%d.brn", synapsesGroup.SynapsegroupID)
	synapsesGroup.fileHandle = utils.ReadConfigFile(synapsesGroup.brainConfig, synapsesGroup.fileName, synapsesGroup.LoadFromFile)
}

func (synapsesGroup *S_SynapsesGroup) appendSynapseToFile(synapse interfaces.I_Synapse) {
	if synapsesGroup.fileHandle == nil {
		return
	}

	dataOffset, err := utils.FileGetEndOffset(synapsesGroup.fileHandle)
	if err != nil {
		return
	}
	synapsesGroup.dataOffset = dataOffset

	if concreteSynapse, ok := synapse.(*S_Synapse); ok {
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.synapseID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE, synapsesGroup.fileName, err.Error())
			return
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.cellID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE, synapsesGroup.fileName, err.Error())
			return
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.score)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE, synapsesGroup.fileName, err.Error())
			return
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.nextSynapseID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE, synapsesGroup.fileName, err.Error())
			return
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.previousSynapseID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE, synapsesGroup.fileName, err.Error())
			return
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.parentSynapseID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE, synapsesGroup.fileName, err.Error())
			return
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.maxChildListSize)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE, synapsesGroup.fileName, err.Error())
			return
		}
		tempList := make([]uint32, concreteSynapse.maxChildListSize)
		copy(tempList, concreteSynapse.childSynapseIDList)
		for _, childID := range tempList {
			dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, childID)
			if err != nil {
				errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE, synapsesGroup.fileName, err.Error())
				return
			}
		}
	}
}

/*******************
*  Functions for the interface I_SynapsesGroup
*******************/
func (synapsesGroup *S_SynapsesGroup) GetSynapsesCount() uint32 {
	synapsesGroup.LoadSynapsesGroupFromFile()
	return synapsesGroup.SynapseCount
}

func (synapsesGroup *S_SynapsesGroup) AppendSynapseToGroup(synapse interfaces.I_Synapse) {
	synapsesGroup.LoadSynapsesGroupFromFile()
	synapsesGroup.SynapseList = append(synapsesGroup.SynapseList, synapse)
	synapsesGroup.SynapseCount += 1

	synapsesGroup.appendSynapseToFile(synapse)

}

func (synapsesGroup *S_SynapsesGroup) GetSynapseFromID(synapseID uint32) interfaces.I_Synapse {
	synapsesGroup.LoadSynapsesGroupFromFile()

	synapse := synapsesGroup.SynapseList[synapseID]
	return synapse
}

/*******************
*  Functions for the interface I_File
*******************/
func (synapsesGroup *S_SynapsesGroup) LoadFromFile(fileHandle *os.File, dataOffset int64, brainConfig interfaces.I_BrainConfig, version uint32) {
	for {
		// Read the synapse group name
		var value uint32
		var err error

		synapse := CreateSynapse(brainConfig, nil, nil, 0)
		concreteSynapse, ok := synapse.(*S_Synapse)
		if !ok {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CELL_CREATE)
			return
		}

		// Read synapse ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			if err == io.EOF {
				break
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ, synapsesGroup.fileName, err.Error())
			return
		}
		concreteSynapse.synapseID = value

		// Read cell ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ, synapsesGroup.fileName, err.Error())
			return
		}
		concreteSynapse.cellID = value

		// Read score
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ, synapsesGroup.fileName, err.Error())
			return
		}
		concreteSynapse.score = value

		// Read next synapse ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ, synapsesGroup.fileName, err.Error())
			return
		}
		concreteSynapse.nextSynapseID = value

		// Read previous synapse ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ, synapsesGroup.fileName, err.Error())
			return
		}
		concreteSynapse.previousSynapseID = value

		// Read parent synapse ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ, synapsesGroup.fileName, err.Error())
			return
		}
		concreteSynapse.parentSynapseID = value

		// Read max child list size
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ, synapsesGroup.fileName, err.Error())
			return
		}
		concreteSynapse.maxChildListSize = value

		concreteSynapse.childSynapseIDList = make([]uint32, 0, concreteSynapse.maxChildListSize)
		for i := uint32(0); i < concreteSynapse.maxChildListSize; i++ {
			dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
			if err != nil {
				errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ, synapsesGroup.fileName, err.Error())
				return
			}
			if value != 0 {
				concreteSynapse.childSynapseIDList = append(concreteSynapse.childSynapseIDList, value)
			}
		}
	}
}

func (synapsesGroup *S_SynapsesGroup) Close() {
	synapsesGroup.Lock()
	defer synapsesGroup.Unlock()

	utils.CloseFile(synapsesGroup.fileHandle)
	synapsesGroup.fileHandle = nil
}

/*******************
* Functions for the interface I_Lock
*******************/
func (synapsesGroup *S_SynapsesGroup) Lock() {
	synapsesGroup.MemoryAccess.Lock()
}

func (synapsesGroup *S_SynapsesGroup) Unlock() {
	synapsesGroup.MemoryAccess.Unlock()
}

/*******************
* SynapsesGroupManagement_Create
*******************/
func SynapsesGroup_Create(brainConfig interfaces.I_BrainConfig, synapsesGroupeID uint32) interfaces.I_SynapsesGroup {
	synapsesGroup := new(S_SynapsesGroup)
	synapsesGroup.Initialize(brainConfig, synapsesGroupeID)
	return synapsesGroup
}
