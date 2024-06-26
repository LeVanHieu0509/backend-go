package main

import (
	"fmt"
)

func main() {
	fmt.Println("There is a cavern entrance here and a path to the east.")
	var command = "go inside"
	main2()
	switch command {
	case "go east":
		// Compares cases to command
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		// A comma- separated list of possible values
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}

}

func main2() {
	var room = "cave"

	switch room {
	case "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}
}
