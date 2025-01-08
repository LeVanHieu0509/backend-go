package main

import (
	"fmt"
	"sync"
	"time"
)

// Định nghĩa struct MessageReal
type MessageReal struct {
	orderID string
	Title   string
	Price   int
}

// Goroutine mua vé, gửi dữ liệu vào kênh buyChannel
func buyTicketReal(channel chan<- MessageReal, orders []MessageReal, wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu hoàn thành
	for _, order := range orders {
		time.Sleep(time.Second * 1) // Giả lập thời gian xử lý
		fmt.Printf("\nSend Buy::%s\n", order.orderID)
		channel <- order
	}
	close(channel) // Đóng kênh sau khi hoàn thành
}

// Goroutine hủy vé, gửi dữ liệu vào kênh cancelChannel
func cancelTicketReal(channel chan<- string, cancelOrders []string, wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu hoàn thành
	for _, orderId := range cancelOrders {
		time.Sleep(time.Second * 10) // Giả lập thời gian xử lý
		fmt.Printf("\nCancel ticket::%s\n", orderId)
		channel <- orderId
	}
	close(channel) // Đóng kênh sau khi hoàn thành
}

// Xử lý đồng thời các kênh bằng select
func handlerOrderSelectReal(orderChannel <-chan MessageReal, cancelChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu hoàn thành
	for {
		select {
		case order, orderOK := <-orderChannel:
			if orderOK {
				fmt.Printf("Handle Buy ticket::%s - %s - %d\n", order.orderID, order.Title, order.Price)
			} else {
				fmt.Printf("Order channel is closed.\n")
				orderChannel = nil
			}
		case cancel, cancelOK := <-cancelChannel:
			if cancelOK {
				fmt.Printf("Handle Cancel ticket::%s\n", cancel)
			} else {
				fmt.Printf("Cancel channel is closed.\n")
				cancelChannel = nil
			}
		}

		// Thoát nếu cả hai kênh đều đã đóng
		if orderChannel == nil && cancelChannel == nil {
			break
		}
	}
}

func main() {
	// Tạo các kênh
	buyChannel := make(chan MessageReal)
	cancelChannel := make(chan string)

	// Khởi tạo WaitGroup
	var wg sync.WaitGroup

	fmt.Println("Start buying ticket")

	// Danh sách đơn hàng mua và hủy
	buyOrders := []MessageReal{
		{orderID: "Order-01", Title: "Tips Go", Price: 30},
		{orderID: "Order-02", Title: "Tips NodeJs", Price: 31},
		{orderID: "Order-03", Title: "Tips Java", Price: 32},
		{orderID: "Order-04", Title: "Tips Php", Price: 33},
	}

	cancelOrders := []string{"Order-01", "Order-02"}

	// Tăng số lượng công việc cho WaitGroup (3 goroutines).
	wg.Add(3)

	// Chạy các goroutines
	go buyTicketReal(buyChannel, buyOrders, &wg)
	go cancelTicketReal(cancelChannel, cancelOrders, &wg)
	go handlerOrderSelectReal(buyChannel, cancelChannel, &wg)

	// Đợi tất cả goroutines hoàn thành
	// chương trình kết thúc khi tất cả các công việc trong WaitGroup hoàn thành
	wg.Wait()

	fmt.Println("End buying and canceling ...")
}
