package main

type Route struct {
	points []*Point
}

func newRoute(p *Point) *Route {
	route := &Route{
		points: []*Point{p},
	}

	return route
}

func (r *Route) getPoint(index int) *Point {
	return r.points[index]
}

func (r *Route) addPoint(p ...*Point) {
	m := len(r.points)
	n := m + len(p)
	if n > cap(r.points) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]*Point, (n+1)*2)
		copy(newSlice, r.points)
		r.points = newSlice
	}
	r.points = r.points[0:n]
	copy(r.points[m:n], p)
}

func (r *Route) removePoint(index int) {
	r.points = append(r.points[:index], r.points[index+1:]...)
}

func (r *Route) getPoints() []*Point {
	return r.points
}

func (r *Route) analize() {
	// analize track
}
