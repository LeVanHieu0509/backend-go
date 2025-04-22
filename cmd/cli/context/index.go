package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// func OrderHandlerSelect(w http.ResponseWriter, r *http.Request) {
// 	orderID := "GO-12345"
// 	resultChan := make(chan error, 1)

// 	// Xử lý đơn hàng trong goroutine
// 	go func() {
// 		err := placeOrderWithoutContext(orderID)
// 		resultChan <- err
// 	}()

// 	select {
// 	case err := <-resultChan:
// 		if err != nil {
// 			log.Printf("Xử lý đơn hàng %s thất bại\n", orderID)
// 			http.Error(w, "Lỗi xử lý đơn hàng", http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Đặt hàng thành công!"))

// 	case <-time.After(2 * time.Second):
// 		log.Printf("Xử lý đơn hàng %s quá 2 giây, trả lời lại với client\n", orderID)
// 		http.Error(w, "Quá thời gian xử lý, vui lòng thử lại sau", http.StatusGatewayTimeout)
// 	}
// }

// OrderHandlerWithContext xử lý yêu cầu HTTP với context và timeout
func OrderHandlerWithContext(w http.ResponseWriter, r *http.Request) {
	orderID := "GO-12345"

	// Tạo context có timeout 2 giây
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	// Gọi hàm đặt hàng với context
	err := placeOrderWithContext(ctx, orderID)
	if err != nil {
		log.Printf("Xử lý đơn hàng %s thất bại: %v\n", orderID, err)
		http.Error(w, "Lỗi xử lý đơn hàng hoặc quá thời gian", http.StatusGatewayTimeout)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Đặt hàng thành công!"))
}

// placeOrderWithContext thực hiện việc xử lý đơn hàng với context và timeout
func placeOrderWithContext(ctx context.Context, orderID string) error {
	log.Printf("Bắt đầu xử lý đơn hàng: %s\n", orderID)

	select {
	case <-time.After(3 * time.Second): // Giả lập xử lý mất 3 giây
		log.Printf("Xử lý đơn hàng %s thành công\n", orderID)
		return nil
	case <-ctx.Done(): // Context bị hủy
		log.Printf("Hủy xử lý đơn hàng %s: %v\n", orderID, ctx.Err())
		return ctx.Err()
	}
}

func main() {
	// http.HandleFunc("/order", OrderHandlerWithContext)

	// log.Print("Server đang chạy tại http://localhost:8001")
	// log.Fatal(http.ListenAndServe(":8000", nil))

	ctx := context.Background() // root -> request HTTP

	// Tạo context có timeout 2 giây
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	orderID := "GO-12345"
	err := placeOrderWithContext(ctx, orderID)

	if err != nil {
		log.Printf("Xử lý đơn hàng %s thất bại: %v\n", orderID, err)
	} else {
		log.Printf("Xử lý đơn hàng %s thành công: %v\n", orderID)
	}

	deadline := time.Now().Add(5 * time.Second)
	_, cancel1 := context.WithDeadline(context.Background(), deadline)
	defer cancel1()
}

/*
	Sử dụng context để tránh gọi vô nghĩa giống như metadata trong nestJS

	CLient -> Handle -> Order
	Context thường thì sử dụng thông qua thư viện redis, rabbitMQ, Kafka, Mysql đều bắt buộc phải truyền context vào.

	Mục đích:
	1, Cho phép huỷ bỏ tác vụ nếu các tác vụ hoặc gorouting không còn cần thiết nữa
	2, Giới hạn thời gian 1 tác vụ có thể chạy
	3, Truyền thông tin giữ các gorouting

	Cấu trúc:
	1, Context là một interface trong go
	2, Done(): Trả về 1 channel khi context bị huỷ bỏ hoặc hết thời gian chờ các gorouting có thể lắng nghe
	3, Err(): Trả về lỗi của context nếu nó đã bị huỷ hoặc hết thời gian
	4, Value(): Trả về giá trị được lưu giữ trong context

*/

/*
	1. context.Background():
	- Thường được sử dụng trong hàm main, trong các hàm khởi tạo,
	hoặc làm context gốc cho các request HTTP.
	- Trả về một context rỗng, không có deadline và không có bất kỳ giá trị nào
	2. TODO:
	- Trả về một context rỗng.
	- Thường được sử dụng khi không biết rõ context nào cần sử dụng
	- hoặc khi chưa quyết định được việc hủy bỏ/timeout.
	3. WithCancel()
	- Tạo một context mới có thể bị hủy bỏ.
	Khi gọi cancel(), tất cả các context con của nó cũng sẽ bị hủy bỏ.
	- Sử dụng khi bạn cần hủy bỏ tác vụ khi không còn cần thiết nữa
	4. WithDeadline()
	- Tạo một context có deadline, tức là khi thời gian hết hạn, context sẽ bị hủy bỏ.
	- Thường được sử dụng khi bạn muốn giới hạn thời gian xử lý.
	5. WithTimeout()
	- Tạo một context với thời gian chờ (timeout).
	Nếu tác vụ không hoàn thành trong thời gian này, context sẽ bị hủy bỏ.
	6. WithValue()
	- Tạo một context mới với giá trị gắn kèm.
	- Bạn có thể lưu trữ các giá trị trong context để truyền giữa các goroutines hoặc các hàm.
	- ctx := context.WithValue(context.Background(), "key", "value")

*/

/*
	1. context trong Go chủ yếu dùng để kiểm soát vòng đời của các tác vụ bất đồng bộ
	và quản lý các yêu cầu như hủy bỏ, timeout, và truyền giá trị giữa các hàm hoặc goroutines.
	2. Cần tránh lưu trữ context trong struct, thay vào đó, context nên được truyền trực tiếp qua các hàm.
	3. context rất hữu ích trong các ứng dụng web hoặc phân tán,
	nơi bạn cần đồng bộ hóa các tác vụ và đảm bảo rằng không có tác vụ nào chạy quá lâu
	hoặc chiếm tài nguyên một cách không cần thiết.
*/
