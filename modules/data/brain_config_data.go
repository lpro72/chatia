package data

/*******************
* Import
*******************/
import (
	"os"
	"path/filepath"

	"chatia/modules/errcode"
	"chatia/modules/interfaces"
	"chatia/modules/utils"
)

/*******************
* Globals Varables
*******************/
var g_MainBrain interfaces.I_BrainConfig = nil
var g_Brain interfaces.I_BrainConfig = nil

/*******************
* S_BrainConfig
*******************/
type S_BrainConfig struct {
	mainDirectory  string
	savesDirectory string

	brainContextManagement  interfaces.I_BrainContextManagement
	cellsGroupManagement    interfaces.I_CellsGroupManagement
	synapsesGroupManagement interfaces.I_SynapsesGroupManagement
	MemoryAccess            interfaces.I_Lock
}

/*******************
* Internal functions
*******************/
func (brainConfig *S_BrainConfig) CallFactoryContext() {
	for name, factory := range g_RegisterBrainContext {
		brainContextManagement := brainConfig.brainContextManagement
		brainContext := brainContextManagement.GetBrainContext(name)
		if brainContext == nil {
			brainContext = brainContextManagement.CreateNewBrainContext(brainConfig, name, 0)
		}
		factory(brainContext)
	}
}

/*******************
*  Functions for the interface I_BrainConfig
*******************/
func (brainConfig *S_BrainConfig) SetMainDirectory(directory string) {
	brainConfig.Lock()
	defer brainConfig.Unlock()

	brainConfig.mainDirectory = directory
}

func (brainConfig *S_BrainConfig) GetMainDirectory() string {
	brainConfig.Lock()
	defer brainConfig.Unlock()

	return brainConfig.mainDirectory
}

func (brainConfig *S_BrainConfig) SetSaveDirectory(directory string) {
	brainConfig.Lock()
	defer brainConfig.Unlock()

	brainConfig.savesDirectory = directory
}

func (brainConfig *S_BrainConfig) GetSaveDirectory() string {
	brainConfig.Lock()
	defer brainConfig.Unlock()

	return brainConfig.savesDirectory
}

func (brainConfig *S_BrainConfig) GetBrainContextManagement() interfaces.I_BrainContextManagement {
	brainConfig.Lock()
	defer brainConfig.Unlock()

	return brainConfig.brainContextManagement
}

func (brainConfig *S_BrainConfig) GetCellsGroupManagament() interfaces.I_CellsGroupManagement {
	brainConfig.Lock()
	defer brainConfig.Unlock()

	return brainConfig.cellsGroupManagement

}

func (brainConfig *S_BrainConfig) GetSynapsesGroupManagement() interfaces.I_SynapsesGroupManagement {
	brainConfig.Lock()
	defer brainConfig.Unlock()

	return brainConfig.synapsesGroupManagement
}

/*******************
*  Functions for the interface I_File
*******************/
func (brainConfig *S_BrainConfig) LoadFromFile(fileHandle *os.File, dataOffset int64, brainConfigInterface interfaces.I_BrainConfig, version uint32) {
	// Nothing to do here
	// Only defined to respect the I_File interface
}

func (brainConfig *S_BrainConfig) Close() {
	brainConfig.Lock()
	defer brainConfig.Unlock()

	brainConfig.brainContextManagement.Close()
	brainConfig.cellsGroupManagement.Close()
}

/*******************
* Functions for the interface I_Lock
*******************/
func (brainConfig *S_BrainConfig) Lock() {
	brainConfig.MemoryAccess.Lock()
}

func (brainConfig *S_BrainConfig) Unlock() {
	brainConfig.MemoryAccess.Unlock()
}

/*******************
* BrainConfig_Create
*******************/
func BrainConfig_Create() interfaces.I_BrainConfig {
	brainConfig := new(S_BrainConfig)
	brainConfig.MemoryAccess = &utils.S_Lock{}

	exec, err := os.Executable()
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_PROG_NOT_FOUND)
		panic(err)
	}
	brainConfig.SetMainDirectory(filepath.Dir(exec))
	brainConfig.SetSaveDirectory(filepath.Join(brainConfig.GetMainDirectory(), "save"))

	CellType_Create(brainConfig)
	brainConfig.brainContextManagement = BrainContextManagement_Create(brainConfig)
	brainConfig.cellsGroupManagement = CellsGroupManagement_Create(brainConfig)
	brainConfig.synapsesGroupManagement = SynapsesGroupManagement_Create(brainConfig)
	brainConfig.CallFactoryContext()

	return brainConfig
}

/*******************
* BrainConfig_Init
*******************/
func BrainConfig_Init() {
	g_MainBrain = BrainConfig_Create()
	g_Brain = g_MainBrain
}

/*******************
* BrainConfig_Close
*******************/
func BrainConfig_Close() {
	if g_MainBrain != nil {
		g_MainBrain.Close()
	}
}

/*******************
* BrainConfig_Register
*******************/
func BrainConfig_Register() {
	g_MainBrain.CallFactoryContext()
}

/*******************
* UseMainBrain
*******************/
func UseMainBrain() interfaces.I_BrainConfig {
	g_Brain.Lock()
	defer g_Brain.Unlock()

	g_Brain = g_MainBrain

	return g_Brain
}

/*******************
* UseTemporaryBrain
*******************/
func UseTemporaryBrain() interfaces.I_BrainConfig {
	g_Brain.Lock()
	defer g_Brain.Unlock()

	g_Brain = BrainConfig_Create()

	return g_Brain
}
