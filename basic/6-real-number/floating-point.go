package main

import (
	"fmt"
	"math"
)

//convert days is a float64

func main() {
	// days := 365.2425
	// var days = 365.2425
	// var days float64 = 365.2425

	// var answer float64 = 42
	var price float64
	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64)  //3.141592653589793 //64 bit
	fmt.Println(pi32)  //3.1415927 //32 bit
	fmt.Println(price) //default value = 0

	//Displaying floating-point types
	third := 1.0 / 3
	fmt.Println(third)            //0.3333333333333333
	fmt.Printf("%v\n", third)     //0.3333333333333333
	fmt.Printf("%f\n", third)     //0.333333
	fmt.Printf("%.3f\n", third)   //0.333
	fmt.Printf("%4.2f\n", third)  //0.33
	fmt.Printf("%05.2f\n", third) //00.33

	fmt.Printf("%f\n", third)
	fmt.Printf("%7.4f\n", third)
	fmt.Printf("%06.2f\n", third)

	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank) //0.30000000000000004

	//phép chia
	celsius := 21.0
	fmt.Println((celsius/5.0*9.0)+32, "o F\n")
	fmt.Println((9.0/5.0*celsius)+32, "o F\n")

	//phép nhân
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Println(fahrenheit, "o F")

	//compare

	fmt.Println(piggyBank == 0.3)
	// gía trị sẽ giống như js có 1 xíu lệch số, do là hệ hex
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)

	piggyBankk := 0.0
	for i := 0; i < 11; i++ {
		piggyBankk += 0.1
	}
	fmt.Println(piggyBankk)
}
