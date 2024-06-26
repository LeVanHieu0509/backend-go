package main

import (
	"fmt"
	"math/rand"
)

// Khai báo toàn cục

var era = "AD"
var d int

// Mỗi biến đã được sử dụng trong quá trình biên dịch hay chưa, nếu như không sử dụng thì biến sẽ báo waring hoặc là sẽ bị mất tích
// Có thể khai báo biến toàn cục và biến cục bộ trùng trên.
// Độ ưu tiên giữa biến local và global thì nó sẽ ưu tiên thằng local hơn.
func main() {

	randomDate()
	year := 2018
	// the package.
	// era and year are in scope.
	switch month := rand.Intn(12) + 1; month {
	// era, year, and month are in scope.
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
		// era, year, month, and day are in scope.
		// It’s a new day.
		// year is no longer in scope.
	}

	d = 10 - 7
	fmt.Println("d", d)

}

func randomDate() {
	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}
