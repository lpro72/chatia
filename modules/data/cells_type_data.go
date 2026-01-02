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
* s_CellType
*******************/
type s_CellType struct {
	id                          uint32
	name                        string
	CreateCellFromSerializeData func([]byte) interfaces.I_CellData
}

/*******************
* s_CellTypeManagement
*******************/
type s_CellTypeManagement struct {
	cellTypeFileHandle     *os.File
	registeredCellTypeList map[string]s_CellType
	nextCellTypeID         uint32
	MemoryAccess           interfaces.I_Lock
}

/*******************
* Globals Variables
*******************/
var g_registeredCellTypeManagement interfaces.I_CellTypeManagement = nil

/*******************
* Internal functions
*******************/
func (cellTypeManagement *s_CellTypeManagement) initialize(brainConfig interfaces.I_BrainConfig) {
	cellTypeManagement.MemoryAccess = &utils.S_Lock{}
	cellTypeManagement.nextCellTypeID = 1
	cellTypeManagement.registeredCellTypeList = make(map[string]s_CellType, 10)
	cellTypeManagement.cellTypeFileHandle = utils.ReadConfigFile(brainConfig, "cell_types.brn", cellTypeManagement.LoadFromFile)
}

func (cellTypeManagement *s_CellTypeManagement) appendToFile(cell s_CellType) bool {
	if cellTypeManagement.cellTypeFileHandle == nil {
		return false
	}
	// Write the brain context name
	_, err := utils.FileWriteString(cellTypeManagement.cellTypeFileHandle, -1, cell.name)
	if err != nil {
		return false
	}
	_, err = utils.FileWriteUint32(cellTypeManagement.cellTypeFileHandle, -1, uint32(cell.id))

	return err == nil
}

/*******************
*  Functions for the interface I_CellTypeManagement
*******************/
func (cellTypeManagement *s_CellTypeManagement) GetCellTypeID(name string) uint32 {
	cellTypeManagement.Lock()
	defer cellTypeManagement.Unlock()

	cellType, exists := cellTypeManagement.registeredCellTypeList[name]
	if !exists {
		return 0
	}

	return (cellType.id)
}

func (cellTypeManagement *s_CellTypeManagement) GetCellNextTypeID() uint32 {
	cellTypeManagement.Lock()
	defer cellTypeManagement.Unlock()

	return cellTypeManagement.nextCellTypeID
}

func (cellTypeManagement *s_CellTypeManagement) GetCellTypeName(id uint32) string {
	cellTypeManagement.Lock()
	defer cellTypeManagement.Unlock()

	for name, cellType := range cellTypeManagement.registeredCellTypeList {
		if cellType.id == id {
			return name
		}
	}

	return ""
}

func (cellTypeManagement *s_CellTypeManagement) AddCellType(name string, CreateCellFromSerializeData func([]byte) interfaces.I_CellData) uint32 {
	cellTypeManagement.Lock()
	defer cellTypeManagement.Unlock()

	cellTypeID := uint32(0)
	cellType, exists := cellTypeManagement.registeredCellTypeList[name]
	if !exists {
		addCell := s_CellType{id: cellTypeManagement.nextCellTypeID, name: name, CreateCellFromSerializeData: CreateCellFromSerializeData}
		if !cellTypeManagement.appendToFile(addCell) {
			return 0
		}
		cellTypeManagement.registeredCellTypeList[name] = addCell
		cellTypeManagement.nextCellTypeID++
		cellTypeID = addCell.id
	} else {
		cellType.CreateCellFromSerializeData = CreateCellFromSerializeData
		cellTypeManagement.registeredCellTypeList[name] = cellType
	}
	return cellTypeID
}

func (cellTypeManagement *s_CellTypeManagement) CreateCellDataFromSerializedData(typeID uint32, serializedData []byte) interfaces.I_CellData {
	cellTypeManagement.Lock()
	defer cellTypeManagement.Unlock()

	for _, cellType := range cellTypeManagement.registeredCellTypeList {
		if cellType.id == typeID {
			return cellType.CreateCellFromSerializeData(serializedData)
		}
	}

	return nil
}

/*******************
*  Functions for the interface I_File
*******************/
func (cellTypeManagement *s_CellTypeManagement) LoadFromFile(fileHandle *os.File, dataOffset int64, brainConfig interfaces.I_BrainConfig, version uint32) {
	cellTypeManagement.Lock()
	defer cellTypeManagement.Unlock()

	for {
		var name string
		var err error
		var cellTypeID uint32

		dataOffset, err = utils.FileReadString(fileHandle, dataOffset, &name)
		if err != nil {
			if err == io.EOF {
				break
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &cellTypeID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		if cellTypeID >= cellTypeManagement.nextCellTypeID {
			cellTypeManagement.nextCellTypeID = cellTypeID + 1
		}

		cellTypeManagement.registeredCellTypeList[name] = s_CellType{id: cellTypeID, name: name}
	}

}

func (cellTypeManagement *s_CellTypeManagement) Close() {
	cellTypeManagement.Lock()
	defer cellTypeManagement.Unlock()

	utils.CloseFile(cellTypeManagement.cellTypeFileHandle)
	cellTypeManagement.cellTypeFileHandle = nil
}

/*******************
* Functions for the interface I_Lock
*******************/
func (cellTypeManagement *s_CellTypeManagement) Lock() {
	cellTypeManagement.MemoryAccess.Lock()
}

func (cellTypeManagement *s_CellTypeManagement) Unlock() {
	cellTypeManagement.MemoryAccess.Unlock()
}

/*******************
* CellType_RegisterNewType
*******************/
func CellType_RegisterNewType(name string, CreateCellFromSerializeData func([]byte) interfaces.I_CellData) uint32 {
	cellTypeID := g_registeredCellTypeManagement.AddCellType(name, CreateCellFromSerializeData)

	return (cellTypeID)
}

/*******************
* CellType_GetTypeID
*******************/
func CellType_GetTypeID(name string) uint32 {
	return (g_registeredCellTypeManagement.GetCellTypeID(name))
}

/*******************
* CellType_Create
*******************/
func CellType_Create(brainConfig interfaces.I_BrainConfig) {
	if g_registeredCellTypeManagement == nil {
		registeredCellTypeManagement := &s_CellTypeManagement{}
		registeredCellTypeManagement.initialize(brainConfig)
		g_registeredCellTypeManagement = registeredCellTypeManagement
	}
}

/*******************
* CellType_Close
*******************/
func CellType_Close() {
	if g_registeredCellTypeManagement != nil {
		g_registeredCellTypeManagement.Close()
		g_registeredCellTypeManagement = nil
	}
}

/*******************
* CellType_CreateCellDataFromSerializedData
*******************/
func CellType_CreateCellDataFromSerializedData(typeID uint32, serializedData []byte) interfaces.I_CellData {
	return (g_registeredCellTypeManagement.CreateCellDataFromSerializedData(typeID, serializedData))
}
