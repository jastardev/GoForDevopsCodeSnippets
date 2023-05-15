/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"GoForDevOps/CLI/usingCobra/cmd"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handleSignal(ctx context.Context, cancel context.CancelFunc) chan os.Signal {
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

			select {
			case <-ctx.Done():
				return
			case sig := <-notify:
				switch sig {
				case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
					cancel()
					out <- sig
					return
				default:
					fmt.Println("unhandled signal: ", sig)
				}
			}
		}
	}()
	return out
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	var sigCh chan os.Signal
	go func() {
		handleSignal(ctx, cancel)
	}()

	cmd.Execute(ctx)
	cancel()

	if sig := <-sigCh; sig == syscall.SIGQUIT {
		panic("SIGQUIT")
	}

}
