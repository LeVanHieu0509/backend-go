package main

import "fmt"

func main() {
	//1.Tham Số Không Xác Định (Variadic Parameters)
	PrintNumbers(1, 2, 3, 4, 5)

	//2. Gửi Danh Sách Các Tham Số
	numbers := []int{1, 2, 3, 4, 5}
	PrintNumbers(numbers...) // Chuyển slice như là các tham số variadic

	//3. Tính Năng ... Trong Slice và Array
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(nums...)) // Chuyển nums slice vào hàm variadic

	//4. Khai Báo ... Trong Các Lệnh Chuyển Tiếp (Forwarding)
	p := MyPrinter{}
	numbers2 := []interface{}{1, 2, 3, 4, 5}
	p.Print(numbers2...) // Chuyển nums slice vào hàm variadic
}

/*
1. Tham Số Không Xác Định (Variadic Parameters)
Khi khai báo hàm, dấu ba chấm dùng để chỉ rằng hàm có thể nhận một số lượng tham số không xác định. Đây là tính năng variadic parameters của Go.
numbers ...int cho phép hàm PrintNumbers nhận một hoặc nhiều tham số kiểu int. Bạn có thể gọi hàm này với một danh sách các số:
*/

func PrintNumbers(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

/*
2. Gửi Danh Sách Các Tham Số
Dấu ba chấm có thể được dùng để chuyển một slice hoặc array như một danh sách các tham số cho một hàm variadic.
*/

/*
3. Tính Năng ... Trong Slice và Array
Khi bạn muốn khai báo một slice với kích thước không xác định hoặc khi muốn truyền một slice hoặc array vào hàm variadic.
*/

func Sum(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

/*
4. Khai Báo ... Trong Các Lệnh Chuyển Tiếp (Forwarding)
Khi khai báo một phương thức trong interface và triển khai nó trong struct, bạn có thể sử dụng dấu ba chấm để chuyển tham số đến một phương thức khác.
*/

type Printer interface {
	Print(args ...interface{})
}

type MyPrinter struct{}

func (p *MyPrinter) Print(args ...interface{}) {

	for _, arg := range args {
		fmt.Println("arg: ", arg)
	}
}

/*
Tóm Tắt
Dấu ba chấm ... là một tính năng mạnh mẽ trong Go, cho phép bạn làm việc với số lượng tham số không xác định
hoặc xử lý các slice và array một cách linh hoạ
*/
