package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/gertd/go-pluralize/cmd/pluralize/version"
)

const (
	appName = "pluralize"
)

func main() {

	var (
		word        = flag.String("word", "", "input value")
		cmd         = flag.String("cmd", "All", "command [All|IsPlural|IsSingular|Plural|Singular]")
		showVersion = flag.Bool("version", false, "display version info")
	)

	flag.Parse()

	if showVersion != nil && *showVersion {
		displayVersionInfo(appName)
		return
	}

	if word == nil || len(*word) == 0 {
		fmt.Printf("-word not specified\n")
		return
	}

	pluralize := pluralize.NewClient()

	testCmd := TestCmdString(*cmd)
	if testCmd.Has(TestCmdUnknown) {
		fmt.Printf("Unknown -cmd value\nOptions: [All|IsPlural|IsSingular|Plural|Singular]\n")
		return
	}

	if testCmd.Has(TestCmdIsPlural) {
		fmt.Printf("IsPlural(%s)   => %t\n", *word, pluralize.IsPlural(*word))
	}
	if testCmd.Has(TestCmdIsSingular) {
		fmt.Printf("IsSingular(%s) => %t\n", *word, pluralize.IsSingular(*word))
	}
	if testCmd.Has(TestCmdPlural) {
		fmt.Printf("Plural(%s)     => %s\n", *word, pluralize.Plural(*word))
	}
	if testCmd.Has(TestCmdSingular) {
		fmt.Printf("Singular(%s)   => %s\n", *word, pluralize.Singular(*word))
	}
}

func displayVersionInfo(name string) {

	vi := version.GetInfo()
	fmt.Fprintf(os.Stdout, "%s - %s@%s [%s].[%s].[%s]\n",
		name,
		vi.Version,
		vi.Commit,
		vi.Date,
		runtime.GOOS,
		runtime.GOARCH,
	)
}
