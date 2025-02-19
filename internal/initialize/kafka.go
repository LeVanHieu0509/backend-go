package initialize

import (
	"context"
	"log"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/segmentio/kafka-go"
)

var (
	// Là một con trỏ tới một Kafka writer để gửi tin nhắn.
	kafkaProducer *kafka.Writer
	// kafkaConsumer *kafka.Reader
)

const (
	kafkaURL   = "192.168.0.109:9193" //Kafka broker URL
	kafkaTopic = "otp-auth-topic"     //Kafka topic name
)

func InitKafka() {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}

	// Test connection
	err := writer.WriteMessages(context.Background(), kafka.Message{Value: []byte("Test connection")})
	if err != nil {
		log.Fatalf("Failed to connect Kafka producer: %v", err)
	}

	log.Println("Kafka producer connected successfully.")
	global.KafkaProducer = writer
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatal("Failed to close kafka producer: %v", err)
	}
}
