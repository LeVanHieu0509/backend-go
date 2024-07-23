package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	age := 41
	marsAge := float64(age) // Muốn tính toán thì cần phải cùng kiểu dữ liệu thì mới đươc.
	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am", marsAge, "years old on Mars.")
	// Prints I am 21.797587336244543 years old on Mars.

	// Có thể convert kiểu float qua kiểu int
	fmt.Println(int(earthDays))

	var bh float64 = 32767
	var h = int16(bh)
	fmt.Println(h, math.MaxUint8) //32767
	// To-do: add rocket science

	// Giá trị số nguyên này có nghĩa là đại diện cho một Điểm mã Unicode (Unicode là siêu tập hợp các ký tự ASCII )
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Println(string(pi), string(alpha), string(omega), string(bang))

	//convert number to string
	countdown := 10
	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Println(str)

	//Convert boolean -> string
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)
	var yesNo string

	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)

	//convert string to boolean
	yesNo1 := "no"
	launch1 := (yesNo1 == "yes")
	fmt.Println("Ready for launch1:", launch1)

}
