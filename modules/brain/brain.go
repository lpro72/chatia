package brain

/*******************
* Import
*******************/
import (
	"os"
	"sync"

	"chatia/modules/brain/cell"
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
var g_MainBrain *s_BrainConfig = nil
var g_Brain *s_BrainConfig = nil
var g_RegisterBrainContext map[string]func(I_Brain) = make(map[string]func(I_Brain))

/*******************
* addNewContext
*******************/
func addNewContext(brain *s_BrainConfig, name string) *s_Brain {
	brain.mutex.Lock()
	defer brain.mutex.Unlock()
	brainContext := new(s_Brain)
	brain.contextList[name] = brainContext

	return brainContext
}

/*******************
* createBrain
*******************/
func createBrain() *s_BrainConfig {
	brain := new(s_BrainConfig)
	brain.contextList = make(map[string]*s_Brain)

	brain.mutex = sync.RWMutex{}
	brain.fileHandle = nil
	brain.mainDirectory = ""
	brain.savesDirectory = ""

	for name, factory := range g_RegisterBrainContext {
		brainContext := addNewContext(brain, name)
		factory(brainContext)
	}

	return brain
}

/*******************
* registerBrainContext
*******************/
func registerBrainContext(name string, factory func(I_Brain)) {
	g_RegisterBrainContext[name] = factory
}

/*******************
* UseTemporaryBrain
*******************/
func UseTemporaryBrain() {
	g_Brain.mutex.Lock()
	defer g_Brain.mutex.Unlock()

	g_Brain = createBrain()
}

/*******************
* UseMainBrain
*******************/
func UseMainBrain() {
	g_Brain.mutex.Lock()
	defer g_Brain.mutex.Unlock()

	g_Brain = g_MainBrain
}
