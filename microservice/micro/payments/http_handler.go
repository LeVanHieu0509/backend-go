package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/LeVanHieu0509/backend-go/microservice/micro/common"
	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
	"github.com/LeVanHieu0509/backend-go/microservice/micro/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

// PaymentHTTPHandler handles HTTP requests for payment services
type PaymentHTTPHandler struct {
	channel *amqp.Channel
}

// NewPaymentHTTPHandler creates a new PaymentHTTPHandler
func NewPaymentHTTPHandler(channel *amqp.Channel) *PaymentHTTPHandler {
	return &PaymentHTTPHandler{channel}
}

// registerRoutes registers HTTP routes for the payment service
func (h *PaymentHTTPHandler) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("/webhook", h.handleCheckoutWebhook) // Fixed the leading space in route
}

// handleCheckoutWebhook handles the incoming webhook requests from checkout
func (h *PaymentHTTPHandler) handleCheckoutWebhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a webhook request")

	const MaxBodyBytes = int64(65536)
	// Use MaxBytesReader to limit the request body size
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Unable to read request body", http.StatusServiceUnavailable)
		return
	}

	// Log the request body (in real production, avoid printing sensitive data)
	log.Printf("Request body: %s", body)

	// TODO: Add signature verification here (currently omitted)

	// Create a sample order to simulate webhook processing
	o := &pb.Order{
		ID:     "421123123", // This ID should be dynamically fetched from the webhook payload
		Status: "paid",      // You may want to adjust the status based on the webhook
	}

	// Marshal the order object to JSON
	orderBytes, err := json.Marshal(o)
	if err != nil {
		log.Printf("Error marshaling order: %v", err)
		http.Error(w, "Failed to process order", http.StatusInternalServerError)
		return
	}

	// Publish the order to RabbitMQ
	if err := h.publishOrderPaidEvent(orderBytes); err != nil {
		log.Printf("Error publishing order event: %v", err)
		http.Error(w, "Failed to publish order event", http.StatusInternalServerError)
		return
	}

	// Respond with the order information
	common.WriteJSON(w, http.StatusOK, o)
}

// publishOrderPaidEvent publishes the order paid event to RabbitMQ
func (h *PaymentHTTPHandler) publishOrderPaidEvent(orderBytes []byte) error {
	return h.channel.Publish(
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
}
