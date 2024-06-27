package basic

import "fmt"

func AddOne(num int) int {

	return num + 1
}

func AddTwo(num int) int {
	if 1 == 2 {
		fmt.Println("Failed 1 == 2 ")
	}
	if num == 4 {
		fmt.Println("Failed 4")
	}
	if num == 6 {
		fmt.Println("Failed 4")
	}
	return num + 1
}
