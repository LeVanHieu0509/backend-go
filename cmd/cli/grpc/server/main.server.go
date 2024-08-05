package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	proto "github.com/LeVanHieu0509/backend-go/cmd/cli/grpc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The Port to connect")
)

type server struct {
	// Khai báo 1 server struct được kết thừa từ struct của proto để tạo ra 1 server protoc
	proto.UnimplementedOrderServiceServer
}

// tạo 1 function để xử lý yêu cầu đặt hàng từ client

/*
- ctx context.Context: Tham số ngữ cảnh, cung cấp thông tin về thời gian chờ, hủy bỏ, và các giá trị metadata khác
- in *proto.NewRequestOrder: Tham số in là con trỏ tới thông điệp NewRequestOrder mà client gửi đến
- (*proto.NewResponseOrder, error): Phương thức trả về con trỏ tới thông điệp NewResponseOrder và một lỗi (nếu có).
*/
func (s *server) NewOrder(ctx context.Context, in *proto.NewRequestOrder) (*proto.NewResponseOrder, error) {
	log.Printf("Receive order:::%v", in.GetOrderRequest())

	// Lấy giá trị của trường orderRequest từ thông điệp NewRequestOrder.

	// Tạo một phản hồi NewResponseOrder mới với trường orderResponse chứa một chuỗi kết hợp "new orderID"
	// và giá trị của orderRequest từ yêu cầu.
	return &proto.NewResponseOrder{OrderResponse: "new orderID" + in.GetOrderRequest()}, nil

}

// tạo ra 1 server mới và log ra
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatal(err)
	}

	// Giới hạn 4mb
	// s := grpc.NewServer() //Dang ki server GRPC moi
	// proto.RegisterOrderServiceServer(s, &server{})

	// Thiết lập cấu hình server gRPC 20mb
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(20 * 1024 * 1024), // 20 MB
		grpc.MaxSendMsgSize(20 * 1024 * 1024), // 20 MB
	}

	// Tạo một server gRPC với các tùy chọn cấu hình và đăng ký dịch vụ OrderService với server.
	s := grpc.NewServer(opts...)
	proto.RegisterOrderServiceServer(s, &server{})

	log.Printf("Server listening on port %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

/*
Tổng quan quy trình
Khi client gọi phương thức NewOrder, server sẽ nhận được yêu cầu này.
Server ghi nhật ký nội dung yêu cầu.
Server tạo một phản hồi NewResponseOrder mới, kết hợp một chuỗi "new orderID" với giá trị từ yêu cầu.
Server trả về phản hồi và không có lỗi.
*/
