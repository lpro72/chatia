package data

/*******************
* Import
*******************/
import (

	// "io"
	"fmt"
	"io"
	"os"

	// "chatia/modules/errcode"
	"chatia/modules/errcode"
	"chatia/modules/interfaces"
	"chatia/modules/utils"
)

/*******************
* Types
*******************/
type S_CellsGroupManagement struct {
	fileHandle *os.File

	cellsGroupList  []interfaces.I_CellsGroup
	cellsGroupCount uint32

	MemoryAccess interfaces.I_Lock

	brainConfig interfaces.I_BrainConfig
}

/*******************
* Internal functions
*******************/
func (cellsGroupManagement *S_CellsGroupManagement) Initialize(brainConfig interfaces.I_BrainConfig) {
	cellsGroupManagement.brainConfig = brainConfig
	cellsGroupManagement.MemoryAccess = &utils.S_Lock{}
	cellsGroupManagement.cellsGroupList = make([]interfaces.I_CellsGroup, 0)
	cellsGroupManagement.cellsGroupCount = 0
	cellsGroupManagement.fileHandle = utils.ReadConfigFile(brainConfig, "cells_group_management.brn", cellsGroupManagement.LoadFromFile)
}

func (cellsGroupManagement *S_CellsGroupManagement) addGroupeToFile() {

	if cellsGroupManagement.fileHandle == nil {
		return
	}

	_, err := utils.FileWriteUint32(cellsGroupManagement.fileHandle, 4, uint32(cellsGroupManagement.cellsGroupCount))
	if err != nil {
		fmt.Println("Error writing group count:", err)
	}
}

/*******************
*  Functions for the interface I_CellsGroupManagement
*******************/
func (cellsGroupManagement *S_CellsGroupManagement) AppendCellToGroup(cell interfaces.I_Cell) uint32 {
	cellsGroupManagement.Lock()
	defer cellsGroupManagement.Unlock()

	lastCellGroupID := int32(cellsGroupManagement.cellsGroupCount) - 1
	if lastCellGroupID == -1 || cellsGroupManagement.cellsGroupList[lastCellGroupID].GetCellCount() >= 1024 {
		newCellGroup := CellsGroup_Create(cellsGroupManagement.brainConfig, cellsGroupManagement.cellsGroupCount)
		cellsGroupManagement.cellsGroupList = append(cellsGroupManagement.cellsGroupList, newCellGroup)
		cellsGroupManagement.cellsGroupCount++
		lastCellGroupID++
		cellsGroupManagement.addGroupeToFile()
	}
	cellsGroup := cellsGroupManagement.cellsGroupList[lastCellGroupID]
	cellsGroup.AppendCellToGroup(cell)

	return uint32(lastCellGroupID*1024) + cellsGroup.GetCellCount()
}

func (cellsGroupManamgement *S_CellsGroupManagement) GetCellFromID(cellID uint32) interfaces.I_Cell {
	cellsGroupManamgement.Lock()
	defer cellsGroupManamgement.Unlock()

	cellsGroupListCount := len(cellsGroupManamgement.cellsGroupList)
	if cellID == 0 || cellsGroupListCount == 0 {
		return nil
	}

	cellID -= 1
	groupeID := cellID / 1024
	cellIDInGroup := cellID % 1024

	if groupeID > uint32(cellsGroupListCount-1) {
		return nil
	}
	cellGroup := cellsGroupManamgement.cellsGroupList[groupeID]
	if cellIDInGroup >= cellGroup.GetCellCount() {
		return nil
	}
	return cellGroup.GetCellFromID(cellIDInGroup)
}

func (cellsGroupManamgement *S_CellsGroupManagement) GetCellGroupsCount() int {
	cellsGroupManamgement.Lock()
	defer cellsGroupManamgement.Unlock()

	println("cells_group_management_data/GetCellGroupsCount")
	// 	return len(cellsGroupManamgement.cellGroupList)
	return 0
}

func (cellsGroupManamgement *S_CellsGroupManagement) GetCellCount(groupID int) int {
	cellsGroupManamgement.Lock()
	defer cellsGroupManamgement.Unlock()

	println("cells_group_management_data/GetCellCount")
	// 	if groupID >= 0 && groupID < len(cellsGroupManamgement.cellGroupList) {
	// 		return cellsGroupManamgement.cellGroupList[groupID].CellCount
	// 	}

	// 	// If groupID is -1, return the total cell count across all groups
	// 	if groupID == -1 {
	// 		total := 0
	// 		for _, group := range cellsGroupManamgement.cellGroupList {
	// 			total += group.CellCount
	// 		}
	// 		return total
	// 	}

	// Invalid groupID
	return 0
}

func (cellsGroupManagement *S_CellsGroupManagement) GetNextCellID() uint32 {
	cellsGroupManagement.Lock()
	defer cellsGroupManagement.Unlock()

	if cellsGroupManagement.cellsGroupCount == 0 {
		return 1
	}

	lastCellGroupID := int32(cellsGroupManagement.cellsGroupCount) - 1
	cellsGroup := cellsGroupManagement.cellsGroupList[lastCellGroupID]
	nextCellID := uint32(lastCellGroupID*1024) + cellsGroup.GetCellCount() + 1
	return nextCellID
}

/*******************
*  Functions for the interface I_File
*******************/
func (cellsGroupManagement *S_CellsGroupManagement) LoadFromFile(fileHandle *os.File, dataOffset int64, brainConfig interfaces.I_BrainConfig, version uint32) {
	for {
		// Read the cell group name
		var cellsGroupSize uint32
		var err error
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &cellsGroupSize)
		if err != nil {
			if err == io.EOF {
				return
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		for i := uint32(0); i < cellsGroupSize; i++ {
			newCellGroup := CellsGroup_Create(brainConfig, i)
			cellsGroupManagement.cellsGroupList = append(cellsGroupManagement.cellsGroupList, newCellGroup)
			cellsGroupManagement.cellsGroupCount++
		}
	}

}

func (cellsGroupManagement *S_CellsGroupManagement) Close() {
	cellsGroupManagement.Lock()
	defer cellsGroupManagement.Unlock()

	utils.CloseFile(cellsGroupManagement.fileHandle)
	cellsGroupManagement.fileHandle = nil
}

/*******************
* Functions for the interface I_Lock
*******************/
func (cellsGroupManagement *S_CellsGroupManagement) Lock() {
	cellsGroupManagement.MemoryAccess.Lock()
}

func (cellsGroupManagement *S_CellsGroupManagement) Unlock() {
	cellsGroupManagement.MemoryAccess.Unlock()
}

/*******************
* CellsGroupManagement_Create
*******************/
func CellsGroupManagement_Create(brainConfig interfaces.I_BrainConfig) interfaces.I_CellsGroupManagement {
	cellsGroupManagement := new(S_CellsGroupManagement)
	cellsGroupManagement.Initialize(brainConfig)
	return cellsGroupManagement
}
