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

	SynapsegroupID uint32
	SynapseCount   uint32
	SynapseList    []interfaces.I_Synapse

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
	synapsesGroup.SynapseCount = 0
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

	dataOffset, err = utils.FileWriteUint32(synapsesGroup.fileHandle, dataOffset, synapse.GetID())
	if err != nil {
		return
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
		var synapseID uint32
		var err error

		// Read synapse ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &synapseID)
		if err != nil {
			if err == io.EOF {
				return
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		synapse := CreateSynapse(brainConfig, nil, nil)
		synapsesGroup.AppendSynapseToGroup(synapse)
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
