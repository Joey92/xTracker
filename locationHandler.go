package main

import (
	"net/http"
	"strconv"
)

func locationHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	lat, errLat := strconv.ParseFloat(r.FormValue("lat"), 64)
	lng, errLng := strconv.ParseFloat(r.FormValue("lng"), 64)
	time, errTime := strconv.ParseInt(r.FormValue("time"), 10, 64)
	rideID, errRide := strconv.ParseInt(r.FormValue("ride"), 10, 64)
	source := r.FormValue("source")

	if errLat != nil || errLng != nil || errRide != nil || errTime != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	p := newPoint(lat, lng, time, source)

	_, exists := rides[rideID]

	if exists {
		rides[rideID].addPoint(p)
	} else {
		rides[rideID] = newRide(rideID)
		rides[rideID].addPoint(p)
	}

	// add to rideUpdates channel to notify clients using the livemap about the change
	rideUpdates <- rides[rideID]
}
