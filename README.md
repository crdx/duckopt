# duckopt

**duckopt** is a docopt wrapper library for Go.

## Installation

```sh
go get crdx.org/duckopt/v2
```

## Usage

```go
import (
    "fmt"

    "crdx.org/duckopt"
)

func getUsage() string {
    return `
        Usage:
            $0 [options] command <arg>

        Options:
            --dry-run         Do a dry run
            -v, --verbose     Be verbose
            -C, --no-color    Disable colours
            -h, --help        Show help
    `
}

type Opts struct {
    Command bool   `docopt:"command"`
    Arg     string `docopt:"<arg>"`
    DryRun  bool   `docopt:"--dry-run"`
    Verbose bool   `docopt:"--verbose"`
    NoColor bool   `docopt:"--no-color"`
}

func main() {
    opts := duckopt.MustBind[Opts](getUsage(), "$0")
    fmt.Printf("%+v\n", opts)
}
```

## Contributions

Open an [issue](https://github.com/crdx/duckopt/issues) or send a [pull request](https://github.com/crdx/duckopt/pulls).

## Licence

[GPLv3](LICENCE).
