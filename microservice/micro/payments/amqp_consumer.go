package main

import (
	"encoding/json"
	"log"

	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
	"github.com/LeVanHieu0509/backend-go/microservice/micro/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	service PaymentsService
}

func NewConsumer() *consumer {
	return &consumer{}
}

func (c *consumer) Listen(ch *amqp.Channel) {
	log.Println("Consumer started listening for messages")

	// Declare the queue to listen for messages
	q, err := ch.QueueDeclare(
		broker.OrderCreatedEvent, // Queue name
		true,                     // Durable
		false,                    // Auto-delete
		false,                    // Exclusive
		false,                    // No-wait
		nil,                      // Arguments
	)
	if err != nil {
		log.Fatalf("Error declaring queue: %v", err)
	}

	// Start consuming messages from the declared queue
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer tag
		false,  // Auto-ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("Error starting consumer: %v", err)
	}

	// Use a goroutine to process messages concurrently
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// Unmarshal the message
			o := &pb.Order{}
			if err := json.Unmarshal(d.Body, o); err != nil {
				d.Nack(false, false)
				log.Printf("Failed to unmarshal order: %v", err)
				continue
			}

			// Log the order information
			log.Printf("Payment link created for Order ID: %s", o.ID)

			// Acknowledge the message
			d.Ack(false)
		}
	}()

	// Keep the consumer running
	select {}
}
