package data

/*******************
* Import
*******************/
import (
	"io"
	"os"

	"chatia/modules/errcode"
	"chatia/modules/interfaces"
	"chatia/modules/utils"
)

/*******************
* Types
*******************/
type S_CellsGroupManagement struct {
	fileHandle *os.File
	fileName   string

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
	cellsGroupManagement.fileName = "cells_group_management.brn"
	cellsGroupManagement.fileHandle = utils.ReadConfigFile(brainConfig, cellsGroupManagement.fileName, cellsGroupManagement.LoadFromFile)
}

func (cellsGroupManagement *S_CellsGroupManagement) addGroupeToFile() {

	if cellsGroupManagement.fileHandle == nil {
		return
	}

	cellsGroupManagement.Lock()
	defer cellsGroupManagement.Unlock()

	_, err := utils.FileWriteUint32(cellsGroupManagement.fileHandle, 4, uint32(cellsGroupManagement.cellsGroupCount))
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_FILE_WRITE, cellsGroupManagement.fileName, err.Error())
		return
	}
}

/*******************
*  Functions for the interface I_CellsGroupManagement
*******************/
func (cellsGroupManagement *S_CellsGroupManagement) AppendCellToGroup(cell interfaces.I_Cell) uint32 {
	cellsGroupManagement.Lock()
	defer cellsGroupManagement.Unlock()

	lastCellGroupID := int32(cellsGroupManagement.cellsGroupCount) - 1
	if lastCellGroupID == -1 || cellsGroupManagement.cellsGroupList[lastCellGroupID].GetCellsCount() >= 1024 {
		newCellGroup := CellsGroup_Create(cellsGroupManagement.brainConfig, cellsGroupManagement.cellsGroupCount)
		cellsGroupManagement.cellsGroupList = append(cellsGroupManagement.cellsGroupList, newCellGroup)
		cellsGroupManagement.cellsGroupCount++
		lastCellGroupID++
		cellsGroupManagement.addGroupeToFile()
	}
	cellsGroup := cellsGroupManagement.cellsGroupList[lastCellGroupID]
	cellsGroup.AppendCellToGroup(cell)

	return uint32(lastCellGroupID*1024) + cellsGroup.GetCellsCount()
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
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_NOT_FOUND, cellID)
		return nil
	}
	cellGroup := cellsGroupManamgement.cellsGroupList[groupeID]
	if cellIDInGroup >= cellGroup.GetCellsCount() {
		errcode.PrintMsgFromErrorCode(errcode.WARNING_CELL_NOT_FOUND, cellID)
		return nil
	}
	return cellGroup.GetCellFromID(cellIDInGroup)
}

func (cellsGroupManamgement *S_CellsGroupManagement) GetCellGroupsCount() uint32 {
	cellsGroupManamgement.Lock()
	defer cellsGroupManamgement.Unlock()

	return uint32(cellsGroupManamgement.cellsGroupCount)
}

func (cellsGroupManamgement *S_CellsGroupManagement) GetCellsCount() uint32 {
	cellsGroupManamgement.Lock()
	defer cellsGroupManamgement.Unlock()

	var total uint32 = 0
	for _, cellGroup := range cellsGroupManamgement.cellsGroupList {
		total += cellGroup.GetCellsCount()
	}

	return total
}

func (cellsGroupManagement *S_CellsGroupManagement) GetNextCellID() uint32 {
	cellsGroupManagement.Lock()
	defer cellsGroupManagement.Unlock()

	if cellsGroupManagement.cellsGroupCount == 0 {
		return 1
	}

	lastCellGroupID := int32(cellsGroupManagement.cellsGroupCount) - 1
	cellsGroup := cellsGroupManagement.cellsGroupList[lastCellGroupID]
	nextCellID := uint32(lastCellGroupID*1024) + cellsGroup.GetCellsCount() + 1
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
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ, cellsGroupManagement.fileName, err.Error())
			return
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
