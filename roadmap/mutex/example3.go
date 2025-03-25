package main

import (
	"fmt"
	"sync"
	"time"
)

type intLock struct {
	val int
	sync.Mutex
}

func (n *intLock) isEven() bool {
	return n.val%2 == 0
}

func main() {
	n := &intLock{val: 0}

	go func() {
		n.Lock()
		defer n.Unlock()
		nIsEven := n.isEven()
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n.val, " is even")
			return
		}
		fmt.Println(n.val, "is odd")
	}()

	go func() {
		n.Lock()
		n.val++
		n.Unlock()
	}()

	time.Sleep(time.Second)
}

/*
	1. Nếu bạn có nhiều trường hợp dữ liệu cần quyền truy cập độc quyền,
	2. Việc đóng gói mutex cùng với chính dữ liệu sẽ giúp dữ liệu bớt khó hiểu hơn và dễ đọc hơn.
	3. Tóm lại, mutex là một công cụ tuyệt vời để ngăn chặn việc truy cập dữ liệu không theo thứ tự.
	Có nhiều cách để sử dụng mutex và nhiều cạm bẫy có thể xảy ra,
	vì vậy hãy đảm bảo đánh giá trường hợp sử dụng của bạn trước khi quyết định phương pháp phù hợp.
*/
