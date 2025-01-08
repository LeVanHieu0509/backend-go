package main

import (
	"fmt"
	"time"
)

// pub sub use channel and go routines

// giống interface để định nghĩa các biến trong go
type Message struct {
	orderID string
	Title   string
	Price   int
}

// buyTicket gửi các đơn hàng vào buyChannel.
func buyTicket(channel chan<- Message, orders []Message) {
	// Lặp qua danh sách các đơn hàng (orders) và gửi từng đơn hàng (Message) vào channel.
	for _, order := range orders {
		time.Sleep(time.Second * 1)
		fmt.Printf("\nSend Buy::%s\n", order.orderID)

		channel <- order
	}
	close(channel)
}

// cancelTicket gửi mã hủy vào cancelChannel
func cancelTicket(channel chan<- string, cancelOrders []string) {
	// Lặp qua danh sách mã đơn hàng cần hủy (cancelOrders) và in thông báo hủy.
	for _, orderId := range cancelOrders {
		time.Sleep(time.Second * 10) // giả lập network chậm
		fmt.Printf("\nCancel ticket::%s\n", orderId)
		channel <- orderId
	}

	close(channel)
}

func handlerOrder(orderChannel <-chan Message, cancelChanel <-chan string) {
	for {
		// Đọc từ orderChannel để xử lý đơn hàng
		order, orderOK := <-orderChannel
		if orderOK {
			fmt.Printf("Handle Buy ticket::%s - %s - %d\n", order.orderID, order.Title, order.Price)
		} else {
			fmt.Printf("Order chanel is closed.")
			break
		}

		// Đọc từ cancelChannel để xử lý yêu cầu hủy đơn hàng.
		cancel, cancelOK := <-cancelChanel //nếu chạy qua mà không có dữ liệu để xử lý thì sẽ bị đứng hình -> bị trì hoãn
		if cancelOK {
			fmt.Printf("Handle Cancel ticket::%s\n", cancel)
		} else {
			fmt.Printf("Cancel chanel is closed.")
			break
		}
	}
}

// Khi có dữ liệu từ bất kỳ kênh nào, handlerOrderSelect xử lý ngay lập tức mà không làm gián đoạn các kênh khác
func handlerOrderSelect(orderChannel <-chan Message, cancelChanel <-chan string) {
	for {
		// Khi một kênh không có dữ liệu, select sẽ chờ dữ liệu từ các kênh khác, tránh tình trạng "đứng hình" (deadlock) khi một kênh không có dữ liệu.
		select {
		// Đọc từ orderChannel để xử lý đơn hàng
		case order, orderOK := <-orderChannel:
			if orderOK {
				fmt.Printf("Handle Buy ticket::%s - %s - %d\n", order.orderID, order.Title, order.Price)
			} else {
				fmt.Printf("Order chanel is closed.")
				orderChannel = nil
			}

		// Đọc từ cancelChannel để xử lý yêu cầu hủy đơn hàng.
		case cancel, cancelOK := <-cancelChanel: //nếu chạy qua mà không có dữ liệu để xử lý thì sẽ bị đứng hình -> bị trì hoãn
			if cancelOK {
				fmt.Printf("Handle Cancel ticket::%s\n", cancel)
			} else {
				fmt.Printf("Cancel chanel is closed.")
				cancelChanel = nil
			}
		}

		// exit when all chanel is closed.
		if orderChannel == nil && cancelChanel == nil {
			break
		}
	}
}

// này là chúa -> chúa chết thì các go routines đều sẽ chết theo
func main() {
	buyChannel := make(chan Message)
	cancelChannel := make(chan string)
	fmt.Println("start buying ticket")

	buyOrders := []Message{
		{orderID: "Order-01", Title: "Tips Go", Price: 30},
		{orderID: "Order-02", Title: "Tips NodeJs", Price: 31},
		{orderID: "Order-03", Title: "Tips Java", Price: 32},
		{orderID: "Order-04", Title: "Tips Php", Price: 33},
	}

	cancelOrders := []string{"Order-01", "Order-02"}

	/*
		- Khi chạy đồng thời 3 go routing thì sẽ không biết được goroutines nào chạy trước
		- nên khi goroutines chạy thì sẽ phải dựa vào channel để biết được goroutines nào chạy trước và chạy sau.
		- 1 goroutines cho phép nhận nhiều channel thông qua từ khoá select
	*/

	go buyTicket(buyChannel, buyOrders)
	go cancelTicket(cancelChannel, cancelOrders)
	// go handlerOrder(buyChannel, cancelChannel)
	go handlerOrderSelect(buyChannel, cancelChannel)

	time.Sleep(time.Second * 25)
	fmt.Println("End buying and canceling ...")

}

/*

Lệnh hủy (Cancel ticket::Order-01) bị xử lý sau khi đơn hàng thứ hai (Order-02)
được gửi nhưng không đồng bộ với hành động "buy".

Goroutines không chờ nhau hoàn thành.
Điều này dẫn đến việc một số hành động bị bỏ qua (ví dụ: không xử lý Order-03 và Order-04).

Các thông báo "Handle" (xử lý) cho một số đơn hàng không xuất hiện.
Chương trình kết thúc trước khi các goroutines hoàn thành.
*/
