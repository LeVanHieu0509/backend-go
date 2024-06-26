package main

import (
	"fmt"
	"io"
	"strings"
)

// Quét, Fscan, Sscan coi các "dòng mới trong đầu vào là khoảng trắng".
// Scanln, Fscanln và Sscanln "dừng quét ở dòng mới" và yêu cầu các mục phải được theo sau bởi dòng mới hoặc EOF.

func main() {
	// var (
	// 	i int
	// 	b bool
	// 	s string
	// )
	// r := strings.NewReader("5 true gophers")
	// n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	// }
	// fmt.Println(i, b, s)
	// fmt.Println(n)
	testF()

}

func testF() {
	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	var i int = 1

	for {
		n, err := fmt.Fscanln(r, &a, &b, &c) // phân tích từng dòng của chuỗi thành các biến với kiểu dữ liệu tương ứng.
		fmt.Println("ok", i)

		if err == io.EOF { // đọc từng dòng cho đến khi đạt đến cuối chuỗi
			break
		}
		if err != nil {
			panic(err)
		}
		i++
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}
}
