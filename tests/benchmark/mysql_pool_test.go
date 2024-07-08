package main

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "Tipjs"}

	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

func BenchmarkMaxOpenConns1(b *testing.B) {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "tipjs:tipjs@tcp(127.0.0.1:3306)/shopGo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SkipDefaultTransaction: false, //Tính nhất quán dữ liệu nên set true (nếu tắt thì cải thiện tầm 305 hiệu suất)

	})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	//Check if table exists
	if db.Migrator().HasTable(&User{}) {
		//Drop table if it exists
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// Handle error if you want
			// fmt.Println("Error dropping table")
		}
	}

	// Create table if not exists
	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Failed to get SQL.DB from gorm.DB: %v", err)
	}

	// Setup arguments connect to db
	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

//go test -bench=. -benchmem
// Kiểm tra insert được bao nhiêu record trong vòng bấy nhiêu s và memory và tỉ lệ phản hồi trong cái lần đầu tiên

/*
	BenchmarkMaxOpenConns1-8             548           1986118 ns/op            5905 B/op         77 allocs/op
PASS
ok      github.com/LeVanHieu0509/backend-go/tests/benchmark     2.826s

- 548: Hàm benchmark này đã chạy được bao nhiêu lần, là số lần insert thành công
- 1986118 ns/op: Thời gian trung bình mà mỗi lần chạy hàm Benchmark hết bao nhiêu nano giây - cho biết hiệu suất mỗi khi chạy hàm bench
- 5905 B/op: số byte được phân bổ cho mỗi lần chạy hàm - chạy càng cao thì càng kém
- 77 allocs/op: Số lần cầp phát bộ nhớ - đánh giá hiệu quả sử dụng của bộ nhớ chạy khi nào.
-

*/
