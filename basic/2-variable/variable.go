// My weight loss program.
package main

import "fmt"

func main() {
	var a = 1
	var b = 2

	b, a = a, b //gán 2 biến cho nhau

	fmt.Println(a, b)

	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km

	var (
		distance1 = 56000000
		speed1    = 100800
	)

	const hoursPerDay = 24
	var speed = 100800       // km/h
	var distance2 = 96300000 // km

	var weight = 149.0
	weight = weight * 0.3783
	weight *= 0.3783

	var age = 41
	age = age + 1
	age += 1
	age++
	age--

	fmt.Println("age", age)
	fmt.Println(weight, "days")
	fmt.Println(distance2/speed/hoursPerDay, "days")

	fmt.Println(distance1/speed1, "seconds")

	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	//--------------------
	//Biến này sẽ không được sử dụng và không chiếm dung lượng bộ nhớ, là một biến vô tri

	m1, _ := GetMoney()

	fmt.Println("m1::", m1)

}

func GetMoney() (int, int) {
	return 123, 234
}
