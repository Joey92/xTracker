package main

import (
	"net/http"
)

var webHandler *http.ServeMux = http.NewServeMux()

var webServer = &http.Server{
	Addr:    ":8080",
	Handler: webHandler,
}

func webServerInit() {

	// register handlers
	webHandler.HandleFunc("/location", locationHandler)

	go webServer.ListenAndServe()
}
