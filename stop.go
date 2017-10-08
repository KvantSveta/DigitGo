package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func stopProgram(runService *bool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
	*runService = false
}