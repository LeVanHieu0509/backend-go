package initialize

import (
	"fmt"
	"time"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)

	}
}
func InitMySql() {
	m := global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false, //Tính nhất quán dữ liệu nên set true (nếu tắt thì cải thiện tầm 305 hiệu suất)

	})

	checkErrorPanic(err, "Init mysql initialization error")

	global.Logger.Info("Initializing Mysql Successfully!")
	global.Mdb = db

	// set Pool (mở nhóm kết nối giúp hiệu suất tăng lên rất nhiều => Mở sẵn các kết nối cho user vào sài)
	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql error: %s ::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

// po: persistent object: đại diện cho một đối tượng thực thể được lưu trữ và truy xuất trong csdl
// Mục đích để đóng gói dữ liệu, mỗi file trong folder sẽ tương ứng với 1 bảng. trong các field của file này tương ứng với 1 trường ở trong bảng.

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		fmt.Println("Migrate tables error:", err)
	}
}
