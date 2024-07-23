package main

import (
	"fmt"
	"math"
)

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decimal(), long.decimal())

	curiosity := location{lat.decimal(), long.decimal()}
	curiosity = newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity)

	var mars = world{radius: 3389.5}
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}
	// Uses the distance method on mars
	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist)

	// practice
	fmt.Println("-----------------------------------------")
	locations := []location1{
		{lat: 14.3462, long: 23.1232, rover: "Spitit", landing: "Columbia memorial Station"},
		{lat: 14.3462, long: 23.1232, rover: "Opportunity", landing: "Challenger Memorial Station"},
		{lat: 14.3462, long: 23.1232, rover: "Curiosity", landing: "Bradbury Landing"}}

	for _, el := range locations {
		el.printLocation()
	}

}

// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

/*
Việc kết hợp các phương pháp và cấu trúc cung cấp phần lớn những gì ngôn ngữ
cổ điển cung cấp mà không cần đưa ra một tính năng ngôn ngữ mới.
Hàm xây dựng là các hàm thông thường.
*/

/*
viết chương trình khai báo vị trí cho từng vị trí trong bảng 22.1. In ra từng vị trí theo độ thập phân.
*/

type location1 struct {
	lat, long      float32
	rover, landing string
}

func (lc *location1) printLocation() {
	fmt.Printf("%v - %v - %v - %v\n", lc.rover, lc.landing, lc.lat, lc.long)

}
