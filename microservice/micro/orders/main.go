package main

import (
	"context"
	"log"
	"net"

	"github.com/LeVanHieu0509/backend-go/microservice/micro/common"
	"github.com/LeVanHieu0509/backend-go/microservice/micro/common/broker"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
	amqpUser = common.EnvString("RABBITMQ_USER", "rabbitmq")
	amqpPass = common.EnvString("RABBITMQ_PASS", "rabbitmq")
	amqpHost = common.EnvString("RABBITMQ_HOST", "localhost")
	amqpPort = common.EnvString("RABBITMQ_PORT", "5672")
)

func main() {
	ch, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func() {
		close()
		ch.Close()
	}()

	consumer := NewConsumer()
	go consumer.Listen(ch)

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer, svc, ch)

	svc.CreateOrder(context.Background())
	log.Println("GRPC Server started at", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
