package main

/*******************
* Import
*******************/
import (
        "os"
        "flag"
        "chatia/modules/error"
        "chatia/modules/brain"
)

/*******************
* Types
*******************/
type ContextStruct struct {
    IsDaemon bool
    LaunchUnitTest bool
}

/*******************
* main
*******************/
func main() {
    // Create the context of the execution
    context := new(ContextStruct)
    parseArguments(context)
    
    // Launch as a daemon
    if context.IsDaemon {
        ExecAsDaemon(context)
        os.Exit(error.SUCCESS)
    } else if context.LaunchUnitTest {
        brain.Unittest()
        os.Exit(error.SUCCESS)
    }

    os.Exit(error.SUCCESS)
}

/*******************
* parseArguments
*******************/
func parseArguments(context *ContextStruct) {
    flag.BoolVar(&context.IsDaemon, "daemon", false, "Launch as a daemon") 
    flag.BoolVar(&context.LaunchUnitTest, "unittest", false, "Launch unittest")
    flag.Parse()
}

func ExecAsDaemon(context *ContextStruct) {
}

