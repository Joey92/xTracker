package main

/**
 * xTrackHttp Websocket
 * For a live map integration
 */

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
)

func livemap(ws *websocket.Conn) {

	defer ws.Close() // close connection when function ends

	msg, _ := json.Marshal(rides)

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

		var sent []byte
		_, errread := ws.Read(sent)

		if errread != nil {
			return
		}

		// handle command
		retMsg := "success"

		_, errwrite := ws.Write([]byte(retMsg))
		if errwrite != nil {
			return
		}

	}
}
