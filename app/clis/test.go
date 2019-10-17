package clis

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

// Test View Action Test Lib in Command
func Test() {
	usage := `Example of program with many options using docopt.
Usage:
  options [-hf NAME] [--select=ERRORS | --ignore=ERRORS] PATH...
Arguments:
    PATH               Arguments
Options:
  -h --help            show this help message and exit
  --version            show version and exit
  -f NAME --file=NAME  when parsing directories, only check filenames matching
                       these comma separated patterns [default: *.go]
  --select=ERRORS      select errors and warnings (e.g. E,W6)
  --ignore=ERRORS      skip errors and warnings (e.g. E4,W)`

	arguments, _ := docopt.ParseArgs(usage, nil, "1.0.0")
	path := arguments["PATH"]
	fmt.Println(path)
}
