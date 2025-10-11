package brain

import "chatia/modules/errcode"

/*******************
* Import
*******************/

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
		firstCell := brainContext.GetFirstCell()
		if firstCell != nil {
			brainContext.CallLearnFunction(data, firstCell)
		}
	} else {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_CRITICAL_BRAIN_NOT_FOUND, brainName)
	}
}

/*******************
* Exec
*******************/
func Exec(command string, brainName string, extraVar ...any) string {
	brainContext := GetBrainContext(brainName)
	if brainContext != nil {
		return brainContext.CallExecFunction(command, extraVar...)
	} else {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_CRITICAL_BRAIN_NOT_FOUND, brainName)
	}

	return ""
}

/*******************
* Unittest
*******************/
func Unittest() int {
	for _, brainContext := range g_Brain.contextList {
		if brainContext != nil {
			brainContext.CallUnittestFunction()
		}
	}

	return errcode.SUCCESS
}

/*******************
* DumpMemory
*******************/
func DumpMemory(name string) {
	brainContext := GetBrainContext(name)

	if brainContext != nil {
		brainContext.CallDumpMemoryFunction()
	} else {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_CRITICAL_BRAIN_NOT_FOUND, name)
	}
}
