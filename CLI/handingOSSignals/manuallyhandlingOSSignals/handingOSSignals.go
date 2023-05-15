package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup() {
	fmt.Println("Cleaning up some files")
	fmt.Println("Cancelling remote called.")
}

func main() {

	signals := make(chan os.Signal, 1)

	signal.Notify(
		signals,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		switch <-signals {
		case syscall.SIGINT, syscall.SIGTERM:
			fmt.Println("Keyboard Interrupt called...")
			cleanup()
			os.Exit(1)
		case syscall.SIGQUIT:
			cleanup()
			panic("SIGQUIT called")
		}
	}()

	fmt.Println("Sleeping for 30 seconds to simulate a long running job.")
	time.Sleep(30 * time.Second)
}
