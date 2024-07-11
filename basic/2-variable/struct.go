package main

import "fmt"

type Student struct {
	Name  string
	Age   uint
	IsGay bool
	Msg   string
}

func main() {

	var student = Student{Name: "Hieu", Age: 18, IsGay: false, Msg: "pro"}

	// s là con trỏ => Student
	// là 1 biến lưu trữ địa chỉ vùng nhớ của 1 biến khác
	var pointer *Student

	// s sẽ trỏ về địa chỉ student
	//&: Lấy địa chỉ của biến student và gán cho con trỏ b (& là toán tử lấy địa chỉ, nó trả về địa chỉ bộ nhớ của biến a)
	pointer = &student

	fmt.Println("student", student)
	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Printf("%T\n", pointer)

	var c = student

	var pointer2 *Student = &c //là một con trỏ trỏ tới c.
	apply(pointer2)

	fmt.Println(c)
	c.GetAge()

	var d = 30

	//pointer3 là một con trỏ đến kiểu int và được gán địa chỉ của d.
	var pointer3 *int = &d
	applyInt(pointer3)
	fmt.Println(d)

}

func (p *Student) GetAge() {
	fmt.Println(p.Age)

}

// Cả hai hàm đều nhận vào con trỏ và thay đổi giá trị mà con trỏ trỏ tới.

// Hàm apply: Không cần sử dụng dấu * khi truy cập các trường của struct thông qua con trỏ vì Go tự động giải tham chiếu con trỏ khi truy cập các trường của struct.
// Hàm applyInt: Cần sử dụng dấu * để giải tham chiếu con trỏ và truy cập giá trị thực tế của biến kiểu nguyên thủy (như int) mà con trỏ trỏ tới.

func apply(pointer *Student) {
	// Trong trường hợp này, không cần sử dụng dấu * trước pointer vì trong Go,
	// khi chúng ta truy cập trường của struct thông qua con trỏ,
	// ngôn ngữ tự động giải tham chiếu (dereference) con trỏ cho chúng ta.

	pointer.Age = 20
}

// Thay đổi giá trị trong biến mà con trỏ này trỏ tới
func applyInt(pointer *int) {
	// Trong trường hợp này, cần sử dụng dấu * để giải tham chiếu con trỏ và
	// truy cập giá trị thực tế của biến int mà con trỏ pointer trỏ tới.
	*pointer = 50
}
