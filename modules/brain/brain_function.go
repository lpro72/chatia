package brain

/*******************
* Import
*******************/
import (
    "chatia/modules/error"
)

/*******************
* LearnFromString
*******************/
func LearnFromString(data string, brainName string) {
    Learn([]byte(data), brainName)
}

/*******************
* Learn
*******************/
func Learn(data []byte, brainName string) {
    brainContext := GetBrainContext(brainName)
    if brainContext != nil {
        if brainContext.FirstCell != nil && brainContext.Learn != nil {
            brainContext.Learn(data, brainContext.FirstCell)
        }
    } else {
        error.PrintMsgFromErrorCode(error.ERROR_CRITICAL_BRAIN_NOT_FOUND, brainName)
    }
}

/*******************
* Exec
*******************/
func Exec(command string, brainName string, extraVar ...any) string {
    brainContext := GetBrainContext(brainName)
    if brainContext != nil {
        if brainContext.Exec != nil {
            return brainContext.Exec(command, extraVar...)
        }
    } else {
        error.PrintMsgFromErrorCode(error.ERROR_CRITICAL_BRAIN_NOT_FOUND, brainName)
    }
    
    return ""
}


