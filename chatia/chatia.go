package main

/*******************
* Import
*******************/
import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"chatia/modules/brain"
	"chatia/modules/errcode"
)

/*******************
* Types
*******************/
type AppContext struct {
	IsDaemon       bool
	LaunchUnitTest bool
}

/*******************
* main
*******************/
func main() {
	// Create the context of the execution
	ctx := &AppContext{}
	exitCode := parseArguments(ctx)
	if exitCode != errcode.SUCCESS {
		os.Exit(exitCode)
	}

	brain.InitBrain()

	// capturer signaux pour shutdown propre
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		// nettoyage avant arrêt forcé
		brain.CloseBrain()
		os.Exit(0)
	}()

	// Launch as a daemon
	switch {
	case ctx.IsDaemon:
		exitCode = ExecAsDaemon()
	case ctx.LaunchUnitTest:
		exitCode = brain.Unittest()
	default:
		exitCode = ExecShell()
	}

	// Exit
	brain.CloseBrain()
	os.Exit(exitCode)
}

/*******************
* parseArguments
*******************/
func parseArguments(ctx *AppContext) int {
	flag.BoolVar(&ctx.IsDaemon, "daemon", false, "Launch as a daemon")
	flag.BoolVar(&ctx.LaunchUnitTest, "unittest", false, "Launch unittest")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExamples:")
		fmt.Fprintf(os.Stderr, "  %s -daemon\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -unittest\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s\n", os.Args[0])
	}

	flag.Parse()

	// Refuse incompatible flags
	if ctx.IsDaemon && ctx.LaunchUnitTest {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_CRITICAL_COMMAND_INVALID_DATA, "--daemon", "--unittest")
		flag.Usage()
		return errcode.ERROR_CRITICAL_COMMAND_INVALID_DATA
	}

	return errcode.SUCCESS
}
