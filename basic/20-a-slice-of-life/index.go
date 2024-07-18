package main

type Universe [][]bool

func main() {
	const (
		width  = 80
		height = 15
	)

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
