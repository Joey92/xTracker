package main

type Point struct {
	Lat, Lng  float64
	Source    string
	Timestamp int64
}

func newPoint(lat float64, lng float64, time int64, source string) *Point {
	point := &Point{
		Lat:       lat,
		Lng:       lng,
		Source:    source,
		Timestamp: time,
	}

	return point
}

func (p *Point) setSource(source string) *Point {
	p.Source = source

	return p
}

func (p *Point) getSource() string {
	return p.Source
}

func (p *Point) setLat(latitude float64) *Point {
	p.Lat = latitude

	return p
}

func (p *Point) getlat() float64 {
	return p.Lat
}

func (p *Point) setLng(longitude float64) *Point {
	p.Lng = longitude

	return p
}

func (p *Point) getLng() float64 {
	return p.Lng
}
