package main

import (
	"fmt"
	"time"
)

func main() {

	selectExample()

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selectExample() {
	c1 := make(chan string) // Tạo channel c1 kiểu string
	c2 := make(chan string) // Tạo channel c2 kiểu string

	// hai goroutine sẽ được khởi tạo đồng thời.
	// Goroutine đầu tiên gửi "one" vào c1 sau 1 giây
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	// Goroutine thứ hai gửi "two" vào c2 sau 2 giây
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Vòng lặp 2 lần, sử dụng select để nhận giá trị từ cả 2 channel
	for i := 0; i < 2; i++ {
		// select sẽ chờ và thực thi trường hợp (case) nào có dữ liệu đến trước,
		// sau đó in ra thông báo và tiếp tục chờ lần tiếp theo.
		select {
		case msg1 := <-c1: // Nếu có dữ liệu từ c1, in ra và tiếp tục
			fmt.Println("received", msg1)
		case msg2 := <-c2: // Nếu có dữ liệu từ c2, in ra và tiếp tục
			fmt.Println("received", msg2)
		}
	}

}

/*
	- select giúp ta chờ và chọn lựa các channel mà có dữ liệu đến,
	giúp tối ưu trong việc xử lý các tác vụ đồng thời mà không phải chờ đợi một cách tuần tự.
	- Trong trường hợp này, mặc dù hai goroutine chạy đồng thời, nhưng select đảm bảo rằng ta
	sẽ nhận giá trị từ channel đầu tiên có dữ liệu mà không cần phải đợi cả hai goroutine hoàn thành.
*/
