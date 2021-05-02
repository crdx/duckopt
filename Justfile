_help:
    @just --list --unsorted

upgrade:
    go get github.com/crdx/hereduck
    go get github.com/docopt/docopt-go
    go mod tidy
