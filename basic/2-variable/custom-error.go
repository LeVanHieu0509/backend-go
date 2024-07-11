package main

import (
	"errors"
	"fmt"
)

// là một cấu trúc (struct) chứa ba trường
type DivisionError struct {
	IntA int    //Lưu trữ giá trị số nguyên đầu tiên.
	IntB int    //Lưu trữ giá trị số nguyên thứ hai (mẫu số).
	Msg  string //Lưu trữ thông báo lỗi.
}

// Phương thức Error giúp DivisionError thỏa mãn giao diện error trong Go.
// Nó trả về thông báo lỗi được lưu trữ trong trường Msg.

func (e *DivisionError) Error() string {
	return e.Msg
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{
			IntA: a,
			IntB: b,
			Msg:  fmt.Sprintf("cannot divide '%d' by zero", a),
		}
	} else {
		return a / b, nil
	}
}

func main() {
	a, b := 10, 2
	result, err := Divide(a, b)

	if err == nil {
		fmt.Printf("%d / %d = %d", a, b, result)
	} else {
		// Khai báo biến divErr để lưu trữ lỗi loại DivisionError.
		var divErr *DivisionError

		switch {
		// Kiểm tra xem lỗi có phải là DivisionError hay không. Nếu phải, in ra thông báo lỗi chi tiết.
		case errors.As(err, &divErr):
			fmt.Printf("%d / %d is not mathematically valid: %s\n", divErr.IntA, divErr.IntB, divErr.Error())
		default:
			fmt.Printf("un expected divide error: %s\n", err)
		}
	}

}
