package main

import "fmt"

func main() {
	const lightYearInKm = 9.461e12
	distanceInKm := 236e15
	distanceInLightYears := distanceInKm / lightYearInKm
	fmt.Println("Distance to Canis Major Dwarf in light years:", distanceInLightYears)
}
