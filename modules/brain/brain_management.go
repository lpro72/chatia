package brain

/*******************
* Import
*******************/
import (
	"os"
	"path/filepath"

	"chatia/modules/brain/cell"
	"chatia/modules/errcode"
)

/*******************
* ManagementFactory
*******************/
func ManagementFactory(brain I_Brain) {
	brain.SetDumpMemoryFunction(cell.ManagementDumpMemory)
}

/*******************
* init
*******************/
func init() {
	registerBrainContext("__Management__", ManagementFactory)
}

/*******************
* initBrain
*******************/
func InitBrain() {
	g_MainBrain = createBrain()
	g_Brain = g_MainBrain

	exec, err := os.Executable()
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_PROG_NOT_FOUND)
		panic(err)
	}
	g_MainBrain.mainDirectory = filepath.Dir(exec)
	g_MainBrain.savesDirectory = filepath.Join(g_MainBrain.mainDirectory, "save")

	readConfigFile()
}

/*******************
* CloseBrain
*******************/
func CloseBrain() {
	g_MainBrain.mutex.Lock()
	defer g_MainBrain.mutex.Unlock()
	if g_MainBrain.fileHandle != nil {
		_ = g_MainBrain.fileHandle.Close()
		g_MainBrain.fileHandle = nil
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
