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

    "crdx.org/duckopt/v2"
)

func getUsage() string {
    return `
        Usage:
            $0 [options] run <task>

        Options:
            -v, --verbose    Be verbose
    `
}

type Opts struct {
    Command bool   `docopt:"run"`
    Arg     string `docopt:"<task>"`
    Verbose bool   `docopt:"--verbose"`
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
