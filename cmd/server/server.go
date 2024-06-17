package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1"

func main() {
	// допишите код
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of ????????????????/%s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}
