package main

import "fmt"

func main() {
	var command = "go east"
	// If command is equal to “go east”
	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of yourOtherwise, if command is equal to “go inside”➥life.")
	} else {
		// Or, if anything else
		fmt.Println("Didn't quite get that.")
	}

	fmt.Println("apple" > "banana") // false

	logical()
	LeapYear()
}

func logical() {
	var room = "cave"
	if room == "cave" {
		fmt.Println("You find yourself in a dimly lit cavern.")
	} else if room == "entrance" {
		fmt.Println("There is a cavern entrance here and a path to the east.")
	} else if room == "mountain" {
		fmt.Println("There is a cliff here. A path leads west down the mountain.")
	} else {
		fmt.Println("Everything is white.")
	}
}

func LeapYear() {
	fmt.Println("The year is 2100, should you leap?")
	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}

	var haveTorch = true
	var litTorch = false
	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}
}
