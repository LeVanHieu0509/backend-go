package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "192.168.0.109:9193" // Kafka broker URL
	kafkaTopic = "user_topic_vip1"    // Kafka topic name
)

/*
Tạo Kafka reader (Consumer)
*/
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	log.Printf("Connecting to Kafka brokers: %v, topic: %s, groupID: %s", brokers, topic, groupID)

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,             // 10KB
		MaxBytes:       10e6,             // 10MB
		SessionTimeout: 60 * time.Second, // Tăng thời gian session timeout
		StartOffset:    kafka.LastOffset,
	})

}

/*
Tạo Kafka writer (Producer)
*/
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(kafkaURL),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne, //
	}

	// Test connection
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// err := writer.WriteMessages(ctx, kafka.Message{Value: []byte("Test connection")})
	// if err != nil {
	// 	log.Fatalf("Failed to connect Kafka producer: %v\n", err)
	// }

	// log.Println("Kafka producer connected successfully.")
	return writer
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	return &StockInfo{
		Message: msg,
		Type:    typeMsg,
	}
}

/*
API gửi tin nhắn tới Kafka
*/
func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))

	body := make(map[string]interface{}, 0)

	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	// Tạo ra 1 message trong kafka
	msg := kafka.Message{
		Key:   []byte(s.Message),
		Value: []byte(jsonBody),
	}

	// Gửi tin nhắn tới Kafka
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("Error writing to Kafka: %v\n", err)
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	log.Printf("Message sent successfully: %s\n", string(jsonBody))
	c.JSON(200, gin.H{
		"err": "",
		"msg": "Action successfully",
	})
}

// Consumer hóng mua ATC => Khớp lệnh => ATC: Khi mà phiên giao dịch hết
// - có 15p để xảy ra giá ATC, khi mà đóng lệnh thì giá ATC là giá bắt đầu cho ngày mai
// Đăng kí topic nào để lắng nghe

/*
Hàm này đăng ký một Kafka consumer để lắng nghe tin nhắn từ Kafka.
Consumer được cấu hình với một GroupID duy nhất và đọc tin nhắn từ Kafka topic.
Tin nhắn nhận được sẽ được in ra console.
*/
func RegisterConsumerATC(id int) {
	// group consumer???
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id)
	// Consumer đăng kí đọc ở topic nào và ở group ID nào
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	log.Printf("Consumer(%d) started listening for topic: %s\n", id, kafkaTopic)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Consumer(%d) error: %v. Retrying...\n", id, err)
			time.Sleep(2 * time.Second) // Retry sau 2 giây
			continue
		}

		log.Printf("Consumer(%d), topic: %s, partition: %d, offset: %d, key: %s, value: %s\n",
			id, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

/*
Chạy ứng dụng
*/
func main() {
	r := gin.Default()
	getIpHost()
	// Tạo Kafka producer
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	// Endpoint API
	r.POST("action/stock", actionStock)

	// Khởi chạy consumer
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	// Chạy HTTP server
	log.Println("Server running on http://localhost:8999")
	if err := r.Run(":8999"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getIpHost() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("Local IP Address:", ipNet.IP.String())
			}
		}
	}
}

// docker exec -it kafka1 kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic user_topic_vip1 --from-beginning
// tăng partition: docker exec -it kafka1 kafka-topics.sh --bootstrap-server localhost:9092 --alter --topic user_topic_vip1 --partitions 3
// kiểm tra partiion: docker exec -it kafka1 kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic user_topic_vip1
// thêm tin nhắn: docker exec -it kafka1 kafka-console-producer.sh --bootstrap-server localhost:9092 --topic user_topic_vip1
// kiểm tra phiên bản kafka: docker exec -it kafka1 kafka-topics.sh --version
