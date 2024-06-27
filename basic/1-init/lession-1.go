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

	1. %v giá trị ở định dạng mặc định
	2. %T in ra kiểu dữ liệu là gì
	3. %t in ra kiểu boolean
	4. %f in ra dấu thập phân nhưng không có số mũ, ví dụ: 123,456 = 3f
	5. %c in ra ký tự tương ứng với giá trị byte.
	6. %w giữ nguyên chuỗi lỗi để gỡ lỗi
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
