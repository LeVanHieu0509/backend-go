package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var curiosity struct {
	lat  float64
	long float64
}

type location struct {
	lat  float64
	long float64
}

type location1 struct {
	lat, long float64
}

func main() {
	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)

	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734
	fmt.Println(spirit, opportunity)

	opportunity1 := location1{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity1)
	insight := location1{lat: 4.5, long: 135.9}
	fmt.Println(insight)

	var curiosity location1
	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	spirit1 := location1{-14.5684, 175.472636}
	fmt.Println(spirit1)

	bradbury := location{-4.5895, 137.4417}
	curiosity1 := bradbury
	curiosity1.long += 0.0106
	fmt.Println(bradbury, curiosity1)
	fmt.Println("-------------------------Marshalling------------------------")

	Marshalling()
	fmt.Println("-------------------------CustomizingJSON------------------------")

	CustomizingJSON()

}

func TowSliceOfFloats() {
	type location struct {
		name string
		lat  float64
		long float64
	}

	lats := []float64{-4.5895, -14.5684, -1.9462}
	longs := []float64{137.4417, 175.472636, 354.4734}

	locations := []location{
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}
	fmt.Println(lats, longs, locations)

}

func Marshalling() {
	type location struct {
		Lat, Long float64
	}

	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))

	//result: {"Lat":-4.5895,"Long":137.4417}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CustomizingJSON() {
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	// Struct tags alter the output.

	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}
