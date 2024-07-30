package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

/*
1. Giải Phóng Bộ Nhớ: Khi làm việc với tài nguyên như bộ nhớ động hoặc tài nguyên bên ngoài, bạn có thể sử dụng defer để đảm bảo tài nguyên được giải phóng đúng cách.
2. Đóng Kết Nối Mạng
3. Mở Khóa Mutex: Khi làm việc với các phần tử đồng thời (concurrency), bạn thường cần mở khóa một mutex để tránh deadlock.
4. Xử Lý Hệ Thống Tập Tin Tạm Thời
5. Ghi Log: Ghi nhật ký khi thoát khỏi hàm, hữu ích trong việc theo dõi luồng thực thi hoặc ghi nhận sự kiện.
6. Hoàn Tác Các Thay Đổi: Khi thực hiện một loạt các thay đổi có thể cần hoàn tác nếu xảy ra lỗi.
7. Đo Lường Thời Gian Thực Thi
8.  Thực Hiện Nhiều Hành Động Deferred
*/
func main() {
	// deferFile()
	deferMutex()
	MeasureTime()

	db := NewDatabase()
	db.data["key1"] = "value1"
	db.data["key2"] = "value2"

	fmt.Println("Trạng thái ban đầu:", db.data)

	err := updateDatabase(db)
	if err != nil {
		log.Println("Lỗi:", err)
	}

	fmt.Println("Trạng thái sau khi cập nhật:", db.data)

	// demo các defer sẽ được push vào stack
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d \n", i)
	}

	// a()
	fmt.Println("---------------------B-------")
	b()
	fmt.Println("---------------------B-------")
}

func deferFile() {
	file, err := os.Open("../log/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

// Mở Khóa Mutex: Khi làm việc với các phần tử đồng thời (concurrency), bạn thường cần mở khóa một mutex để tránh deadlock.
func deferMutex() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
}

// 7.
func trackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func MeasureTime() {
	defer trackTime(time.Now(), "trackTime")

	// Thực hiện các hoạt động trong hàm
}

// func updateRecord() {
// 	// Lưu trạng thái ban đầu
// 	originalState := getState()
// 	defer restoreState(originalState)

// 	// Thực hiện các thay đổi
// 	if err := applyChanges(); err != nil {
// 		return // restoreState sẽ được gọi tại đây
// 	}

// 	// Thay đổi thành công, không cần hoàn tác
// }

// ví dụ về restore data
type Database struct {
	data map[string]string
}

func NewDatabase() *Database {
	return &Database{data: make(map[string]string)}
}

func (db *Database) getState() map[string]string {
	state := make(map[string]string)
	for k, v := range db.data { // Sửa lỗi: thêm từ khóa :=
		state[k] = v
	}
	return state
}

func (db *Database) restoreState(state map[string]string) {
	db.data = state
}

func (db *Database) applyChanges() error {
	// Thực hiện các thay đổi trên cơ sở dữ liệu
	db.data["key1"] = "newValue1"
	db.data["key2"] = "newValue2"

	// thành công
	// return nil

	//failed
	// Gây ra lỗi giả lập
	return errors.New("cập nhật không thành công")
}

func updateDatabase(db *Database) error {
	// Lưu trạng thái ban đầu của cơ sở dữ liệu
	originalState := db.getState()
	defer func() {
		if originalState != nil {
			// Điều này đảm bảo rằng dữ liệu của bạn vẫn giữ nguyên trạng thái nhất quán.
			db.restoreState(originalState)
		}
	}()

	// Thực hiện các thay đổi
	if err := db.applyChanges(); err != nil {
		return err // restoreState sẽ được gọi tại đây
	}

	// Nếu không có lỗi, hoãn defer để không hoàn tác thay đổi
	defer func() {
		originalState = nil
	}()

	// Thay đổi thành công, không cần hoàn tác
	return nil
}

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
	trace("a")
	defer untrace("a")

	// do something....
}

// / process

/*
1. trace1("b") in ra "entering: b".
2. defer un("b") được thiết lập để chạy khi b kết thúc.
3. In ra "in b".
4. Gọi hàm a1.
5. trace1("a") in ra "entering: a".
6. defer un("a") được thiết lập để chạy khi a1 kết thúc.
7. In ra "in a".
8. a1 kết thúc và un("a") được gọi, in ra "leaving: a".
9. b kết thúc và un("b") được gọi, in ra "leaving: b".
*/

/*
Note: defer đảm bảo rằng hàm un sẽ được gọi khi hàm hiện tại kết thúc, bất kể hàm đó kết thúc theo cách nào.
Trong ví dụ này, trace1 và un được sử dụng để theo dõi việc vào và ra khỏi các hàm a1 và b.
Điều này rất hữu ích cho việc gỡ lỗi và theo dõi luồng thực thi trong chương trình của bạn.

--- Hàm hiện tại ---
1. main()
2. b() (khi b() được gọi từ main())
3. a1() (khi a1() được gọi từ b())
4. a1 kết thúc và b tiếp tục là hàm hiện tại
5. b kết thúc và main tiếp tục là hàm hiện tại, sau đó main kết thúc
6. Chức năng defer đảm bảo các hàm un được gọi ngay trước khi hàm hiện tại kết thúc, giúp dễ dàng quản lý tài nguyên và theo dõi luồng thực thi.
*/

func trace1(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a1() {
	// defer là một từ khóa đặc biệt trong Go dùng để trì hoãn việc thực thi một hàm cho đến khi hàm bao quanh nó (a1 trong trường hợp này) kết thúc.
	// trace1("a") được thực thi ngay lập tức và không phải là hàm mà chúng ta muốn trì hoãn.
	// un("a") là hàm mà chúng ta muốn trì hoãn đến cuối hàm a1, do đó chúng ta sử dụng defer trước hàm un.
	defer un(trace1("a"))
	fmt.Println("in a")
}

func b() {
	// hàm un sẽ được gọi khi hàm hiện tại kết thúc
	defer un(trace1("b"))
	fmt.Println("in b")
	a1()
}
