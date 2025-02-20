package main

import (
	"context"
	"log"

	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
	"google.golang.org/grpc"
)

type GrpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrdersService
}

func NewGRPCHandler(grpcHandler *grpc.Server, service OrdersService) {
	handler := &GrpcHandler{
		service: service,
	}

	pb.RegisterOrderServiceServer(grpcHandler, handler)
}

func (h *GrpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received! Order %v", p)

	o := &pb.Order{
		ID: "42",
	}

	return o, nil
}
