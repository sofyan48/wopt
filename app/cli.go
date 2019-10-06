package app

import (
	"fmt"
	"os"

	"github.com/meongbego/wopt/app/clis"
)

func Init() {
	argsWithProg := os.Args[1]
	if argsWithProg == "test" {
		clis.Test()
		os.Exit(1)
	} else {
		fmt.Println("No Commands")
		os.Exit(1)
	}
}
