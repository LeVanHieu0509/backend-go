package main

import "fmt"

func main() {
	// add a number larger than one var red uint8 = 255
	var red uint8 = 255
	red += 2
	fmt.Println("red", red)
	// Figure 7.1 Carrying the 1 in binary addition
	// Integers wrap around 59
	// Quick check 7.5 Use the Go Playground to experiment with the wrapping behavior of integers:
	// 1 In listing 7.2, the code increments red and number by 1. What happens when you add a larger number to either variable?
	// 2 Go the other way. What happens if you decrement red when it’s 0 or number when it’s equal to –128?
	// 3 Wrapping applies to 16-bit, 32-bit, and 64-bit integers too. What happens if you declare a uint16 assigned to the maximum value of 65535 and then increment it by 1?
	// TIP The math package defines math.MaxUint16 as 65535 and similar min/max constants for each architecture-independent integer type. Remember that int and uint could be either 32-bit or 64-bit, depending on the underlying hardware.
	var number1 int8 = 127
	number1 += 3
	fmt.Println("number1", number1)
	// 2 wrap the other way
	red = 0
	red--
	fmt.Println("red", red)
	number1 = -128
	number1--
	fmt.Println("number1", number1)
	// Prints 1
	// Prints -126
	// Prints 255
	// Prints 127
	//3  wrapping with a 16-bit unsigned integer
	var green uint16 = 65535
	green++
	fmt.Println("green", green) //Prints 0
}
