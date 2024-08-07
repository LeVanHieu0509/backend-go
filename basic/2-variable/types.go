package main

import (
	"fmt"
	"io"
	"math"
	"math/bits"
	"math/cmplx"
	"os"
	"reflect"
	"unsafe"
)

/*
int	Platform Dependent - Số nguyên - Default int - vòng lặp nên sử dụng
int8	8 bits/1 byte - Scope:  -2^7  đến 2^7 - 1.
int16	16 bits/2 byte - Scope:  -2^15  đến 2^15 - 1.
int32	32 bits/4 byte - Scope:  -2^31  đến 2^31 - 1.
int64	64 bits/8 byte - Scope:  -2^63  đến 2^63 - 1, time.Duration

uint	Phụ thuộc nền tảng: máy 32 bit, phạm vi của int sẽ là -2^31  đến 2^31 - 1, máy 64 bit, phạm vi của int sẽ là -2^63  đến 2^63 - 1.
uint8	8 bit/1 byte - scope: 0 đến 255 hoặc 0 đến 2^8  - 1.
uint16	16 bit/2 byte - scope: 0 đến 2^16  - 1
uint32	32 bit/4 byte - scope: 0 đến 2^32  - 1
uint64	64 bit/8 byte - scope: 0 đến 2^64  - 1

float32	32 bit hoặc 4 byte - scope: 1.18E-38 (1.18 × 10^-38) -> 3.4E+38 (3.4 × 10^38) - Value default:  0.0
float64	64 bit hoặc 8 byte - Default float64 - scope: 2,23E-38 đến 1.8E+308  - Value default:  0.0

complex64	Cả phần thực và phần ảo đều là float32
complex128	Cả phần thực và phần ảo đều là float64
*/

func main() {
	var a int = 10
	var b int32 = 20

	fmt.Println(a, b)

	//Khai báo biến song song
	a, b = 1, 2

	var c float32 = 10.34

	// Lấy phần nguyên
	fmt.Println(int(c)) //10.34

	fmt.Println("a,b", a, b)
	fmt.Print(5 % 3) // 2 lấy phần nguyên

	var i int32
	var j int64

	i, j = 2, 3

	if i == 2 || j == 3 {
		fmt.Println("equal")
	}

	//boolean
	ok := true
	fmt.Println(ok)

	m := 1

	if m == 1 {
		fmt.Println("is true")
	}

	///////string
	s1 := "hello"
	s2 := "tip js go"

	s := `row1\\r\nrow2`

	fmt.Println(s1, s2, s)

	//concat
	s3 := s1 + s2
	fmt.Println(s3)

	//length
	fmt.Println(len(s3))
	fmt.Println(s3[2:4]) // Cắt chuỗi từ 2 cho tới 4

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//int8
	//Declare a int 8
	var abc int8 = 2

	//Size of int8 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(abc))
	fmt.Printf("abc's type is %s\n", reflect.TypeOf(abc))

	//uint => máy của tôi đang chạy 64 bit
	sizeOfUintInBits := bits.UintSize
	fmt.Println(sizeOfUintInBits)
	fmt.Printf("a's type is %s\n", reflect.TypeOf(sizeOfUintInBits))
	fmt.Println("---------------------uintptrFc---------------------")
	uintptrFc()
	fmt.Println("---------------------float---------------------")
	float()
	fmt.Println("---------------------complex---------------------")
	complexInit()
	fmt.Println("---------------------byte---------------------")
	byteInit()
	fmt.Println("---------------------utf8---------------------")
	utf8Init()
	fmt.Println("---------------------string---------------------")
	stringInit()
	fmt.Println("---------------------boolean---------------------")
	booleanInit()
	fmt.Println("---------------------array---------------------")
	arrayInit()
	fmt.Println("---------------------struct---------------------")
	structInit()
	fmt.Println("---------------------defer---------------------")
	deferInit()
	fmt.Println("---------------------make---------------------")
	makeInit()
	fmt.Println("-----------------------channels-------------------")
	channelsInit()
	fmt.Println("---------------------map---------------------")
	mapInit()
	fmt.Println("---------------------pointer---------------------")
	pointerInit()
	fmt.Println("--------------------function----------------------")
	functionInit()
	fmt.Println("--------------------interface----------------------")
	interfaceInit()
	interfaceExample()

}

type sample struct {
	a int
	b string
	c string
}

// uintptrFc
// Con trỏ chủ yếu được sử dụng để truy cập bộ nhớ không an toàn.
// Khi bạn muốn lưu giá trị địa chỉ con trỏ để in hoặc lưu trữ nó

func uintptrFc() {
	s := &sample{a: 1, b: "test", c: "ok"}

	// Getting the address of field b in struct s
	// unsafe.Pointer(s): Chuyển con trỏ của s thành kiểu unsafe.Pointer.
	// uintptr(unsafe.Pointer(s)): Chuyển đổi unsafe.Pointer thành uintptr, một kiểu số nguyên không dấu đủ lớn để chứa một con trỏ.
	// unsafe.Offsetof(s.b): Lấy offset của trường b trong cấu trúc s.
	// uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b): Tính toán địa chỉ của trường b bằng cách thêm offset của b vào địa chỉ của s
	// unsafe.Pointer(...): Chuyển đổi lại từ uintptr thành unsafe.Pointer

	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))

	//Typecasting it to a string pointer and printing the value of it
	// (*string)(p): Ép kiểu unsafe.Pointer thành con trỏ kiểu string.
	// *(*string)(p): Truy xuất giá trị của con trỏ kiểu string.
	// Chỉ in ra được giá trị của string đầu tiên.
	fmt.Println(*(*string)(p)) //fmt.Println(...): In giá trị ra màn hình.
}

func float() {
	var f32 float32 = math.Pi
	var f64 float64 = math.Pi

	fmt.Printf("float32: %.10f\n", f32)
	fmt.Printf("float64: %.10f\n", f64)

	// Kiểm tra phạm vi
	var minFloat32 float32 = 1.18e-38
	var maxFloat32 float32 = 3.4e38
	var minFloat64 float64 = 2.23e-308
	var maxFloat64 float64 = 1.7976931348623157e+308

	fmt.Println("float32 min:", minFloat32)
	fmt.Println("float32 max:", maxFloat32)
	fmt.Println("float64 min:", minFloat64)
	fmt.Println("float64 max:", maxFloat64)

}

func complexInit() {
	var a float32 = 3
	var b float32 = 5

	//Initialize-1
	c := complex(a, b)

	//Initialize-2
	var d complex64
	d = 4 + 5i

	//Print Size
	fmt.Printf("c's size is %d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("d's size is %d bytes\n", unsafe.Sizeof(d))

	//Print type
	fmt.Printf("c's type is %s\n", reflect.TypeOf(c))
	fmt.Printf("d's type is %s\n", reflect.TypeOf(d))

	//Operations on complex number
	fmt.Println(c+d, c-d, c*d, c/d)
}

func byteInit() {
	var r byte = 'a'

	//Print Size
	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

	//Print Type
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	//Print Character
	fmt.Printf("Character: %c\n", r)
	s := "abc"

	//This will the decimal value of byte
	// ASCII của a, b, và c lần lượt là 97, 98, và 99.
	fmt.Println([]byte(s))
}

//trong GO, mọi chuỗi đều được mã hóa bằng utf-8.

func utf8Init() {
	// sử dụng mảng rune khi tất cả các giá trị trong mảng đều có nghĩa là Điểm Mã Unicode.

	fmt.Printf("%U\n", []rune("0b£"))            //In điểm mã Unicode
	fmt.Printf("Character: %c\n", []rune("0b£")) //Ký tự in
}

func stringInit() {
	// 2 ways: "", ``

	s := "ab£"
	fmt.Println([]byte(s))
	//[48 98 194 163] 4 bytes

	x := "this\nthat"
	fmt.Printf("x is: %s\n", x)

	//String in back quotes
	y := `this\nthat`
	fmt.Printf("y is: %s\n", y)

	s = "ab£"

	//This will print the byte sequence.
	//Since character a and b occupies 1 byte each and £ character occupies 2 bytes.
	//The final output will 4 bytes
	fmt.Println([]byte(s))

	// chú ý:  khi bạn cố in độ dài của chuỗi trên bằng len(“ab£”)
	// nó sẽ xuất ra 4 chứ không phải 3 vì nó chứa 4 byte.
	fmt.Println(len(s))

	// phạm vi vòng lặp trên các chuỗi byte tạo thành mỗi ký tự
	for _, c := range s {
		fmt.Println(string(c))
	}

}

func booleanInit() {
	//Default value will be false it not initialized
	var a bool
	fmt.Printf("a's value is %t\n", a)

	//And operation on one true and other false
	andOperation := 1 < 2 && 1 > 3
	fmt.Printf("Ouput of AND operation on one true and other false %t\n", andOperation)

	//OR operation on one true and other false
	orOperation := 1 < 2 || 1 > 3
	fmt.Printf("Ouput of OR operation on one true and other false: %t\n", orOperation)

	//Negation Operation on a false value
	negationOperation := !(1 > 2)
	fmt.Printf("Ouput of NEGATION operation on false value: %t\n", negationOperation)
}

func arrayInit() {
	sample := [3]string{"a", "b", "c"}
	fmt.Println(sample)
}

func structInit() {

	//Declare a struct
	type employee struct {
		name   string
		age    int
		salary float64
	}

	//Initialize a struct without named fields
	employee1 := employee{"John", 21, 1000}
	fmt.Println(employee1)

	//Initialize a struct with named fields
	employee2 := employee{
		name:   "Sam",
		age:    22,
		salary: 1100,
	}
	fmt.Println(employee2)

	//Initializing only some fields. Other values are initialized to default zero value of that type
	employee3 := employee{name: "Tina", age: 24}
	fmt.Println(employee3)

	// In ra kích thước của struct employee3
	fmt.Printf("Kích thước của struct employee3: %d bytes\n", unsafe.Sizeof(employee3))

	// In ra địa chỉ của từng trường trong struct
	fmt.Printf("Địa chỉ của trường Name: %p\n", &employee3.name)
	fmt.Printf("Địa chỉ của trường Age: %p\n", &employee3.age)

	/*

		Struct có vùng nhớ riêng: Khi bạn khai báo và khởi tạo một struct, bộ nhớ sẽ được cấp phát để lưu trữ các trường của struct.
		Kích thước của struct: Phụ thuộc vào các trường mà struct đó chứa.
		Địa chỉ của các trường: Mỗi trường trong struct có địa chỉ bộ nhớ riêng, và bạn có thể lấy địa chỉ này bằng cách sử dụng toán tử &.
	*/

}

// đọc toàn bộ nội dung của tệp đó và trả về nội dung dưới dạng chuỗi
// Thường được sử dụng để close 1 kết nối sau khi open

func deferInit() (string, error) {
	/*
		Câu lệnh trì hoãn của Go lên lịch cho một lệnh gọi hàm (hàm trì hoãn) được chạy ngay lập tức
		trước khi hàm thực hiện lệnh trì hoãn trả về.
		Đó là một cách bất thường nhưng hiệu quả để giải quyết các tình huống như tài nguyên
		phải được giải phóng bất kể hàm đó đi theo đường nào để trả về.
		Các ví dụ điển hình là mở khóa mutex hoặc đóng tệp.
	*/

	f, err := os.Open("../../log/log.txt")

	fmt.Println("f", f)

	if err != nil {
		fmt.Println(err)
		return "error", err
	}

	// đảm bảo rằng tệp sẽ được đóng khi hàm deferInit kết thúc, dù kết thúc theo cách nào.
	// Dù hàm kết thúc theo cách nào (trả về từ giữa hàm hoặc từ cuối hàm), tệp f vẫn sẽ được đóng do lệnh defer f.Close().

	// Khi bạn sử dụng defer, bạn đang lên lịch cho một hàm (gọi là hàm deferred) sẽ được thực thi ngay trước khi hàm chứa defer trả về.
	// Các hàm deferred được xếp vào một ngăn xếp. Khi hàm bao quanh trả về, các hàm deferred được gọi theo thứ tự LIFO (Last In, First Out).
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}

	defer fmt.Println("hieu")
	return string(result), nil // f will be closed if we return here.

}

func makeInit() {
	// có kích thước động, tham chiếu đến các phần tử của mảng
	// là một kiểu tham chiếu vì nó tham chiếu nội bộ một mảng

	// format: make([]TYPE, length, capacity)
	dwarfs1 := make([]string, 0, 10)
	dwarfs1 = append(dwarfs1, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")

	s := make([]string, 1, 1)
	ok := append(s, "Crea")

	fmt.Println(s, ok)

	//Direct intialization
	p := []string{"a", "b", "c"}
	fmt.Println(p)

	//Append function
	p = append(p, "d")
	fmt.Println(p)

	//Iterate over a slice
	// range p: sẽ duyệt qua tất cả các phần tử của mảng p
	// Ở mỗi lần lặp, giá trị của phần tử hiện tại được gán cho biến val.
	for index, value := range p {
		fmt.Printf("%v: %v\n", index, value)
	}
}

func sendEvents(eventsChan chan<- string) {
	//Channel Direction: chan<- string chỉ cho phép gửi giá trị vào channel, không thể nhận giá trị từ channel.
	eventsChan <- "a"
	eventsChan <- "b"
	eventsChan <- "c"
	close(eventsChan)

}

func channelsInit() {

	// cung cấp sự đồng bộ hóa và liên lạc giữa các goroutines
	// goroutine có thể gửi các giá trị và nhận các giá trị

	// Channels: Được sử dụng để truyền dữ liệu giữa các goroutines.
	// Goroutines: Cho phép thực hiện các tác vụ đồng thời.
	// Close Channel: Đóng channel để báo hiệu rằng không còn dữ liệu nào nữa sẽ được gửi.
	/*

		Kênh không có bộ đệm
			- Nó không có khả năng lưu giữ và giá trị và do đó
			- Gửi trên một kênh sẽ bị chặn trừ khi có một goroutine khác để nhận.
			- Việc nhận bị chặn cho đến khi có một con goroutine khác ở phía bên kia gửi đi.
		Kênh đệm
			- Bạn có thể chỉ định kích thước của bộ đệm ở đây và cho chúng
			- Gửi trên kênh đệm chỉ bị chặn nếu bộ đệm đầy
			- Nhận là khối duy nhất là bộ đệm của kênh trống
	*/

	//Kênh đệm
	eventsChan := make(chan string, 3)
	eventsChan <- "a"
	eventsChan <- "b"
	eventsChan <- "c"

	//Closing the channel
	close(eventsChan)
	for event := range eventsChan {
		fmt.Println("Kênh đệm:", event)
	}

	// Kênh không đệm
	eventsChanKhongDem := make(chan string)
	go sendEvents(eventsChanKhongDem)

	for event := range eventsChanKhongDem {
		fmt.Println("Kênh không có bộ đệm:", event)
	}

}

func mapInit() {
	//là loại dữ liệu được tham chiếu
	//Specify values
	// employeeSalary := map[string]int{
	// 	"John": 1000,
	// 	"Sam":  2000,
	// }

	// employeeSalary["John"] = 1000
	// delete(employeeSalary, "John")

	//Declare
	var employeeSalary map[string]int
	fmt.Println(employeeSalary)

	//Intialize using make
	employeeSalary2 := make(map[string]int)
	fmt.Println(employeeSalary2)

	//Intialize using map lieteral
	employeeSalary3 := map[string]int{
		"John": 1000,
		"Sam":  1200,
	}
	fmt.Println(employeeSalary3)

	//Operations
	//Add
	employeeSalary3["Carl"] = 1500

	//Get
	fmt.Printf("John salary is %d\n", employeeSalary3["John"])

	//Delete
	delete(employeeSalary3, "Carl")

	//Print map
	fmt.Println("\nPrinting employeeSalary3 map")
	fmt.Println(employeeSalary3)

}

func pointerInit() {

	//Declare: gán giá trị và truy cập giá trị thông qua con trỏ
	// b là con trỏ kiểu int.
	var b *int

	//biến a sẽ được cấp phát một vị trí cụ thể trong bộ nhớ để lưu trữ giá trị của nó
	a := 2

	//&: Lấy địa chỉ của biến a và gán cho con trỏ b (& là toán tử lấy địa chỉ, nó trả về địa chỉ bộ nhớ của biến a)
	b = &a

	//Will print a address. Output will be different everytime.
	// In địa chỉ của a (địa chỉ mà con trỏ b trỏ tới):
	fmt.Println(b)

	// Toán tử * có thể được sử dụng để hủy đăng ký một con trỏ, nghĩa là lấy giá trị tại địa chỉ được lưu trong con trỏ.
	// Được gọi là dereferencing, nó trả về giá trị của biến tại địa chỉ mà b trỏ tới, tức là giá trị của a.
	fmt.Println(*b)

	// Khởi tạo con trỏ b với một địa chỉ bộ nhớ mới
	b = new(int)

	// gán giá trị 10 cho vùng bộ nhớ mà con trỏ b đang trỏ tới.
	*b = 10

	// in ra giá trị tại địa chỉ mà con trỏ b trỏ tới, tức là giá trị 10
	fmt.Println(b)

}

// func
func doOperation(fn func(int, int) int, x, y int) int {
	return fn(x, y)
}

func functionInit() {
	add := func(a, b int) int {
		return a + b
	}
	multiply := func(a, b int) int {
		return a * b
	}

	x, y := 3, 4

	// Sử dụng hàm doOperation với hàm add
	resultAdd := doOperation(add, x, y)
	fmt.Printf("Addition of %d and %d is %d\n", x, y, resultAdd)

	// Sử dụng hàm doOperation với hàm multiply
	resultMultiply := doOperation(multiply, x, y)
	fmt.Printf("Multiplication of %d and %d is %d\n", x, y, resultMultiply)
}

// interface
// Định nghĩa một interface có tên shape.
type shape interface {

	// Interface này yêu cầu bất kỳ kiểu nào muốn thực hiện interface này
	// phải có phương thức area trả về một giá trị kiểu int.
	area() int
}

// Định nghĩa một kiểu cấu trúc (struct) có tên square.
type square struct {
	//Kiểu square có một trường dữ liệu duy nhất là side, kiểu int, đại diện cho cạnh của hình vuông.
	side int
}

// Định nghĩa một phương thức có tên area cho kiểu square.
// (s *square): Đây là phương thức của kiểu square, với s là một con trỏ tới square.

func (s *square) area() int {
	return s.side * s.side //Phương thức area trả về diện tích của hình vuông
}

func interfaceInit() {
	/*
	   Sử dụng con trỏ thường hiệu quả hơn về mặt bộ nhớ, đặc biệt khi làm việc với các struct lớn.
	   Bằng cách sử dụng con trỏ, bạn tránh được việc sao chép toàn bộ struct mỗi khi truyền nó xung quanh.
	*/

	// Khai báo một biến s kiểu shape. Lúc này, s chưa được gán giá trị nào.
	// khi gán một giá trị cho một biến interface, giá trị đó phải thỏa mãn tất cả các phương thức được định nghĩa trong interface
	var s shape

	// 1, Khởi tạo một giá trị square với side bằng 4 và lấy địa chỉ của nó (tạo một con trỏ tới square).
	// 2, Gán giá trị con trỏ square này cho biến s. Bởi vì square thực hiện interface shape
	// (vì nó có phương thức area), việc gán này hợp lệ.

	// chỉ có con trỏ đến một square mới có thể gọi phương thức area.
	s = &square{side: 4}

	// Gọi phương thức area của đối tượng square thông qua interface shape.
	fmt.Println(s.area())

	/*
		1, Interfaces trong Go: Định nghĩa một tập hợp các phương thức mà một kiểu dữ liệu phải thực hiện.
		Interface shape trong ví dụ yêu cầu một phương thức area.
		2, Phương thức của struct: Kiểu square có một phương thức area trả về diện tích của hình vuông.
		3, Gán giá trị cho interface: Bạn có thể gán một giá trị cụ thể (trong trường hợp này là con trỏ tới square)
		cho một biến interface nếu kiểu của giá trị đó thực hiện interface.
		4, Truy xuất phương thức qua interface: Sau khi gán giá trị,
		bạn có thể gọi các phương thức của interface trên giá trị đó.
	*/
}

// example về interface
type Movement interface {
	up()
}

type Animal interface {
	speak()
}

// Embed Interface
type NextAnimal interface {
	Movement
	Animal
}

// Empty Interface: default trong go là kiểu này
type EmptyAnimal interface {
}

type Dog struct {
	color   string
	age     uint
	typeDog string
}

func (d *Dog) speak() {
	fmt.Printf("DOG: %v \n", d)
}
func (d *Dog) up() {

	fmt.Printf("DOG RUN UP: %v \n", d)
}

// Giống kiểu any trong typescript
func goOut(i interface{}) {
	fmt.Println(i)
}

type Cat struct {
}

func (d Cat) speak() {
	fmt.Print("CAT\n")
}

func (d Cat) up() {
	fmt.Print("CAT UP\n")
}

func interfaceExample() {
	// Sử dụng & trong interface: sử dụng trong các trường hợp cần truyền tham chiếu để tránh sao chép
	// nhiều dữ liệu hoặc cần thay đổi nội dung của struct.

	// Single Interface

	// var animal Animal
	// animal = Dog{}
	// animal.speak()

	//Multi Interface
	// dog := Dog{}

	// var m Movement = dog
	// m.up()
	// var a Animal = dog
	// a.speak()

	// Embed Interface (extend interface)
	dog := &Dog{
		color:   "blu",
		age:     18,
		typeDog: "Viet",
	}

	cat := Cat{}

	var n NextAnimal = dog
	n.up()
	n.speak()

	var c Cat = cat
	c.speak()
	c.up()

	// Empty Interface
	goOut(10)
}
