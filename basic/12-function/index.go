package main

import "fmt"

func FunctionExport() {
	//Trong Go, các hàm, biến và các mã định danh khác bắt đầu bằng chữ in hoa sẽ được export

}

func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func main() {
	// Giữ nguyên vùng nhớ
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	//Golang Passing Address to a Function
	update(&age, &text)

	fmt.Println("After :", text, age)

	//Anonymous Functions in Golang
	fmt.Println(area(20, 30))

	// Gọi ngay 1 function
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	//log được luôn
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	//Closures Functions in Golang
	l := 20
	b := 30

	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()

	// truy cập biến trên mỗi lần lặp của vòng lặp bên trong thân hàm.
	// một vùng nhớ mới được cấp phát cho rad trong mỗi lần lặp.
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch %p \n", i, rad, &rad)
	}

	//Higher Order Functions in Golang
	partial := partialSum(3)
	fmt.Println(partial(7))

	//Returning Functions from other Functions
	fmt.Println(squareSum(7)(6)(5))

}

func squareSum(x int) func(int) func(int) int {
	println("Func 1 x=", x)
	return func(y int) func(int) int {
		println("Func 2 y=", y)
		return func(z int) int {
			println("Func 3 z=", z)
			return x*x + y*y + z*z
		}
	}
}
