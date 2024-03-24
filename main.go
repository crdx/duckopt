package duckopt

import (
	"fmt"
	"os"
	"path"
	"strings"

	"crdx.org/hereduck"
	"github.com/docopt/docopt-go"
)

// Bind parses a usage string and returns an instance of *T, or an error if it can't be parsed.
//
// placeholder is an optional placeholder in the usage string that should be replaced with the name
// of the program.
func Bind[T any](usage string, placeholder ...string) (*T, error) {
	var opts T
	if err := MustParse(usage, placeholder...).Bind(&opts); err != nil {
		return nil, err
	}
	return &opts, nil
}

// MustBind parses a usage string and returns an instance of *T, or panics if it can't be parsed.
//
// placeholder is an optional placeholder in the usage string that should be replaced with the name
// of the program.
func MustBind[T any](usage string, placeholder ...string) *T {
	opts, err := Bind[T](usage, placeholder...)
	if err != nil {
		panic(err)
	}
	return opts
}

// Parse parses a usage string and returns an instance of [docopt.Opts], or an error if it can't be
// parsed.
//
// placeholder is an optional placeholder in the usage string that should be replaced with the name
// of the program.
func Parse(usage string, placeholder ...string) (docopt.Opts, error) {
	usage = hereduck.D(usage)
	usage = strings.ReplaceAll(usage, "\t", "    ")

	if len(placeholder) > 0 {
		usage = strings.Replace(usage, placeholder[0], path.Base(os.Args[0]), -1)
	}

	parser := &docopt.Parser{
		HelpHandler: func(_ error, _ string) {
			fmt.Print(usage)
			os.Exit(2)
		},
	}

	return parser.ParseArgs(usage, os.Args[1:], "")
}

// MustParse parses a usage string and returns an instance of [docopt.Opts], or panics if it can't
// be parsed.
//
// placeholder is an optional placeholder in the usage string that should be replaced with the name
// of the program.
func MustParse(usage string, placeholder ...string) docopt.Opts {
	opts, err := Parse(usage, placeholder...)
	if err != nil {
		panic(err)
	}
	return opts
}
