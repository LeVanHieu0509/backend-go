// My weight loss program.
package main

import (
	"fmt"
	"math"
	"time"
)

/*
	bool:                    %t
	int, int8 etc.:          %d
	uint, uint8 etc.:        %d, %#x if printed with %#v
	float32, complex64, etc: %g
	string:                  %s
	chan:                    %p
	pointer:                 %p

	struct:             {field0 field1 ...}
	array, slice:       [elem0 elem1 ...]
	maps:               map[key1:value1 key2:value2 ...]
	pointer to above:   &{}, &[], &map[]

	Advance:
	1. Chuỗi kí tự
		%s - String: Định dạng chuỗi (string)
		%q - Quoted: Định dạng chuỗi với dấu ngoặc kép (double-quoted string)
		%x - Hexadecimal: Định dạng chuỗi dưới dạng các byte hex (hex bytes)
		%X - Uppercase Hexadecimal: Định dạng chuỗi dưới dạng các byte hex (hex bytes) với chữ cái viết hoa
	2. Số Nguyên
		%d - Decimal: Định dạng số nguyên hệ thập phân (decimal integer)
		%b - Binary: Định dạng số nguyên dưới dạng nhị phân (binary integer)
		%c - Character: Định dạng số nguyên dưới dạng ký tự Unicode (Unicode character)
		%o - Octal: Định dạng số nguyên dưới dạng bát phân (octal integer)
		%q - Quoted character: Định dạng số nguyên dưới dạng ký tự Unicode (single-quoted character)
		%x - Hexadecimal: Định dạng số nguyên dưới dạng thập lục phân (hexadecimal integer)
		%X - Uppercase Hexadecimal: Định dạng số nguyên dưới dạng thập lục phân (hexadecimal integer) với chữ cái viết hoa
		%U - Unicode: Định dạng số nguyên dưới dạng mã Unicode (Unicode code point)
	3. Số Thực
		%f - Floating-point: Định dạng số thực (floating-point) với dấu chấm động (default format)
		%e - Exponential: Định dạng số thực dưới dạng ký hiệu khoa học (scientific notation, e.g., -1.234456e+78)
		%E - Uppercase Exponential: Định dạng số thực dưới dạng ký hiệu khoa học (scientific notation) với chữ E viết hoa (e.g., -1.234456E+78)
		%g - General: Định dạng số thực dưới dạng gọn gàng hơn giữa %e và %f
		%G - Uppercase General: Định dạng số thực dưới dạng gọn gàng hơn giữa %E và %f
	4. Boolean
		%t - Boolean: Định dạng giá trị boolean (true hoặc false)

	5. Con Trỏ và Địa Chỉ
		%p - Pointer: Định dạng con trỏ (pointer)
	6. Các Cấu Trúc Dữ Liệu
		%v - Value: Định dạng giá trị theo cách tiêu chuẩn nhất (default format)
		%+v -  Detailed Value: Định dạng giá trị bao gồm cả tên trường (field names) cho các cấu trúc (structs)
		%#v - Go-syntax: Định dạng giá trị theo cách có thể tái sử dụng trong mã Go (Go-syntax format)
		%T - Type: Định dạng kiểu của giá trị (type of the value)
	7. Lỗi
		%w - Wrapped Error: Định dạng lỗi để bọc lỗi trong lỗi khác (error wrapping, sử dụng với fmt.Errorf)
*/
// A comment for human readers
// main is the function where it all begins.
func main() {

	fmt.Println("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	// Printlns 56.3667 Printlns 21
	// NOTE Though listing 2.1 displays weight in pounds, the chosen unit of measurement doesn’t impact the weight calculation. Whichever unit you choose, the weight on Mars is 37.83% of the weight on Earth.
	fmt.Println(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Println(" years old.")

	// example
	// %v giá trị ở định dạng mặc định
	// khi in cấu trúc, cờ dấu cộng (%+v) sẽ thêm tên trường

	var i interface{} = 23
	fmt.Printf("%v\n", i) //the value in a default format

	integer := 23
	// Each of these prints "23" (without the quotes).
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	// The special verb %T shows the type of an item rather than its value.
	fmt.Printf("%T %T\n", integer, &integer)
	// Result: int *int

	// Booleans print as "true" or "false" with %v or %t.
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// Result: true true

	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// Result: 42 42 2a 52 101010

	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// Result: 3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00

	// Complex numbers format as parenthesized pairs of floats, with an 'i'
	// after the imaginary part.
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// Result: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)

	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// Result: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'

	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// Result: foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`

	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// Result: map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}

	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// Result: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))
	// Result: &{Kim 22} 0x0
	// fmt.Printf("%v %p\n", pointer, pointer)
	// Result: &{Kim 22} 0x010203 // See comment above.

	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// Result: [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// Result: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98

	now := time.Unix(123456789, 0).UTC() // time.Time implements fmt.Stringer.
	fmt.Printf("%v %q\n", now, now)
	// Result: 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"

}
