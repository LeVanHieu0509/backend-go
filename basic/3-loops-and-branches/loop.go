package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10
	// Declares and initializes
	// A condition
	infinity()
	guess()

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")
}

func infinity() {
	var degrees = 0
	for {
		fmt.Println(degrees)
		degrees++

		if degrees >= 10 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

func guess() {
	//Write a guess-the-number program.
	// Make the computer pick random numbers between 1–100 until it guesses your number,
	// which you declare at the top of the program.
	// Dis - play each guess and whether it was too big or too small.

	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("Liftoff!")
	} else {
		fmt.Println("Launch failed.")
	}

}
