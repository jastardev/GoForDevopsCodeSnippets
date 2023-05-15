package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	useProd = flag.Bool("prod", false, "Use a production endpoint")
	useDev  = flag.Bool("dev", false, "Use a development endpoint")
	help    = new(bool)
)

func init() {
	flag.BoolVar(help, "help", false, "Display help text")
	flag.BoolVar(help, "h", false, "Display help text (shorthand)")
}

func main() {
	flag.Parse()

	authors := flag.Args()
	if len(authors) == 0 {
		fmt.Println("did not pass any authors")
		os.Exit(1)
	}
}
