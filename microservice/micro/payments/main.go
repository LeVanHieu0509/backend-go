package main

import (
	"log"
	"net"
	"net/http"

	"github.com/LeVanHieu0509/backend-go/microservice/micro/common"
	"github.com/LeVanHieu0509/backend-go/microservice/micro/common/broker"
	"google.golang.org/grpc"
)

var (
	amqpUser = common.EnvString("RABBITMQ_USER", "rabbitmq")
	amqpPass = common.EnvString("RABBITMQ_PASS", "rabbitmq")
	amqpHost = common.EnvString("RABBITMQ_HOST", "localhost")
	amqpPort = common.EnvString("RABBITMQ_PORT", "5672")
	grpcAddr = common.EnvString("GRPC_ADDRESS", "localhost:2001")
	httpAddr = common.EnvString("HTTP_ADDR", "localhost:8081")
)

func main() {

	// Broker connection
	ch, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func() {
		close()
		ch.Close()
	}()

	// Tạo một consumer cho payment service lắng nghe các tin nhắn từ order service
	consumer := NewConsumer()
	go consumer.Listen(ch)

	// http server
	mux := http.NewServeMux()

	httpServer := NewPaymentHTTPHandler(ch)
	httpServer.registerRoutes(mux)

	go func() {
		log.Printf("Starting HTTP server at %s", httpAddr)
		if err := http.ListenAndServe(httpAddr, mux); err != nil {
			log.Fatal("failed to start http server")
		}
	}()

	// gRPC server
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	log.Println("GRPC Server Started at ", grpcAddr)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
