package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a = 1
	var b = 2

	b, a = a, b

	count := 0 // Short declaration default like var

	for count < 10 { // A new scope begins.
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	} // End scope begins. After the loop ends, the num variable goes out of scope.

	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}

	shortIf()
	shortSwitch()
}

// Cách viết if else rút gọn cấu trúc
func shortIf() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}

	// fmt.Print(num) //dùng câu lệnh này sẽ ko get được vì num có scope ở trong hàm if thôi
}

// Cách viết switch rút gọn cấu trúc
func shortSwitch() {
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}
}
