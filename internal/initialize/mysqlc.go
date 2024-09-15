package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/model"
	"go.uber.org/zap"
	"gorm.io/gen"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)

	}
}

// Hàm: InitMySql khởi tạo kết nối cơ sở dữ liệu MySQL
func InitMySqlC() {
	// Lấy cấu hình MySQL từ global.Config.Mysql.
	m := global.Config.Mysql

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// Định dạng chuỗi Data Source Name (DSN) sử dụng cấu hình MySQL.
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)

	// Mở kết nối tới cơ sở dữ liệu MySQL sử dụng GORM.
	db, err := sql.Open("mysql", s)

	checkErrorPanic(err, "Init mysql initialization error")

	global.Logger.Info("Initializing Mysql Successfully!")

	// Gán kết nối cơ sở dữ liệu vào global.Mdb
	global.Mdbc = db

	// set Pool (mở nhóm kết nối giúp hiệu suất tăng lên rất nhiều => Mở sẵn các kết nối cho user vào sài)
	SetPool()

	// Gọi hàm migrateTables để tự động di chuyển các bảng cơ sở dữ liệu.
	// migrateTables()

	//
	// genTableDao()
}

// Hàm: SetPool cấu hình nhóm kết nối MySQL
func SetPoolC() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql error: %s ::", err)
	}

	// Lấy kết nối cơ sở dữ liệu SQL gốc từ GORM.
	// Đặt số lượng kết nối nhàn rỗi tối đa, số lượng kết nối mở tối đa, và thời gian sống tối đa của kết nối dựa trên cấu hình.

	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

// po: persistent object: đại diện cho một đối tượng thực thể được lưu trữ và truy xuất trong csdl
// Mục đích để đóng gói dữ liệu, mỗi file trong folder sẽ tương ứng với 1 bảng. trong các field của file này tương ứng với 1 trường ở trong bảng.

// go to mysql
func migrateTablesC() {
	// Gọi global.Mdb.AutoMigrate để di chuyển các bảng dựa trên các mô hình po.User và po.Role.
	// Đại diện cho các thực thể được lưu trữ và truy xuất từ cơ sở dữ liệu.
	// Mỗi file trong gói po tương ứng với một bảng trong cơ sở dữ liệu, với mỗi trường trong file đại diện cho một cột trong bảng.

	err := global.Mdb.AutoMigrate(
		// &po.User{},
		// &po.Role{},
		&model.GoCrmUser{}, // có thể switch qua switch lại giữ go và mysql. Nên ưu tiên làm việc bên mysql
	)

	if err != nil {
		fmt.Println("Migrate tables error:", err)
	}
}

// mysql to go
func genTableDaoC() {
	// Init gen table
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	g.GenerateAllTable()
	g.GenerateModel("go_crm_user")
	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}
