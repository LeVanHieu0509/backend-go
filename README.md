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

echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc # hoặc ~/.zshrc, ~/.profile tùy thuộc vào shell bạn sử dụng

# Kiểm tra cài đặt

which protoc-gen-go
which protoc-gen-go-grpc

# Tạo mã Go từ tệp .proto

Lệnh này sẽ tạo ra hai tệp Go: một cho định nghĩa giao thức và một cho dịch vụ gRPC.

--go_out=.: Chỉ định đầu ra cho mã Go từ Protocol Buffers.
--go-grpc_out=.: Chỉ định đầu ra cho mã gRPC từ Protocol Buffers.

protoc --go_out=. --go-grpc_out=. event.proto
