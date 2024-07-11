package main

import (
	"errors"
	"fmt"
)

type SyntaxError struct {
	msg    string // description of error
	Offset int64  // error occurred after reading Offset bytes
}

type Error interface {
	error
	Timeout() bool   // Is the error a timeout?
	Temporary() bool // Is the error temporary?
}

type error interface {
	Error() string
}

var ErrDivideByZero = errors.New("divide by zero")

func main() {
	// f, err := os.Open("filename.ext")
	// if err != nil {
	// 	log.Fatal("err:", err, f)
	// }
	// print: in ra open filename.ext: no such file or directory <nil> và dừng chương trình

	//2.
	// f1, err1 := Sqrt(-1)
	// if err != nil {
	// 	fmt.Println(err1, f1)
	// }

	// fmt.Errorf("math: square root of negative number %g", 1)
	// fmt.Sprintf("math: square root of negative number %g", float64(123))

	a, b := 10, 0
	result, err := Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)

}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// func Sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		return 0, errors.New("math: square root of negative number")
// 	}
// 	//implementation
// }

func DoSomething() error {
	return errors.New("something didn't work")
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}
	return a / b, nil
}
