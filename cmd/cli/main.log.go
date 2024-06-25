package main

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// suger := zap.NewExample().Sugar()
	// suger.Infof("Hello name:%s, age: %d ", "Tip", 10) // Giống cấu trúc printf

	// //logger
	// loggerTest := zap.NewExample()
	// loggerTest.Info("Hello", zap.String("name", "tip go"), zap.Int("age", 40))

	// Cung cấp 3 config ở 3 môi trường.

	/*
		Kiểu 1: {"level":"info","msg":"Hello NewExample"}
		Kiểu 2: 2024-06-25T15:25:47.491+0700    INFO    cli/main.log.go:19      Hello NewDevelopment
		Kiểu 3: {"level":"info","ts":1719303947.491699,"caller":"cli/main.log.go:23","msg":"Hello NewProduction"}
	*/

	//2. Use Basic
	//Kiểu 1
	// logger1 := zap.NewExample()
	// logger1.Info("Hello NewExample")

	// //Dev
	// loggerDev, _ := zap.NewDevelopment()
	// loggerDev.Info("Hello NewDevelopment")

	// //Production
	// loggerProduction, _ := zap.NewProduction()
	// loggerProduction.Info("Hello NewProduction")

	//3, Customize
	encoder := getEncoderLog()
	sync := getWriterSync()

	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))

}

// customize format logs
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// convert 1719303947.491699 --> 2024-06-25T15:25:47.491+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// convert ts --> Time
	encodeConfig.TimeKey = "time"
	//from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//caller":"cli/main.log.go:23
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	// name: Tên file, ở đây là .log/log.txt.
	// flag: Cờ (flag) để xác định chế độ mở file. os.O_RDWR có nghĩa là mở file ở chế độ đọc-ghi.
	// perm: Quyền truy cập file. os.ModePerm cho phép tất cả quyền (read, write, execute) cho owner, group và others.

	// Hàm os.OpenFile được sử dụng để mở hoặc tạo một file mới tại đường dẫn
	// file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm) //O_RDWR

	// Tạo thư mục nếu chưa tồn tại
	logDir := "log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}

	// Mở file log với cờ phù hợp
	file, err := os.OpenFile(filepath.Join(logDir, "log.txt"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// Tạo một đối tượng WriteSyncer từ file đã mở, cho phép ghi log vào file này.
	syncFile := zapcore.AddSync(file)

	// Tạo một đối tượng WriteSyncer từ os.Stderr, cho phép ghi log ra console (tiêu chuẩn lỗi).
	syncConsole := zapcore.AddSync(os.Stderr)

	// Được sử dụng để kết hợp nhiều WriteSyncer thành một. Ở đây, nó kết hợp syncConsole và syncFile,
	// cho phép ghi log đồng thời vào cả console và file. Hàm này trả về một đối tượng WriteSyncer kết hợp.
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
