package main

import (
	"os"
)

const VERSION = "1.0.0"

func main() {
	cli := NewCLI()
	ret := cli.Run(os.Args)
	os.Exit(ret)
}
