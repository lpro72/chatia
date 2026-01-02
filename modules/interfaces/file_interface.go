package interfaces

/*******************
* Imports
*******************/
import "os"

/*******************
* Interfaces
*******************/
type I_File interface {
	LoadFromFile(fileHandle *os.File, dataOffset int64, brainConfigInterface I_BrainConfig, version uint32)
	Close()
}
