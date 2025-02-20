package main

import (
	"log"
	"net/http"

	common "github.com/LeVanHieu0509/backend-go/microservice/micro/common"
	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serviceName      = "gateway"
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080") // Địa chỉ HTTP mà server sẽ lắng nghe
	orderServiceAddr = "localhost:2000"                       // Địa chỉ của dịch vụ gRPC mà bạn sẽ kết nối tới
)

func main() {
	// Dùng để thiết lập kết nối gRPC, với tùy chọn "insecure" (không sử dụng chứng chỉ bảo mật) trong môi trường phát triển.
	// Thiết lập kết nối đến dịch vụ gRPC tại orderServiceAddr
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()
	log.Println("Dialing orders service at ", orderServiceAddr)

	// Tạo client gRPC để gọi các phương thức của OrderService từ dịch vụ gRPC đã kết nối.
	c := pb.NewOrderServiceClient(conn)

	// Tạo một ServeMux, đây là router HTTP của Go, sẽ dùng để xử lý các route HTTP.
	mux := http.NewServeMux()

	// Tạo một handler với client gRPC đã kết nối.
	// NewHandler là một hàm mà bạn cần định nghĩa, có nhiệm vụ đăng ký các route HTTP và xử lý chúng
	handler := NewHandler(c)

	//  Đăng ký các route vào mux để HTTP server có thể xử lý.
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	// Bắt đầu lắng nghe và xử lý các yêu cầu HTTP đến địa chỉ httpAddr (mặc định là :8080)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
