package main

import (
	"fmt"
	"strings"
)

func main() {
	var age = 41
	var minor = age < 18
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
	fmt.Println("You find yourself in a dimly lit cavern.")
	fmt.Println("You leave the cave:", exit)

}
