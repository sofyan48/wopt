package app

import (
	"fmt"

	"github.com/meongbego/wopt/app/clis"

	"github.com/docopt/docopt-go"
	"github.com/meongbego/wopt/app/utils"
)

// RoutesCommand |  Function Object
var (
	RoutesCommand = map[string]interface{}{
		"test": func() { clis.Test() },
	}
	funcs          = utils.NewFuncs(100)
	version string = "0.0.1"
)

// BindCommand | Bindd allcommand
func BindCommand() {
	for k, v := range RoutesCommand {
		err := funcs.Bind(k, v)
		if k[:3] == "err" {
			if err == nil {
				fmt.Println(fmt.Sprintf("Bind %s: %s", k, "an error should be paniced."))
			}
		} else {
			if err != nil {
				fmt.Println(fmt.Sprintf("Bind %s: %s", k, err))
			}
		}
	}
}

// Init Function to load all cli packages
func Init() {
	usage := `Example of program with many options using docopt.
Usage:
	wopt  COMMAND
Arguments:
	COMMAND  				Command Name
Options:
	-h --help            show this help message and exit
	--version            show version and exit`

	// bind command
	BindCommand()

	// Parsing Docopt Arguments
	arguments, _ := docopt.ParseArgs(usage, nil, version)
	cmd := fmt.Sprintf("%s", arguments["COMMAND"])
	funcs.Call(string(cmd))

}
