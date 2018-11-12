package main

import (
	"os"
)

const VERSION = "0.1.0"

func main() {
	var cli CLI = NewCLI()
	var ret = cli.Run(os.Args)
	os.Exit(ret)
}
