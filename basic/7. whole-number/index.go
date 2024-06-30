package main

import "fmt"

// 1, int8: -128 to 127 =>
// 2, uint8 0 to 255 => default: 1
// 3, int16 -32,768 to 32,767
// 4, uint16 0 to 65535
// 5, int32 -2,147,483,648 to 2,147,483,647
// 6, uint32 0 to 4,294,967,295
// 7, int64 -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
// 8, uint 0 to 18,446,744,073,709,551,615

// nếu bạn đang xử lý các số lớn hơn hai tỷ sử dụng int64 hoặc uint64 thay vì int và uint.

func main() {
	year := 2018 //Prints Type int
	fmt.Printf("Type %T for %v\n", year, year)

	days := 365.2425                        // Prints Type float64
	fmt.Printf("Type %T for %[1]v\n", days) // Thay vì lặp lại biến hai lần, sài [1] cho động từ định dạng thứ hai.

	// Nếu có nhiều màu cần lưu trữ tuần tự, chẳng hạn như trong ảnh không nén,
	// Bạn có thể tiết kiệm bộ nhớ đáng kể bằng cách sử dụng số nguyên 8 bit.

	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a)
	b := 42
	fmt.Printf("Type %T for %[1]v\n", b)
	c := 3.14
	fmt.Printf("Type %T for %[1]v\n", c)
	d := true
	fmt.Printf("Type %T for %[1]v\n", d)

	var red, green, blue uint8 = 0, 141, 213
	// var red, green, blue uint8 = 0x00, 0x8d, 0xd5

	// Để hiển thị số ở dạng thập lục phân, bạn có thể sử dụng động từ định dạng %x hoặc %X với Printf
	fmt.Printf("%x, %x, %x", red, green, blue)

	// Để tạo ra màu sắc quen thuộc trong tệp .css, các giá trị thập lục phân cần có một số phần đệm.
	// Giống như các động từ định dạng %v và %f, bạn có thể chỉ định số chữ số tối thiểu (2) và phần đệm bằng 0 với %02x:
	// %02x - là một chuỗi định dạng để hiển thị các giá trị số nguyên dưới dạng chuỗi hex (hệ thập lục phân)
	fmt.Printf("\n color: #%02x%02x%02x;", red, green, blue)

	// Số nguyên không dấu 8 bit (uint8) có phạm vi từ 0–255.
	// Việc tăng vượt quá 255 sẽ quay về 0.
	// Danh sách sau đây sẽ tăng cả số nguyên 8 bit có dấu và không dấu, khiến chúng bao quanh.

	var red1 uint8 = 255 // (default nếu vượt quá: 0)
	red1++
	fmt.Println("red1", red1) //Việc tăng vượt quá 255 sẽ quay về 0.
	var number int8 = 127     // (default nếu vượt quá: -128)
	number++
	fmt.Println("number", number) //Việc tăng vượt quá 255 sẽ quay về -128.

	// Prints 00000011 green++ Prints 00000100
	// %08b là một chuỗi định dạng để hiển thị một số nguyên dưới dạng chuỗi số nhị phân với độ dài 8 chữ số

	var green1 uint8 = 3
	fmt.Printf("%08b\n", green1) //0 * 2^7 + 0 * 2^6 + 0 * 2^5 + 0 * 2^4 + 0 * 2^3 + 0 * 2^2 + 1 * 2^1 + 1 * 2^0 = = 0 + 0 + 0 + 0 + 0 + 0 + 2 + 1 = 3
	green1++
	fmt.Printf("%08b\n", green1)

}
