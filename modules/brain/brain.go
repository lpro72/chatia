package brain

/*******************
* Import
*******************/
import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"chatia/modules/brain/cell"
	"chatia/modules/errcode"
	"chatia/modules/utils"
)

/*******************
* Interfaces
*******************/
type I_Brain interface {
	GetFirstCell() cell.I_CellManagement
	SetFirstCell(firstCell cell.I_CellManagement)
	SetLearnFunction(learn func(data []byte, firstCell cell.I_CellManagement))
	SetExecFunction(exec func(command string, extraVar ...any) string)
	SetUnittestFunction(unittest func())
	SetDumpMemoryFunction(dumpMemory func())
	CallLearnFunction(data []byte, firstCell cell.I_CellManagement)
	CallExecFunction(command string, extraVar ...any) string
	CallUnittestFunction()
	CallDumpMemoryFunction()
}

/*******************
* Types
*******************/
type s_Brain struct {
	learn    func(data []byte, firstCell cell.I_CellManagement)
	exec     func(command string, extraVar ...any) string
	unittest func()

	// Data
	firstCell cell.I_CellManagement
	mutex     sync.RWMutex

	// Debug only
	dumpMemory func()
}

type s_BrainConfig struct {
	mainDirectory  string
	savesDirectory string
	contextList    map[string]*s_Brain
	mutex          sync.RWMutex
	fileHandle     *os.File
}

/*******************
* s_Brain
*******************/
func (brain *s_Brain) GetFirstCell() cell.I_CellManagement {
	return brain.firstCell
}

func (brain *s_Brain) SetFirstCell(firstCell cell.I_CellManagement) {
	brain.firstCell = firstCell
}

func (brain *s_Brain) SetLearnFunction(learn func(data []byte, firstCell cell.I_CellManagement)) {
	brain.learn = learn
}

func (brain *s_Brain) SetExecFunction(exec func(command string, extraVar ...any) string) {
	brain.exec = exec
}

func (brain *s_Brain) SetUnittestFunction(unittest func()) {
	brain.unittest = unittest
}

func (brain *s_Brain) SetDumpMemoryFunction(dumpMemory func()) {
	brain.dumpMemory = dumpMemory
}

func (brain *s_Brain) CallLearnFunction(data []byte, firstCell cell.I_CellManagement) {
	if brain.learn != nil {
		brain.learn(data, firstCell)
	}
}

func (brain *s_Brain) CallExecFunction(command string, extraVar ...any) string {
	if brain.exec != nil {
		return brain.exec(command, extraVar...)
	}

	return ""
}

func (brain *s_Brain) CallUnittestFunction() {
	if brain.unittest != nil {
		brain.unittest()
	}
}

func (brain *s_Brain) CallDumpMemoryFunction() {
	if brain.dumpMemory != nil {
		brain.dumpMemory()
	}
}

/*******************
* Globals Varables
*******************/
var g_Brain *s_BrainConfig = createBrain()

/*******************
* createBrain
*******************/
func createBrain() *s_BrainConfig {
	brain := new(s_BrainConfig)
	brain.contextList = make(map[string]*s_Brain)

	exec, err := os.Executable()
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_PROG_NOT_FOUND)
		panic(err)
	}
	brain.mainDirectory = filepath.Dir(exec)
	brain.savesDirectory = filepath.Join(brain.mainDirectory, "save")

	return brain
}

/*******************
* addNewContext
*******************/
func addNewContext(name string) *s_Brain {
	g_Brain.mutex.Lock()
	defer g_Brain.mutex.Unlock()
	brainContext := new(s_Brain)
	g_Brain.contextList[name] = brainContext

	return brainContext
}

/*******************
* init
*******************/
func init() {
	readConfigFile()
	managementBrainContext := CreateBrainContext("__Management__", nil)
	managementBrainContext.SetDumpMemoryFunction(cell.ManagementDumpMemory)
}

/*******************
* initConfigFile
*******************/
func initConfigFile() {
	var err error
	confFile := filepath.Join(g_Brain.savesDirectory, "context.brn")

	// Create or truncate the configuration file
	g_Brain.fileHandle, err = os.OpenFile(confFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_CREATE)
		os.Exit(errcode.ERROR_FATAL_CONFIG_CREATE)
	}
	fileHandle := g_Brain.fileHandle

	// Write file version
	_, err = utils.FileWriteUint32(fileHandle, 0, 0x62726e01)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
		os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
	}
}

/*******************
* readConfigFile
*******************/
func readConfigFile() {
	err := os.MkdirAll(g_Brain.savesDirectory, 0700)
	fmt.Println("Creating saves directory:", g_Brain.savesDirectory)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_CREATE)
		os.Exit(errcode.ERROR_FATAL_CONFIG_OPEN)
	}

	confFile := filepath.Join(g_Brain.savesDirectory, "context.brn")

	// Open the configuration file
	g_Brain.fileHandle, err = os.OpenFile(confFile, os.O_RDWR, 0)
	if err != nil {
		initConfigFile()
	}
	fileHandle := g_Brain.fileHandle

	// Read file version
	var version uint32
	dataOffset, err := utils.FileReadUint32(fileHandle, 0, &version)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
		os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
	}
	if version != 0x62726e01 {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_CRITICAL_CONFIG_INVALID_DATA)
		os.Exit(errcode.ERROR_CRITICAL_CONFIG_INVALID_DATA)
	}

	for {
		// Read the brain context name
		var name string
		dataOffset, err = utils.FileReadString(fileHandle, dataOffset, &name)
		if err != nil {
			fmt.Println("Error reading context.brn:", err)
			if err == io.EOF {
				break
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		// Create the brain context
		brainContext := addNewContext(name)
		if brainContext == nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_PROG_MEMORY_ALLOC)
			os.Exit(errcode.ERROR_FATAL_PROG_MEMORY_ALLOC)
		}
	}
}

/*******************
* GetBrainContext
*******************/
func GetBrainContext(name string) I_Brain {
	g_Brain.mutex.RLock()
	defer g_Brain.mutex.RUnlock()
	brainContext, ok := g_Brain.contextList[name]
	if !ok || brainContext == nil {
		return nil
	}
	return brainContext
}

/*******************
* CreateBrainContext
*******************/
func CreateBrainContext(name string, factory func(I_Brain)) I_Brain {
	brainContext := GetBrainContext(name)

	if brainContext == nil {
		addNewContext(name)
		_, err := utils.FileWriteString(g_Brain.fileHandle, -1, name)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
			return nil
		}

		brainContext = GetBrainContext(name)
		if brainContext == nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_PROG_MEMORY_ALLOC)
			return nil
		}
	}

	// Call the factory function to initialize the brain context
	if factory != nil {
		factory(brainContext)
	}

	return brainContext
}

// Close ferme le handle du fichier de configuration de façon thread-safe.
func Close() {
	g_Brain.mutex.Lock()
	defer g_Brain.mutex.Unlock()
	if g_Brain.fileHandle != nil {
		_ = g_Brain.fileHandle.Close()
		g_Brain.fileHandle = nil
	}
}
