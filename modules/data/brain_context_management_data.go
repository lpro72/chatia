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
* S_BrainContext
*******************/
type S_BrainContextManagement struct {
	contextList map[string]interfaces.I_BrainContext
	loaded      bool

	// File management
	fileHandle *os.File

	MemoryAccess interfaces.I_Lock
}

/*******************
* Internal functions
*******************/
func (brainContextManagement *S_BrainContextManagement) initialize(brainConfig interfaces.I_BrainConfig) {
	brainContextManagement.contextList = make(map[string]interfaces.I_BrainContext)
	brainContextManagement.MemoryAccess = &utils.S_Lock{}
	brainContextManagement.fileHandle = utils.ReadConfigFile(brainConfig, "context.brn", brainContextManagement.LoadFromFile)
	brainContextManagement.loaded = true
}

func (brainContextManagement *S_BrainContextManagement) appendToFile(context interfaces.I_BrainContext) {
	if brainContextManagement.fileHandle == nil || brainContextManagement.loaded == false {
		return
	}

	dataOffset, err := utils.FileGetEndOffset(brainContextManagement.fileHandle)
	if err != nil {
		return
	}
	context.SetFileOffset(dataOffset)

	// Write the brain context name
	name := context.GetName()
	_, err = utils.FileWriteString(brainContextManagement.fileHandle, -1, name)
	if err != nil {
		return
	}

	// Write the first cell ID
	firstCellID := context.GetFirstCellID()
	_, err = utils.FileWriteUint32(brainContextManagement.fileHandle, -1, firstCellID)
	if err != nil {
		return
	}
}

/*******************
*  Functions for the interface I_BrainContextManagement
*******************/
func (brainContextManagement *S_BrainContextManagement) AddNewBrainContext(name string, context interfaces.I_BrainContext) {
	brainContextManagement.Lock()
	defer brainContextManagement.Unlock()

	brainContextManagement.contextList[name] = context
	brainContextManagement.appendToFile(context)
}

func (brainContextManagement *S_BrainContextManagement) GetBrainContext(name string) interfaces.I_BrainContext {
	brainContextManagement.Lock()
	defer brainContextManagement.Unlock()

	return brainContextManagement.contextList[name]
}

func (brainContextManagement *S_BrainContextManagement) CreateNewBrainContext(brain interfaces.I_BrainConfig, name string, firstSynapseID uint32) interfaces.I_BrainContext {
	brainContext := new(S_BrainContext)
	brainContext.Initialize(brain, name, firstSynapseID)
	brainContextManagement.AddNewBrainContext(name, brainContext)

	return brainContext
}

func (brainContextManagement *S_BrainContextManagement) UpdateToFile(context interfaces.I_BrainContext) {
	if brainContextManagement.fileHandle == nil || brainContextManagement.loaded == false {
		return
	}

	dataOffset := context.GetFileOffset()

	// Write the brain context name
	name := context.GetName()
	dataOffset, err := utils.FileWriteString(brainContextManagement.fileHandle, dataOffset, name)
	if err != nil {
		return
	}

	// Write the first cell ID
	firstCellID := context.GetFirstCellID()
	dataOffset, err = utils.FileWriteUint32(brainContextManagement.fileHandle, dataOffset, firstCellID)
	if err != nil {
		return
	}
}

/*******************
*  Functions for the interface I_File
*******************/
func (brainContextManagement *S_BrainContextManagement) LoadFromFile(fileHandle *os.File, dataOffset int64, brainConfigInterface interfaces.I_BrainConfig, version uint32) {
	for {
		// Read the brain context name
		var name string
		var err error
		contextDataOffset := dataOffset

		dataOffset, err = utils.FileReadString(fileHandle, dataOffset, &name)
		if err != nil {
			if err == io.EOF {
				break
			}
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		// Read first synapse ID
		synapseID := uint32(0)
		dataOffset, err = utils.FileReadUint32(fileHandle, dataOffset, &synapseID)
		if err != nil {
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}
		// Create the brain context
		brainContext := brainContextManagement.CreateNewBrainContext(brainConfigInterface, name, synapseID)
		brainContext.SetFileOffset(contextDataOffset)
	}
}

func (brainContextManagement *S_BrainContextManagement) Close() {
	brainContextManagement.Lock()
	defer brainContextManagement.Unlock()

	utils.CloseFile(brainContextManagement.fileHandle)
	brainContextManagement.fileHandle = nil
}

/*******************
* Functions for the interface I_Lock
*******************/
// Lock
func (brainContextManagement *S_BrainContextManagement) Lock() {
	brainContextManagement.MemoryAccess.Lock()
}

// Unlock
func (brainContextManagement *S_BrainContextManagement) Unlock() {
	brainContextManagement.MemoryAccess.Unlock()
}

/*******************
* BrainContextManagement_Create
*******************/
func BrainContextManagement_Create(brainConfig interfaces.I_BrainConfig) interfaces.I_BrainContextManagement {
	brainContextManagement := new(S_BrainContextManagement)
	brainContextManagement.initialize(brainConfig)

	return brainContextManagement
}

/*******************
* Global Variables
*******************/
var g_RegisterBrainContext map[string]func(interfaces.I_BrainContext) = make(map[string]func(interfaces.I_BrainContext))

/*******************
* BrainContextManagement_RegisterNewContext
*******************/
func BrainContextManagement_RegisterNewContext(name string, factory func(interfaces.I_BrainContext)) {
	g_RegisterBrainContext[name] = factory
}
