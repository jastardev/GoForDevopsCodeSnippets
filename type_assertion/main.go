package main

import "fmt"

func printInterface(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("i was %d\n", i)
	case string:
		fmt.Printf("i was %s\n", i)
	case float64:
		fmt.Printf("i was %v\n", i)
	default:
		fmt.Printf("i was an unsupported type %T\n", i)
	}
}

func main() {
	var i interface{}

	i = 3
	printInterface(i)
	i = "hello World"
	printInterface(i)
	i = 3.3
	printInterface(i)
}
