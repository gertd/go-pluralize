package main

import (
	"flag"
	"fmt"
	"os"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/gertd/go-pluralize/pkg/tflags"
	"github.com/gertd/go-pluralize/pkg/version"
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

	testCmd := tflags.TestCmdString(*cmd)
	if testCmd.Has(tflags.TestCmdUnknown) {
		fmt.Printf("Unknown -cmd value\nOptions: [All|IsPlural|IsSingular|Plural|Singular]\n")
		return
	}

	if testCmd.Has(tflags.TestCmdIsPlural) {
		fmt.Printf("IsPlural(%s)   => %t\n", *word, pluralize.IsPlural(*word))
	}

	if testCmd.Has(tflags.TestCmdIsSingular) {
		fmt.Printf("IsSingular(%s) => %t\n", *word, pluralize.IsSingular(*word))
	}

	if testCmd.Has(tflags.TestCmdPlural) {
		fmt.Printf("Plural(%s)     => %s\n", *word, pluralize.Plural(*word))
	}

	if testCmd.Has(tflags.TestCmdSingular) {
		fmt.Printf("Singular(%s)   => %s\n", *word, pluralize.Singular(*word))
	}
}

func displayVersionInfo(name string) {
	fmt.Fprintf(os.Stdout, "%s - %s\n",
		name,
		version.GetInfo(),
	)
}
