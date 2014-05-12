package main

import (
	"fmt"
	"time"
)

type Ride struct {
	Id         int64             `json:"id"`
	Routes     map[string]*Route `json:"-"`
	Calendar   string            `json:"-"`
	Location   *Point            `json:"location"`
	Tags       []string          `json:"tags`
	Speed      float64           `json:"speed"`
	UpdateTime int64             `json:"timestamp"`
	UpdateDiff int64             `json:"updateDiff"`
}

func newRide(id int64) *Ride {
	ride := &Ride{
		Id:         id,
		Routes:     map[string]*Route{},
		Tags:       []string{},
		Location:   nil,
		Speed:      0,
		UpdateTime: 0,
		UpdateDiff: 0,
	}

	return ride
}

func (r *Ride) getRoute(name string) (bool, *Route) {
	_, exists := r.Routes[name]

	if exists {
		return true, r.Routes[name]
	}

	return false, nil
}

func (r *Ride) addRoute(name string, route *Route) *Route {

	r.Routes[name] = route

	return route
}

func (r *Ride) getRoutes() map[string]*Route {
	return r.Routes
}

func (r *Ride) getPosition() *Point {
	return r.Location
}

func (r *Ride) setPosition(p *Point) {
	r.Location = p
}

func (r *Ride) addPoint(p *Point) {
	source := p.getSource()
	_, exists := r.Routes[source]
	r.Location = p

	if exists {
		pointAmount := len(r.Routes[source].points)
		pointBefore := r.Routes[source].getPoint(pointAmount - 1)

		timediff := float64(p.Timestamp - pointBefore.Timestamp)

		if timediff > 0 {
			r.Speed = speed(pointBefore, p, timediff)
		}

		r.Routes[source].addPoint(p)
	} else {
		r.Routes[source] = newRoute(p)
	}
	now := time.Now().Unix()
	r.UpdateDiff = now - r.UpdateTime
	r.UpdateTime = now
}

func (r *Ride) flush() {
	// save some stuff
	fmt.Println("Flushed ride ", r.Id)
}

func (r *Ride) addTag(tagToAdd string) {
	for _, tag := range r.Tags {
		if tag == tagToAdd {
			return
		}
	}

	r.Tags = append(r.Tags, tagToAdd)
}

func (r *Ride) hasTag(searchTag string) bool {
	for _, tag := range r.Tags {
		if tag == searchTag {
			return true
		}
	}

	return false
}

func (r *Ride) removeTag(tagToRemove string) {
	for index, tag := range r.Tags {
		if tag == tagToRemove {
			r.Tags = append(r.Tags[:index], r.Tags[index+1:]...)
		}
	}
}
