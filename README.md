# backend-go

## LIBRARY --------------------------------------------------------------------------------------------------------------------

1. viper: go get github.com/spf13/viper
2. gin: go get -u github.com/gin-gonic/gin
3. logger: go get -u go.uber.org/zap
4. logger advance: go get github.com/natefinch/lumberjack
5. gorm: go get -u gorm.io/gorm
6. uuid: go get -u github.com/google/uuid
7. redis: go get -u github.com/redis/go-redis/v9
8. kafka: go get -u github.com/segmentio/kafka-go
9. grpc: go get -u google.golang.org/protobuf/reflect/protoreflect
10. grpc: go get -u google.golang.org/protobuf/runtime/protoimpl
11. go get -u google.golang.org/grpc
12. go get -u github.com/google/wire/cmd/wire
13. go get -u gorm.io/gen
14. go get -u github.com/go-sql-driver/mysql
15. go install github.com/pressly/goose/v3/cmd/goose@latest
16. go install github.com/swaggo/swag/cmd/swag@latest
17. go get -u github.com/swaggo/swag
18. go get -u github.com/golang-jwt/jwt

## Go (3): GIN vs ROUTER --------------------------------------------------------------------------------------------------------------------

run server: go run cmd/server/main.go

Note:

- Viết Hoa Function thì mới được gọi trong go
- Sử dụng zap để ghi log được thiết kế để sử dụng trong ứng dụng GO => có hiệu suất vượt trội với hiệu suất cao.
- Trong Go, chỉ những hàm, phương thức, hoặc biến bắt đầu bằng chữ cái viết hoa mới có thể truy cập từ bên ngoài package.

go run cmd/cli/main.log.go

## Tương tác File trong Go bằng Thư viện OS

## TEST: --------------------------------------------------------------------------------------------------------------------

- Phải chạy 2 lệnh này thì mới test được độ bao phủ của test case user đã viết

- go test -coverprofile=coverage.out
- go tool cover -html=coverage.out -o coverage.html (Check được độ bao phủ của test đi vào từng hàm đã đủ độ bao phủ hay chưa?)

# BLOCKCHAIN --------------------------------------------------------------------------------------------------------------------

PART 4

- go run cmd/blockchain/index.go createblockchain -address "HIEU"
- go run cmd/blockchain/index.go printchain
- go run cmd/blockchain/index.go getbalance -address "HIEU"
- go run cmd/blockchain/index.go send -from "HIEU" -to "HIEP" - amount 50

PART 5

- go run cmd/blockchain/index.go createwallet

## Cli --------------------------------------------------------------------------------------------------------------------

1. go list -m all
2. go list -m -versions

### Proto --------------------------------------------------------------------------------------------------------------------

1. go get -u google.golang.org/protobuf/reflect/protoreflect
1. go get -u google.golang.org/protobuf/runtime/protoimpl

1. go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
1. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Thêm vào tệp cấu hình shell (chỉ cần làm một lần)

1. echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
2. source ~/.bashrc # hoặc ~/.zshrc, ~/.profile tùy thuộc vào shell bạn sử dụng

# Kiểm tra cài đặt

1. which protoc-gen-go
2. which protoc-gen-go-grpc

# Tạo mã Go từ tệp .proto

Lệnh này sẽ tạo ra hai tệp Go: một cho định nghĩa giao thức và một cho dịch vụ gRPC.

1. --go_out=.: Chỉ định đầu ra cho mã Go từ Protocol Buffers.
2. --go-grpc_out=.: Chỉ định đầu ra cho mã gRPC từ Protocol Buffers.
3. protoc --go_out=. --go-grpc_out=. event.proto

### docker

1. docker build . -t go-backend-api
2. docker run -p 8001:8080 go-backend-api

3. rm ~/.docker/config.json
4. docker compose-up -d
5. docker build . -t crm.shopdev.com
6. docker network create my_network_pro
7. docker network connect bridge mysql_con
8. docker run --link mysql_con:mysql_con -p 8003:8001 crm.shopdev.com

### SQLC

- Sau khi tạo bảng database -> dùng sqlc để combine ra goose

1. Tạo ra file sqlc trong schema: make create_migration name=0001_pre_go_acc_user_verify_9999
2. Migration từ code qua db: make up_by_one
3. Sau khi tạo xong DB rồi thì sẽ viết cho từng func để và định nghĩa nó trong file querie.
4. Sqlc sẽ combind ra code của goose: make sqlgen
