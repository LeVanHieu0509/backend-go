package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
	"github.com/LeVanHieu0509/backend-go/microservice/micro/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PaymentHTTPHandler struct {
	channel *amqp.Channel
}

func NewPaymentHTTPHandler(channel *amqp.Channel) *PaymentHTTPHandler {
	return &PaymentHTTPHandler{channel}
}

func (h *PaymentHTTPHandler) registerRoutes(router *http.ServeMux) {
	router.HandleFunc(" /webhook", h.handleCheckoutWebhook)
}

func (h *PaymentHTTPHandler) handleCheckoutWebhook(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleCheckoutWebhook")

	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)

	body, err := io.ReadAll(r.Body)

	fmt.Printf("body: %s", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}

	o := &pb.Order{
		ID:     "421123123",
		Status: "paid",
	}

	orderBytes, err := json.Marshal(o)

	h.channel.Publish(
		broker.OrderPaidEvent,      // Exchange, gửi tới Exchange đã khai báo
		broker.OrderPaidRoutingKey, // Routing key
		false,                      // Mandatory
		false,                      // Immediate
		amqp.Publishing{
			ContentType:  "application/json", // Định dạng dữ liệu là JSON
			Body:         orderBytes,         // Thân tin nhắn (dữ liệu đơn hàng)
			DeliveryMode: amqp.Persistent,    // Đảm bảo tin nhắn không bị mất
		},
	)

	w.WriteHeader(http.StatusOK)
}
