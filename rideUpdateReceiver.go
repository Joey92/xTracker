package main

type rideUpdateReceiver interface {
	update(ride *Ride)
}
