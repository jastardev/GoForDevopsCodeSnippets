package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	go func() {
		for _, word := range []string{"hello", "world"} {
			ch <- word
		}
		close(ch)
	}()
	for word := range ch {
		fmt.Println(word)
	}
}
