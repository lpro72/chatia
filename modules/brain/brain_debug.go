//go:build debug

package brain

/*******************
* Import
*******************/
import (
    "chatia/modules/error"
)

/*******************
* DumpMemory
*******************/
func DumpMemory(name string) {
    brainContext := GetBrainContext(name)
    
    if brainContext != nil {
        if brainContext.DumpMemory != nil {
            brainContext.DumpMemory()
        }
    } else {
        error.PrintMsgFromErrorCode(error.ERROR_CRITICAL_BRAIN_NOT_FOUND, name)
    }
}


