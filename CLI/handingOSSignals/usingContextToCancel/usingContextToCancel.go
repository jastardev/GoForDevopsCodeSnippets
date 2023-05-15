package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

func cleanup(tmpFiles string) {
	if err := os.RemoveAll(tmpFiles); err != nil {
		fmt.Println("problem doing file cleanup: ", err)
		return
	}
	fmt.Println("Cleanup done")
}

func createFiles(ctx context.Context, tmpFiles string) error {
	for i := 0; i < 30; i++ {
		if err := ctx.Err(); err != nil {
			return ctx.Err()
		}
		_, err := os.Create(filepath.Join(tmpFiles, strconv.Itoa(i)))
		if err != nil {
			return err
		}
		fmt.Println("Created file: ", i)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func handleSignal(cancel context.CancelFunc) chan os.Signal {
	out := make(chan os.Signal, 1)
	notify := make(chan os.Signal, 10)
	signal.Notify(
		notify,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		defer close(out)
		for {
			sig := <-notify
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				cancel()
				out <- sig
				return
			default:
				fmt.Println("unhandled signal: ", sig)
			}
		}
	}()
	return out
}

func main() {
	tmpFiles, err := os.MkdirTemp("", "myApp_*")
	if err != nil {
		log.Println("error creating temp file directory: ", err)
		os.Exit(1)
	}
	fmt.Println("temp files lcoated at: ", tmpFiles)

	ctx, cancel := context.WithCancel(context.Background())
	recvsig := handleSignal(cancel)

	if err := createFiles(ctx, tmpFiles); err != nil {
		cleanup(tmpFiles)
		select {
		case sig := <-recvsig:
			if sig == syscall.SIGQUIT {
				panic("SIGQUIT called")
			}
		default:
			// prevents waiting on a signal if none exists.
		}
		log.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Println("Done")
}
