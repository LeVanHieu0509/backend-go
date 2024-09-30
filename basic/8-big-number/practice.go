package main

import "fmt"

func main() {
	// 9.461e12 là cách viết ngắn gọn trong lập trình
	const lightYearInKm = 9.461e12 // 9.461×10 12km.
	distanceInKm := 236e15
	distanceInLightYears := distanceInKm / lightYearInKm
	fmt.Println("Distance to Canis Major Dwarf in light years:", distanceInLightYears)
}
