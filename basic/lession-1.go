// My weight loss program.
package main

import "fmt"

// A comment for human readers
// main is the function where it all begins.
func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	// Prints 56.3667 Prints 21
	// NOTE Though listing 2.1 displays weight in pounds, the chosen unit of measurement doesn’t impact the weight calculation. Whichever unit you choose, the weight on Mars is 37.83% of the weight on Earth.
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.")
}
