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
	synapsesGroup.fileHandle = utils.ReadConfigFile(synapsesGroup.brainConfig, fmt.Sprintf("synapses_group_%d.brn", synapsesGroup.SynapsegroupID), synapsesGroup.LoadFromFile)
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
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
			os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.cellID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
			os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.score)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
			os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.nextSynapseID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
			os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.previousSynapseID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
			os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
		}
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, concreteSynapse.parentSynapseID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
			os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
		}
		totalChildren := uint32(len(concreteSynapse.childSynapseIDList))
		dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, totalChildren)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
			os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
		}
		tempList := make([]uint32, 10)
		copy(tempList, concreteSynapse.childSynapseIDList)
		for _, childID := range tempList {
			dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, childID)
			if err != nil {
				errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
				os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
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

		synapse := CreateSynapse(brainConfig, nil, nil)
		concreteSynapse, ok := synapse.(*S_Synapse)
		if !ok {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CELL_INVALID_DATA)
			os.Exit(errcode.ERROR_FATAL_CELL_INVALID_DATA)
		}

		// Read synapse ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			if err == io.EOF {
				break
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}
		concreteSynapse.synapseID = value

		// Read cell ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}
		concreteSynapse.cellID = value

		// Read score
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}
		concreteSynapse.score = value

		// Read next synapse ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}
		concreteSynapse.nextSynapseID = value

		// Read previous synapse ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}
		concreteSynapse.previousSynapseID = value

		// Read parent synapse ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}
		concreteSynapse.parentSynapseID = value

		// Read child list size
		var childListSize uint32 = 0
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &childListSize)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		for i := uint32(0); i < childListSize; i++ {
			dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &value)
			if err != nil {
				errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
				os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
			}
			concreteSynapse.childSynapseIDList = append(concreteSynapse.childSynapseIDList, value)
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
