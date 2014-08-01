package main

/**
 * xTrackHttp Websocket
 * For a live map integration
 */

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"time"
)

func livemap(ws *websocket.Conn) {

	defer ws.Close() // close connection when function ends

	var currentRides []*Ride

	for _, ride := range rides {
		currentRides = append(currentRides, ride)
	}

	msg, errMarshal := json.Marshal(currentRides)

	currentRides = nil

	if errMarshal != nil {
		fmt.Println(errMarshal.Error())
	}

	// send all current rides
	_, errwrite := ws.Write(msg)
	if errwrite != nil {
		fmt.Println("connection closed at initial connect")
		return
	}

	msg = nil

	// add to clients list to receive updates
	livemapClients = append(livemapClients, ws)

	for true {

		time.Sleep(20 * time.Second)

		// send all current rides
		_, errPing := ws.Write([]byte("ping"))
		if errPing != nil {
			fmt.Println("Connection to a client closed")
			return
		}

	}
}
