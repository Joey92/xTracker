package main

import (
	"time"
)

type App struct {
	Rides           map[int64]*Ride
	RideUpdates     chan *Ride
	UpdateReceivers []*rideUpdateReceiver
}

func (app *App) shutdown() {

	for id, ride := range app.Rides {
		ride.flush()
		delete(app.Rides, id)
	}
}

func (app *App) getRides() map[int64]*Ride {
	return app.Rides
}

func (app *App) getRideById(id int64) *Ride {

	ride, exists := app.Rides[id]

	if exists {
		return app.Rides[id]
	}

	return nil
}

func (app *App) addRide(r *Ride) *Ride {

	ride, exists := app.Rides[r.Id]

	if exists {
		return app.Rides[r.Id]
	}

	app.Rides[r.Id] = r

	return app.Rides[r.Id]
}

func (app *App) removeRide(id int64) bool {

	ride, exists := app.Rides[id]

	if !exists {
		return false
	}

	delete(app.Rides, id)
	return true
}

func (app *App) cleanUp() {
	for id, ride := range app.Rides {
		if ride.UpdateTime < (time.Now().Unix() - 45*60) {
			app.removeRide(id)
		}
	}
}

func (app *App) addUpdateReceiver(receiver *rideUpdateReceiver) {
	app.UpdateReceivers = append(app.UpdateReceivers, receiver)
}

func (app *App) run() int {

	for true {
		ride := <-app.RideUpdates

		for _, receiver := range app.UpdateReceivers {
			receiver.update(ride)
		}

		if ride.hasTag("Deleted") {
			app.removeRide(ride.Id)
		} else {
			app.addRide(ride)
		}
	}

	return -1
}
