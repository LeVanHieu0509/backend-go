package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Hàm main bản chất là 1 go routing chúa => sẽ chạy từ main vào trước
// Khi hàm main kết thúc thì các go routing con cũng kết thúc

// khi mà main kết thúc rồi mà gorouting con chưa kết thúc thì sẽ ko in ra được kết quả vì main kết thúc trước gorouting
// Main kết thúc thì các go routing con trong hệ thống đều kết thúc hết

func main() {

	// cach 1: Finished ... 5.005119875s
	fmt.Println("Starting 1 ...")

	ids := []int{1, 2, 3, 4, 5} //productId

	start := time.Now()

	for _, id := range ids {
		getProductByIdAPI(id)
	}

	fmt.Println("Finished 1 ...", time.Since(start))

	// cach 2: go routing - được người dùng tạo ra chứ không phải là do hệ thống tạo ra
	// Finished 2 ... 1.0012165s

	fmt.Println("Starting 2 ...")
	start2 := time.Now()

	for _, id := range ids {
		go getProductByIdAPI(id)

	}

	// Tại sao chỗ này có thì mới in ra kết quả?
	time.Sleep(time.Second + 1)

	// --> giải thích:
	fmt.Println("Finished 2 ...", time.Since(start2))

	// wait group
	waitGroup()

}

// go routing cha
// note: Khi mà go routing cha kết thúc mà go routing con vẫn sẽ chạy cho đến khi nào xong thì thôi.
func getProductByIdAPI(id int) {
	fmt.Println(">>> Data ProductId: ", id)
	sleep(1)
}

// go routing con
func sleep(s time.Duration) {
	time.Sleep(time.Second + s)
}

func getWaitGroup(id int, wg *sync.WaitGroup) {
	// Khi công việc trong goroutine hoàn tất, giảm số lượng công việc đang chờ trong WaitGroup.
	defer wg.Done()

	fmt.Println(">>> Data ProductId: ", id)
	sleep(1)
}

func waitGroup() {
	fmt.Println("Starting ...")

	// 1. Khai bao  WaitGroup sử dụng để quản lý trạng thái của các goroutine.
	var wg sync.WaitGroup

	ids := []int{1, 2, 3, 4, 5}

	start := time.Now()

	for _, id := range ids {
		// Tăng số lượng công việc đang chờ lên 1, thông báo rằng một goroutine mới sẽ được thêm vào.
		wg.Add(1)

		// Việc sử dụng goroutine cho phép thực hiện song song các công việc.

		// go getWaitGroup(id, &wg)

		// Ví dụ thực tế, có thể áp dụng được hàng trăm ngàn tác vụ đồng thời.
		// Phải quản lý bộ nhớ thật tốt và chặn chẽ, phải control được CPU

		go getProductById(id, &wg)
	}

	// Chờ đến khi tất cả các goroutine báo hoàn thành (khi số lượng công việc chờ trong WaitGroup giảm về 0).
	wg.Wait()
	fmt.Println("Finished WaitGroup ...", time.Since(start))

}

func getProductById(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return // Thoát nếu có lỗi trong quá trình gửi yêu cầu HTTP
	}
	defer resp.Body.Close() // Đảm bảo đóng response body sau khi đọc xong

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return // Thoát nếu không đọc được nội dung
	}

	fmt.Printf(">>> Data ProductId: %d: %s\n", id, string(body))
}

/*
	1. Đồng thời: là các chương trình đa luồng chạy trên 1 lõi của CPU
		- 1 người làm 1 lúc multi task (vừa đánh bài, vừa hút thuốc)
	2. Song song: Là chương trình đa luồng chạy trên nhiều lõi của CPU (4 người làm 1 lúc multi task)
		- 1 người (đánh bài hút thuốc), 1 người (uống nước, đánh bài)
		- 1 người (bấm điện thoại, hút thuốc), 1 người (uống nước, bấm điện thoại)

	3. wait group
		-  là công cụ trong Go để đồng bộ hóa các goroutine.
		- Nó đảm bảo rằng chương trình chính chờ cho đến khi tất cả các goroutine được quản lý bởi WaitGroup hoàn thành.


	4. channel: 2 ông go routing sẽ nói chuyện với nhau thông qua channel
*/
