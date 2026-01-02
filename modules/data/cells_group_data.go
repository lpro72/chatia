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
* S_CellsGroup
*******************/
type S_CellsGroup struct {
	brainConfig interfaces.I_BrainConfig

	CellgroupID uint32
	CellCount   uint32
	CellList    []interfaces.I_Cell

	fileHandle *os.File
	loaded     bool
	dataOffset int64

	MemoryAccess interfaces.I_Lock
}

/*******************
* Internal functions
*******************/
func (cellsGroup *S_CellsGroup) Initialize(brainConfig interfaces.I_BrainConfig, cellsGroupeID uint32) {
	cellsGroup.MemoryAccess = &utils.S_Lock{}
	cellsGroup.CellList = make([]interfaces.I_Cell, 0)
	cellsGroup.CellCount = 0
	cellsGroup.CellgroupID = cellsGroupeID
	cellsGroup.brainConfig = brainConfig
}

func (cellsGroup *S_CellsGroup) LoadCellsGroupFromFile() {
	if cellsGroup.loaded {
		return
	}
	cellsGroup.loaded = true
	cellsGroup.fileHandle = utils.ReadConfigFile(cellsGroup.brainConfig, fmt.Sprintf("cells_group_%d.brn", cellsGroup.CellgroupID), cellsGroup.LoadFromFile)
}

func (cellsGroup *S_CellsGroup) appendCellToFile(cell interfaces.I_Cell) {
	if cellsGroup.fileHandle == nil {
		return
	}

	dataOffset, err := utils.FileGetEndOffset(cellsGroup.fileHandle)
	if err != nil {
		return
	}
	cellsGroup.dataOffset = dataOffset

	dataOffset, err = utils.FileWriteUint32(cellsGroup.fileHandle, dataOffset, cell.GetID())
	if err != nil {
		return
	}

	dataOffset, err = utils.FileWriteUint32(cellsGroup.fileHandle, dataOffset, cell.GetCellType())
	if err != nil {
		return
	}

	dataOffset, err = utils.FileWriteData(cellsGroup.fileHandle, dataOffset, cell.GetSerializedData())
	if err != nil {
		return
	}
}

/*******************
*  Functions for the interface I_CellsGroup
*******************/
func (cellsGroup *S_CellsGroup) GetCellCount() uint32 {
	cellsGroup.LoadCellsGroupFromFile()
	return cellsGroup.CellCount
}

func (cellsGroup *S_CellsGroup) AppendCellToGroup(cell interfaces.I_Cell) {
	cellsGroup.LoadCellsGroupFromFile()
	cellsGroup.CellList = append(cellsGroup.CellList, cell)
	cellsGroup.CellCount += 1

	cellsGroup.appendCellToFile(cell)

}

func (cellsGroup *S_CellsGroup) GetCellFromID(cellID uint32) interfaces.I_Cell {
	cellsGroup.LoadCellsGroupFromFile()

	cell := cellsGroup.CellList[cellID]
	return cell
}

/*******************
*  Functions for the interface I_File
*******************/
func (cellsGroup *S_CellsGroup) LoadFromFile(fileHandle *os.File, dataOffset int64, brainConfig interfaces.I_BrainConfig, version uint32) {
	for {
		// Read the cell group name
		var cellID uint32
		var cellType uint32
		var err error

		// Read cell ID
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &cellID)
		if err != nil {
			if err == io.EOF {
				return
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		// Read cell type
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &cellType)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		// Read cell data
		var dataBuffer []byte
		dataOffset, err = utils.FileReadData(fileHandle, dataOffset, &dataBuffer)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}
		cellData := CellType_CreateCellDataFromSerializedData(cellType, dataBuffer)
		cell := CreateCell(brainConfig, cellData, cellType)
		cellsGroup.AppendCellToGroup(cell)
	}
}

func (cellsGroup *S_CellsGroup) Close() {
	cellsGroup.Lock()
	defer cellsGroup.Unlock()

	utils.CloseFile(cellsGroup.fileHandle)
	cellsGroup.fileHandle = nil
}

/*******************
* Functions for the interface I_Lock
*******************/
func (cellsGroup *S_CellsGroup) Lock() {
	cellsGroup.MemoryAccess.Lock()
}

func (cellsGroup *S_CellsGroup) Unlock() {
	cellsGroup.MemoryAccess.Unlock()
}

/*******************
* CellsGroupManagement_Create
*******************/
func CellsGroup_Create(brainConfig interfaces.I_BrainConfig, cellsGroupeID uint32) interfaces.I_CellsGroup {
	cellsGroup := new(S_CellsGroup)
	cellsGroup.Initialize(brainConfig, cellsGroupeID)
	return cellsGroup
}
