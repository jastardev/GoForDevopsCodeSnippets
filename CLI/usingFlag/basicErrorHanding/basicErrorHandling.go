package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	useProd = flag.Bool("prod", false, "User a production endpoint")
	useDev  = flag.Bool("dev", false, "Use a development endpoint")
	help    = flag.Bool("help", false, "Display help text")
)

func main() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	switch {
	case *useProd && *useDev:
		fmt.Println("Error: --prod and --dev cannot both be set")
		flag.PrintDefaults()
		os.Exit(1)
	case !(*useProd || *useDev):
		fmt.Println("Error: either --prod or --dev must be set")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
