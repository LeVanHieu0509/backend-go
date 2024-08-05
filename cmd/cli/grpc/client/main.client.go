package main

import (
	"context"
	"log"
	"time"

	proto "github.com/LeVanHieu0509/backend-go/cmd/cli/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
Bên client tạo ra 1 order thì đẩy về phía server để xử lý thông qua protoc
Nếu xử lý file buffer thì có vấn đề nếu file lớn hơn 10mb

Theo mặc định, gRPC giới hạn kích thước thông điệp là 4MB
*/
func main() {
	address := "localhost:9000"
	// max is 4mb
	// conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials())) //ko bảo mật

	//max is 20mb
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(20*1024*1024), grpc.MaxCallRecvMsgSize(20*1024*1024)), // 20 MB
	)

	if err != nil {
		log.Fatal("did not connect: %v", err)
	}

	defer conn.Close()

	c := proto.NewOrderServiceClient(conn)

	ticket := time.NewTicker(2 * time.Second) //Trong trường hợp này, ticket được khởi tạo với khoảng thời gian

	defer ticket.Stop()

	for range ticket.C { // là một cách tiện lợi để lặp vô hạn và chờ đợi tín hiệu từ trong ticket của go
		orderId := "123"                                                        //make([]byte, 10*1024*1024)                                   // make([]byte, 10*1024*1024)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Tăng thời gian timeout lên 5 giây

		r, err := c.NewOrder(ctx, &proto.NewRequestOrder{OrderRequest: string(orderId)})

		if err != nil {
			log.Fatalf("could not greet %v", err)
		}

		log.Printf("Client order: %s", r.GetOrderResponse())

		cancel()
	}
}
