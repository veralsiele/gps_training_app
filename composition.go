package main

import (
	"fmt"
	"math"
)

type gps struct {
	world
	current     location
	destination location
}

type rover struct {
	gps
}

type world struct {
	radius float64
}

type location struct {
	name      string
	lat, long float64
}

func (w world) distance(l1, l2 location) float64 {
	s1, c1 := math.Sincos(rad(l1.lat))
	s2, c2 := math.Sincos(rad(l2.lat))
	clong := math.Cos(rad(l1.long - l2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("Km left:%v", g.world.distance(g.current, g.destination))
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (l location) description() string {
	return fmt.Sprintf("%v's longitude is: %1.f and latitude is: %1.f\n", l.name, l.long, l.lat)
}

func main() {
	Mars := world{radius: 3389.5}

	Bradbury := location{
		name: "BradburyLanding",
		lat:  -4.5895,
		long: 137.4417,
	}

	Elysium := location{
		name: "Elysium Planitia",
		lat:  4.5,
		long: 135.9,
	}

	gps := gps{
		world:       Mars,
		current:     Bradbury,
		destination: Elysium,
	}

	curiosity := rover{gps}

	fmt.Print(curiosity.message())

}
