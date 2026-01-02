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
* Types
*******************/
type S_SynapsesGroupManagement struct {
	fileHandle *os.File

	synapsesGroupList  []interfaces.I_SynapsesGroup
	synapsesGroupCount uint32

	MemoryAccess interfaces.I_Lock

	brainConfig interfaces.I_BrainConfig
}

/*******************
* Internal functions
*******************/
func (synapsesGroupManagement *S_SynapsesGroupManagement) Initialize(brainConfig interfaces.I_BrainConfig) {
	synapsesGroupManagement.brainConfig = brainConfig
	synapsesGroupManagement.MemoryAccess = &utils.S_Lock{}
	synapsesGroupManagement.synapsesGroupList = make([]interfaces.I_SynapsesGroup, 0)
	synapsesGroupManagement.synapsesGroupCount = 0
	synapsesGroupManagement.fileHandle = utils.ReadConfigFile(brainConfig, "synapses_group_management.brn", synapsesGroupManagement.LoadFromFile)
}

func (synapsesGroupManagement *S_SynapsesGroupManagement) addGroupToFile() {

	if synapsesGroupManagement.fileHandle == nil {
		return
	}

	_, err := utils.FileWriteUint32(synapsesGroupManagement.fileHandle, 4, uint32(synapsesGroupManagement.synapsesGroupCount))
	if err != nil {
		fmt.Println("Error writing group count:", err)
	}
}

/*******************
*  Functions for the interface I_SynapsesGroupManagement
*******************/
func (synapsesGroupManagement *S_SynapsesGroupManagement) AppendSynapseToGroup(synapse interfaces.I_Synapse) uint32 {
	synapsesGroupManagement.Lock()
	defer synapsesGroupManagement.Unlock()

	lastSynapseGroupID := int32(synapsesGroupManagement.synapsesGroupCount) - 1
	if lastSynapseGroupID == -1 || synapsesGroupManagement.synapsesGroupList[lastSynapseGroupID].GetSynapsesCount() >= 1024 {
		newSynapseGroup := SynapsesGroup_Create(synapsesGroupManagement.brainConfig, synapsesGroupManagement.synapsesGroupCount)
		synapsesGroupManagement.synapsesGroupList = append(synapsesGroupManagement.synapsesGroupList, newSynapseGroup)
		synapsesGroupManagement.synapsesGroupCount++
		lastSynapseGroupID++
		synapsesGroupManagement.addGroupToFile()
	}
	synapsesGroup := synapsesGroupManagement.synapsesGroupList[lastSynapseGroupID]
	synapsesGroup.AppendSynapseToGroup(synapse)

	return uint32(lastSynapseGroupID*1024) + synapsesGroup.GetSynapsesCount()
}

func (synapsesGroupManagement *S_SynapsesGroupManagement) GetSynapseFromID(synapseID uint32) interfaces.I_Synapse {
	synapsesGroupManagement.Lock()
	defer synapsesGroupManagement.Unlock()

	synapsesGroupListCount := len(synapsesGroupManagement.synapsesGroupList)
	if synapseID == 0 || synapsesGroupListCount == 0 {
		return nil
	}

	synapseID -= 1
	groupeID := synapseID / 1024
	synapseIDInGroup := synapseID % 1024

	if groupeID > uint32(synapsesGroupListCount-1) {
		return nil
	}
	synapseGroup := synapsesGroupManagement.synapsesGroupList[groupeID]
	if synapseIDInGroup >= synapseGroup.GetSynapsesCount() {
		return nil
	}
	return synapseGroup.GetSynapseFromID(synapseIDInGroup)
}

func (synapsesGroupManagement *S_SynapsesGroupManagement) GetSynapsesGroupsCount() int {
	synapsesGroupManagement.Lock()
	defer synapsesGroupManagement.Unlock()

	println("synapses_group_management_data/GetSynapsesGroupsCount")
	// 	return len(synapsesGroupManagement.synapsesGroupList)
	return 0
}

func (synapsesGroupManagement *S_SynapsesGroupManagement) GetSynapsesCount(groupID int) int {
	synapsesGroupManagement.Lock()
	defer synapsesGroupManagement.Unlock()

	println("synapses_group_management_data/GetSynapsesCount")
	// 	if groupID >= 0 && groupID < len(synapsesGroupManagement.synapseGroupList) {
	// 		return synapsesGroupManagement.synapseGroupList[groupID].SynapseCount
	// 	}

	// 	// If groupID is -1, return the total synapse count across all groups
	// 	if groupID == -1 {
	// 		total := 0
	// 		for _, group := range synapsesGroupManagement.synapseGroupList {
	// 			total += group.SynapseCount
	// 		}
	// 		return total
	// 	}

	// Invalid groupID
	return 0
}

func (synapsesGroupManagement *S_SynapsesGroupManagement) GetNextSynapseID() uint32 {
	synapsesGroupManagement.Lock()
	defer synapsesGroupManagement.Unlock()

	if synapsesGroupManagement.synapsesGroupCount == 0 {
		return 1
	}

	lastSynapseGroupID := int32(synapsesGroupManagement.synapsesGroupCount) - 1
	synapsesGroup := synapsesGroupManagement.synapsesGroupList[lastSynapseGroupID]
	nextSynapseID := uint32(lastSynapseGroupID*1024) + synapsesGroup.GetSynapsesCount() + 1
	return nextSynapseID
}

/*******************
*  Functions for the interface I_File
*******************/
func (synapsesGroupManagement *S_SynapsesGroupManagement) LoadFromFile(fileHandle *os.File, dataOffset int64, brainConfig interfaces.I_BrainConfig, version uint32) {
	for {
		// Read the synapse group name
		var synapsesGroupSize uint32
		var err error
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &synapsesGroupSize)
		if err != nil {
			if err == io.EOF {
				return
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		for i := uint32(0); i < synapsesGroupSize; i++ {
			newSynapsesGroup := SynapsesGroup_Create(brainConfig, i)
			synapsesGroupManagement.synapsesGroupList = append(synapsesGroupManagement.synapsesGroupList, newSynapsesGroup)
			synapsesGroupManagement.synapsesGroupCount++
		}
	}

}

func (synapsesGroupManagement *S_SynapsesGroupManagement) Close() {
	synapsesGroupManagement.Lock()
	defer synapsesGroupManagement.Unlock()

	utils.CloseFile(synapsesGroupManagement.fileHandle)
	synapsesGroupManagement.fileHandle = nil
}

/*******************
* Functions for the interface I_Lock
*******************/
func (synapsesGroupManagement *S_SynapsesGroupManagement) Lock() {
	synapsesGroupManagement.MemoryAccess.Lock()
}

func (synapsesGroupManagement *S_SynapsesGroupManagement) Unlock() {
	synapsesGroupManagement.MemoryAccess.Unlock()
}

/*******************
* SynapsesGroupManagement_Create
*******************/
func SynapsesGroupManagement_Create(brainConfig interfaces.I_BrainConfig) interfaces.I_SynapsesGroupManagement {
	synapsesGroupManagement := new(S_SynapsesGroupManagement)
	synapsesGroupManagement.Initialize(brainConfig)
	return synapsesGroupManagement
}
