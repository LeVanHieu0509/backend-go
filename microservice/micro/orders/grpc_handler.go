package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
	"github.com/LeVanHieu0509/backend-go/microservice/micro/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type GrpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrdersService
	channel *amqp.Channel
}

// Tạo mới GRPC handler
func NewGRPCHandler(grpcHandler *grpc.Server, service OrdersService, channel *amqp.Channel) {
	handler := &GrpcHandler{
		service: service,
		channel: channel,
	}

	// Đảm bảo khai báo exchange một lần khi server khởi động
	err := handler.setupExchange()
	if err != nil {
		log.Fatalf("Failed to set up exchange: %v", err)
	}

	pb.RegisterOrderServiceServer(grpcHandler, handler)
}

// Khai báo Exchange chỉ một lần trong toàn bộ ứng dụng
func (h *GrpcHandler) setupExchange() error {
	// Exchange đã được khai báo ở đây, tránh việc khai báo lại mỗi lần gọi
	return h.channel.ExchangeDeclare(
		broker.OrderCreatedEvent, // Tên Exchange
		"direct",                 // Kiểu Exchange
		true,                     // Durable
		false,                    // Không auto-delete
		false,                    // Không internal
		false,                    // Không no-wait
		nil,                      // Tham số thêm
	)
}

func (h *GrpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New order received! Order %v", p)

	// Tạo đơn hàng mới (chỉ ví dụ)
	o := &pb.Order{
		ID:     "421123123", // Tạo ID đơn hàng giả
		Status: "Created",   // Trạng thái của đơn hàng mới
	}

	// Chuyển đối tượng đơn hàng thành JSON để gửi đi
	orderBytes, err := json.Marshal(o)
	if err != nil {
		log.Printf("Failed to marshal order: %v", err)
		return nil, err
	}

	// Gửi tin nhắn đến RabbitMQ (publish)
	err = h.publishOrderToRabbitMQ(orderBytes)

	if err != nil {
		log.Printf("Failed to publish order to RabbitMQ: %v", err)
		return nil, err
	}

	log.Printf("Order published to RabbitMQ: %v", o)

	// Trả về đơn hàng đã tạo
	return o, nil
}

// Gửi tin nhắn đến RabbitMQ
func (h *GrpcHandler) publishOrderToRabbitMQ(orderBytes []byte) error {
	return h.channel.Publish(
		broker.OrderCreatedEvent,      // Exchange, gửi tới Exchange đã khai báo
		broker.OrderCreatedRoutingKey, // Routing key
		false,                         // Mandatory
		false,                         // Immediate
		amqp.Publishing{
			ContentType:  "application/json", // Định dạng dữ liệu là JSON
			Body:         orderBytes,         // Thân tin nhắn (dữ liệu đơn hàng)
			DeliveryMode: amqp.Persistent,    // Đảm bảo tin nhắn không bị mất
		},
	)
}
