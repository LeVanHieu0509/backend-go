package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vo C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	// Nếu tìm không được key thì trả về default 0
	moon := temperature["Moon"]
	fmt.Println(moon)

	// Câu lệnh if trong Go có thể kết hợp với một câu lệnh khai báo và gán giá trị để kiểm tra điều kiện ngay lập tức
	// Đoạn mã bạn đưa ra sử dụng tính năng này để kiểm tra xem một giá trị có tồn tại trong bản đồ (map) hay không

	//ok sẽ là một biến boolean nhận giá trị true nếu khóa "Moon" tồn tại
	if moon, oke1 := temperature["Moon"]; oke1 {
		fmt.Printf("On average the moon is %vo C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
	//           Prints Where is the moon?

	// Map lưu trữ số lượng sách bán được tại các cửa hàng sách
	// Kiểm tra và in số lượng sách bán được tại "Bookstore A"

	bookStore := map[string]int{
		"Book store A": 20,
		"Book store B": 21,
		"Book store C": 22,
	}

	if book, checkA := bookStore["Book store A"]; checkA {
		fmt.Println("checkA", checkA, "value:", book)
	} else {
		fmt.Println("book", book)
	}

	fmt.Println("--------Frequency of temperatures---------")
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	for _, t := range temperatures {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}

	fmt.Println("--------Grouping data with maps and slices---------")

	temperaturesDeg := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	groups := make(map[float64][]float64)
	for _, t := range temperaturesDeg {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	// Rounds temperaturesDeg down to -20, -30, and so on
	for g, temperaturesDeg := range groups {
		fmt.Printf("%v: %v\n", g, temperaturesDeg)
	}

	//Repurposing maps as sets
	fmt.Println("--------Repurposing maps as sets---------")
	var temperatures2 = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	set := make(map[float64]bool)
	for _, t := range temperatures2 {
		set[t] = true
	}
	if set[-28.0] {
		fmt.Println("set member")
	}
	fmt.Println(set)

	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)
	fmt.Println(unique)
}
