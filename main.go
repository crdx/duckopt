package duckopt

import (
	"fmt"
	"crdx.org/hereduck"
	"github.com/docopt/docopt-go"
	"log"
	"os"
	"path"
	"strings"
)

func Parse(usage string, programNamePlaceholder string) docopt.Opts {
	usage = hereduck.D(usage)
	usage = strings.Replace(usage, programNamePlaceholder, path.Base(os.Args[0]), -1)

	parser := &docopt.Parser{
		HelpHandler: func(_ error, _ string) {
			fmt.Print(usage)
			os.Exit(2)
		},
	}

	opts, err := parser.ParseArgs(usage, os.Args[1:], "")

	if err != nil {
		log.Fatal("duckopt parse error: " + err.Error())
	}

	return opts
}
