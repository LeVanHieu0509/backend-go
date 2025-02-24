package main

import (
	"encoding/json"
	"log"

	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
	"github.com/LeVanHieu0509/backend-go/microservice/micro/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Consumer listens for messages from RabbitMQ and processes them
type consumer struct {
}

// NewConsumer creates a new instance of consumer
func NewConsumer() *consumer {
	return &consumer{}
}

// Listen listens for incoming messages from the queue
func (c *consumer) Listen(ch *amqp.Channel) {
	log.Println("Consumer started listening for messages")

	// Declare the queue (if it does not exist already)
	q, err := c.declareQueue(ch)
	if err != nil {
		log.Printf("Failed to declare queue: %v", err)
		return
	}

	// Start consuming messages from the declared queue
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Printf("Error starting consumer: %v", err)
		return
	}

	// Use a goroutine to process messages concurrently
	go c.processMessages(msgs)

	// Keep the consumer running
	select {}
}

// declareQueue ensures the queue is declared
func (c *consumer) declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		broker.OrderCreatedEvent, // Queue name
		true,                     // Durable
		false,                    // Auto-delete
		false,                    // Exclusive
		false,                    // No-wait
		nil,                      // Arguments
	)
}

// processMessages processes each message received from the queue
func (c *consumer) processMessages(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)

		// Unmarshal the message into order
		o, err := c.unmarshalOrder(d.Body)
		if err != nil {
			c.handleMessageFailure(d, err)
			continue
		}

		// Log the order information
		log.Printf("Payment link created for Order ID: %s", o.ID)

		// Acknowledge the message after processing
		if err := d.Ack(false); err != nil {
			log.Printf("Failed to acknowledge message: %v", err)
		}
	}
}

// unmarshalOrder unmarshals the byte data into a pb.Order
func (c *consumer) unmarshalOrder(body []byte) (*pb.Order, error) {
	var o pb.Order
	if err := json.Unmarshal(body, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

// handleMessageFailure handles failed message processing (Nack the message)
func (c *consumer) handleMessageFailure(d amqp.Delivery, err error) {
	log.Printf("Failed to unmarshal order: %v", err)
	if nackErr := d.Nack(false, false); nackErr != nil {
		log.Printf("Failed to Nack message: %v", nackErr)
	}
}
