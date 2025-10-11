package brain

/*******************
* Import
*******************/
import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"chatia/modules/errcode"
	"chatia/modules/utils"
)

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
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading context.brn:", err)
			errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
			os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
		}

		// Create the brain context
		if _, exists := g_RegisterBrainContext[name]; !exists {
			errcode.PrintMsgFromErrorCode(errcode.WARNING_CONFIG_NOT_FOUND, name)
			os.Exit(errcode.WARNING_CONFIG_NOT_FOUND)
		}
	}
}
