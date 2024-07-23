package main

import "fmt"

type Universe [][]bool

func main() {
	const (
		width  = 80
		height = 15
	)

	fmt.Println(sum(1, 2, 3)) // Output: 6

	numbers := []int{1, 2, 3}
	fmt.Println(sum(numbers...)) // Output: 6

}

func (u Universe) Show()

func (u Universe) Seed()

func (u Universe) Alive(x, y int) bool {
	return true
}

func (u Universe) Neighbors(x, y int) int {
	return 1
}

func (u Universe) Next(x, y int) bool {
	return true
}

func Step(a, b Universe) {
	a, b = b, a
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
