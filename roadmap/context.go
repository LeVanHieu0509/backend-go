package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
Mã Go này mô phỏng quá trình truy xuất dữ liệu của người dùng từ một bên thứ ba chậm
và sử dụng ngữ cảnh (context) để kiểm soát thời gian chờ.
Dưới đây là phân tích chi tiết từng phần của mã:
*/
func main() {
	// Ghi nhận thời gian bắt đầu
	start := time.Now()

	// Tạo context với giá trị "foo" là "bar"
	ctx := context.WithValue(context.Background(), "foo", "bar")

	// Giả định một userID là 10
	userID := 10

	val, err := fetchUserData(ctx, userID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", val)             // In kết quả
	fmt.Println("took:", time.Since(start)) // In thời gian đã thực hiện

}

type Response struct {
	value int
	err   error
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	// Giả lập quá trình truy xuất chậm bằng cách ngủ 150 microsecond
	time.Sleep(time.Microsecond * 150)

	return 666, nil
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	// Lấy giá trị "foo" từ context
	value := ctx.Value("foo")

	// In giá trị đó
	fmt.Println(value.(string))

	// Tạo context con với timeout là 200 milliseconds
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)

	// Đảm bảo hàm cancel() sẽ được gọi để giải phóng tài nguyên context
	defer cancel()

	// Tạo một channel respch để nhận kết quả từ goroutine
	respch := make(chan Response)

	// Khởi chạy một goroutine gọi hàm fetchThirdPartyStuffWhichCanBeSlow() và gửi kết quả (gồm value và err) vào channel respch.
	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	// Sử dụng vòng lặp for với select để chờ nhận dữ liệu từ channel hoặc nhận thông báo hết thời gian từ context.
	for {
		select {

		// Nếu context hết thời gian (timeout), trả về lỗi
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took to long")

		// Nếu nhận được kết quả từ channel respch, trả về giá trị và lỗi (nếu có) từ response.
		case resp := <-respch:
			return resp.value, resp.err
		}
	}

}

/*
1. Truyền giá trị: Giá trị "foo" với giá trị "bar" được truyền qua context
và được lấy ra trong hàm fetchUserData để in ra màn hình.
2. Quản lý thời gian chờ: Context được sử dụng để tạo một timeout 200 milliseconds.
Nếu việc lấy dữ liệu từ hàm fetchThirdPartyStuffWhichCanBeSlow mất quá nhiều thời gian,
hàm fetchUserData sẽ hủy bỏ hoạt động đó và trả về lỗi "fetching data from third party took too long".
3. Hủy bỏ: Nếu việc lấy dữ liệu mất quá nhiều thời gian và context bị hết hạn,
goroutine sẽ bị hủy bỏ và tài nguyên sẽ được giải phóng thông qua việc gọi cancel().
*/
