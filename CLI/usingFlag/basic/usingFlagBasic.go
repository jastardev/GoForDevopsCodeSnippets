package main

import (
	"flag"
	"fmt"
)

var (
	endpoint = flag.String(
		"endpoint",
		"myserver.aws.com",
		"The server this app will connect to",
	)
)

func main() {
	flag.Parse()
	fmt.Println("server endpoint is: ", *endpoint)
}
