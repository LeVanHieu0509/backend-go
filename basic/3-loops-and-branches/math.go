package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var walkOutside = true
	var takeTheBluePill = false

	var num = rand.Intn(10) + 1
	var distance = rand.Intn(345000001) + 56000000
	fmt.Println(num)
	num = rand.Intn(10) + 1

	fmt.Println(distance)
	fmt.Println(num)
	fmt.Println(walkOutside)
	fmt.Println(takeTheBluePill)
}
