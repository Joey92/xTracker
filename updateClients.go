package main

import (
	"encoding/json"
	"fmt"
)

func updateClients() {

	for true {

		ride := <-rideUpdates

		responseString, err := json.Marshal(ride)

		if err != nil {
			fmt.Println("Ride could not be encoded:", err.Error())
		}

		for index, ws := range livemapClients {
			_, errwrite := ws.Write(responseString)
			if errwrite != nil {
				// error happened. dump the client
				ws.Close()
				livemapClients = append(livemapClients[:index], livemapClients[index+1:]...)
			}
		}
	}
}
