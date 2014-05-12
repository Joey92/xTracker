package main

import (
	"math"
)

func distance(p1 *Point, p2 *Point) float64 {

	var earthRadius float64
	var dLon, dLat, a, c, d float64

	earthRadius = 6371 //earth radius in kilometers

	dLat = (p2.Lat - p1.Lat) * math.Pi / 180
	dLon = (p2.Lng - p1.Lng) * math.Pi / 180

	a = math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos((p1.Lat*math.Pi/180))*math.Cos((p2.Lat*math.Pi/180))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c = 2 * math.Asin(math.Sqrt(a))
	d = earthRadius * c

	return d
}

func speed(p1 *Point, p2 *Point, timediff float64) float64 {

	if timediff <= 0 {
		return 0
	}

	distance := distance(p1, p2)

	speed := (distance / timediff * 3600)

	if speed > 0 {
		return speed
	}

	return 0
}
