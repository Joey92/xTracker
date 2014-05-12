package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

var application *App
var shutdownFlag bool

func main() {

	application = &App{
		Rides:       map[int64]*Ride{},
		RideUpdates: make(chan *Ride),
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, syscall.SIGTERM)
	go func() {
		<-sigchan
		shutdownFlag = true
	}()

	// run main application in goroutine
	go application.run()

	for !shutdownFlag {
		time.Sleep(1 * time.Second)
	}

	application.shutdown()
	os.Exit(1)
}
