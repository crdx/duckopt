# duckopt

A docopt wrapper library for Go.

## Usage

```go
import (
    "fmt"
    "github.com/crdx/duckopt"
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
    var opts Opts
    if err := duckopt.Parse(getUsage(), "$0").Bind(&opts); err != nil {
        panic(err)
    }
    fmt.Println(opts)
}
```
