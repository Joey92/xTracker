package main

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var shutdownFlag bool
var rides map[int64]*Ride
var rideUpdates chan *Ride
var livemapClients []*websocket.Conn

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, syscall.SIGTERM)
	go func() {
		<-sigchan
		shutdownFlag = true
	}()

	rides = map[int64]*Ride{}
	rideUpdates = make(chan *Ride, 0)

	http.Handle("/live", websocket.Handler(livemap))
	http.HandleFunc("/location", locationHandler)

	// run main application in goroutine
	go http.ListenAndServe(":8080", nil)

	go updateClients()

	go rideCleanup()

	for !shutdownFlag {
		time.Sleep(1 * time.Second)
	}

	shutdown()

	os.Exit(1)
}
