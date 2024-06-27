# backend-go

## LIBRARY

1. viper: go get github.com/spf13/viper
2. gin: go get -u github.com/gin-gonic/gin
3. logger: go get -u go.uber.org/zap
4. logger advance: go get github.com/natefinch/lumberjack

## Go (3): GIN vs ROUTER

run server: go run cmd/server/main.go

Note:

- Viết Hoa Function thì mới được gọi trong go
- Sử dụng zap để ghi log được thiết kế để sử dụng trong ứng dụng GO => có hiệu suất vượt trội với hiệu suất cao.
- install go: go get -u go.uber.org/zap

go run cmd/cli/main.log.go

## Tương tác File trong Go bằng Thư viện OS

## TEST: Phải chạy 2 lệnh này thì mới test được độ bao phủ của test case user đã viết

- go test -coverprofile=coverage.out
- go tool cover -html=coverage.out -o coverage.html (Check được độ bao phủ của test đi vào từng hàm đã đủ độ bao phủ hay chưa?)
