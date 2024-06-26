package main

import (
	"fmt"
	"math/rand"
	"time"
)

func convertToDollars(cents int) string {
	//convert to dollar
	dollars := float64(cents) / 100

	return fmt.Sprintf("$%.2f", dollars)
}
func main() {
	rand.Seed(time.Now().UnixNano())

	balance := 0
	for balance < 2000 {
		coin := []int{5, 10, 15}[rand.Intn(3)] //chọn ngẫu nhiên vị trí 1,2,3 trong array [5,10,15]

		balance += coin

		fmt.Println("Running balance:", convertToDollars(balance))
	}

	// Display the final balance in dollars
	fmt.Println("Final balance:", convertToDollars(balance))
}
