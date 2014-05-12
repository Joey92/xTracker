package main

type clientUpdater struct {
}

func (updater *clientUpdater) update(ride *Ride) bool {

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
