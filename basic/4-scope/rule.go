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
// Câu lệnh if trong Go cho phép khai báo và khởi tạo một biến cục bộ ngay bên trong cấu trúc điều kiện if.
func shortIf() {

	// biến num được khai báo và khởi tạo bằng giá trị ngẫu nhiên từ hàm rand.Intn(3)
	// Biến num chỉ tồn tại trong khối if-else
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

	// num được khởi tạo bằng giá trị ngẫu nhiên từ rand.Intn(10)
	// num chỉ tồn tại trong khối switch và không thể sử dụng bên ngoài khối này.
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
