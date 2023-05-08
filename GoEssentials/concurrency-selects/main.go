package main

import (
	"fmt"
	"sync"
)

func printWords(in1, in2 chan string, exit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-exit:
			fmt.Println("exiting")
			return
		case str := <-in1:
			fmt.Println("in1: ", str)
		case str := <-in2:
			fmt.Println("in2: ", str)
		}
	}
}

func main() {
	in1 := make(chan string)
	in2 := make(chan string)

	wg := &sync.WaitGroup{}
	exit := make(chan struct{})

	wg.Add(1)
	go printWords(in1, in2, exit, wg)

	in1 <- "hello"
	in2 <- "world"

	close(exit)

	wg.Wait()
}
