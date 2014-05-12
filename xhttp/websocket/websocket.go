package main

import (
	"code.google.com/p/go.net/websocket"
)

var livemapClients []*websocket.Conn

func websocketInit(app *App) {

	var websocketReceiver = &clientUpdater{}
	app.addUpdateReceiver(websocketReceiver)

	webHandler.Handle("/live", websocket.Handler(livemap))
}
