package initialize

import (
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
	kafkaURL   = "localhost:9092" //Kafka broker URL
	kafkaTopic = "otp-auth-topic" //Kafka topic name
)

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{}, // Cân bằng tải
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatal("Failed to close kafka producer: %v", err)
	}
}
