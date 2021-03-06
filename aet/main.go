package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	flags    = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	commands = map[string]func(){
		"init": initSourcesCommand,
		"test": runTestsCommand,
	}
	// Expect appengine-go source files (repo) to be int appengineDir/src
	appengineDir string
)

func main() {
	parseArgs()
	action := strings.ToLower(flags.Arg(0))
	cmd, ok := commands[action]
	if !ok {
		log.Fatalf("Unknown action %q", action)
	}
	cmd()
}

func parseArgs() {
	flags.Parse(os.Args[1:])
	if flags.NArg() == 0 {
		flags.Usage()
		log.Fatal("Too few arguments")
	}
}

func init() {
	defAppengineDir := os.Getenv("APPENGINE_GO_SRC")
	if defAppengineDir == "" {
		gop := strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))[0]
		defAppengineDir = filepath.Join(gop, "src", "appengine-go")
	}
	flags.StringVar(&appengineDir, "d", defAppengineDir,
		"expect appengine-go sources to be in d/src; required")

	flags.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage: %s [flags] {init|test} ./path/to/*_test.go\n", os.Args[0])
		flags.PrintDefaults()
	}
}
