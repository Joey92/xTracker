package main

import (
	"time"
)

func rideCleanup() {

	for true {
		time.Sleep(5 * time.Minute)

		for id, ride := range rides {
			if ride.UpdateTime < (time.Now().Unix() - 45*60) {
				ride.flush()
				delete(rides, id)
			}
		}
	}
}
