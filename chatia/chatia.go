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

	"chatia/modules/errcode"
)

/*******************
* Types
*******************/
type appContext struct {
	isDaemon bool
	execFile string
}

/*******************
* main
*******************/
func main() {
	// Create the context of the execution
	ctx := &appContext{}
	exitCode := parseArguments(ctx)
	if exitCode != errcode.SUCCESS {
		os.Exit(exitCode)
	}

	// Capturer signaux pour shutdown propre
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		// Nettoyage avant arrêt forcé
		closeAll()
		os.Exit(0)
	}()

	// Launch as a daemon
	switch {
	case ctx.isDaemon:
		initAll()

		exitCode = execAsDaemon()
	case ctx.execFile != "":
		exitCode = execFile(ctx.execFile)
	default:
		exitCode = execShell()
	}

	// Exit
	closeAll()
	os.Exit(exitCode)
}

/*******************
* parseArguments
*******************/
func parseArguments(ctx *appContext) int {
	flag.BoolVar(&ctx.isDaemon, "daemon", false, "Launch as a daemon")
	flag.StringVar(&ctx.execFile, "exec", "", "Execute a file and exit")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExamples:")
		fmt.Fprintf(os.Stderr, "  %s -daemon\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -exec <file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s\n", os.Args[0])
	}

	flag.Parse()

	return errcode.SUCCESS
}
