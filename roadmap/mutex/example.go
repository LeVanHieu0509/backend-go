package main

import (
	"fmt"
	"sync"
	"time"
)

// a simple function that returns true if a number is even
func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0
	var m sync.Mutex

	// goroutine 1
	// reads the value of n and prints true if its even
	// and false otherwise
	go func() {
		m.Lock()
		defer m.Unlock()

		nIsEven := isEven(n)
		// we can simulate some long running step by sleeping
		// in practice, this can be some file IO operation
		// or a TCP network call
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	// goroutine 2
	// modifies the value of n
	go func() {
		m.Lock()
		n++
		m.Unlock()
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}

/*
	- 2 go routines trên đang được call cùng 1 thời điểm
	- go routine 2 đang thay đổi giá trị của n của go routine 1 nên sẽ bị xung đột dữ liệu.
	- Sử dụng khoá mutex để giải quyết vấn đề trên

	Solution:
	1. Khi m.Lock được gọi thì nó sẽ chặn các luồng khác lại theo hàng đợi cho đến khi nào m.Unlock được gọi.
	2. Điều này đảm bảo rằng chỉ có một go routine có thể truy cập vào biến n tại một thời điểm.
*/
