package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	// Là một con trỏ tới một Kafka writer để gửi tin nhắn.
	kafkaProducer *kafka.Writer
	// kafkaConsumer *kafka.Reader
)

const (
	kafkaURL   = "localhost:9092" //Kafka broker URL
	kafkaTopic = "user_topic_vip" //Kafka topic name
)

/*
For Consumer,
Hàm này tạo và trả về một Kafka reader để tiêu thụ tin nhắn từ Kafka.
Nó cấu hình Kafka reader với các tham số như brokers, GroupID, topic, MinBytes, MaxBytes và CommitInterval.
*/
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, // danh sách brockers => nếu chứa nhiều địa chỉ: []string{"localhost:2001", "localhost:200"}
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,        // 10KB - Vận chuyển dữ liệu
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // là khoảng thời gian giữa các lần commit offset là 1 s

		// Nếu user không muốn nhận tin nhắn đầu mà chỉ lấy tin nhắn cuối thôi thì sài LastOffset
		StartOffset: kafka.FirstOffset, // Các user vào lấy tin nhắn - đạt giá trị offset ban đầu khi mà user lắng nghe
	})
}

/*
For Producer,
Hàm này tạo và trả về một Kafka writer để gửi tin nhắn tới Kafka.
Writer được cấu hình với địa chỉ Kafka broker, tên topic và một balancer (LeastBytes) để cân bằng tải.
*/
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // Cân bằng tải
	}
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}

	s.Message = msg
	s.Type = typeMsg

	//return pointer s
	return &s
}

/*
Gửi tin nhắn này tới Kafka bằng Kafka producer.
Nếu gửi thành công, trả về JSON với thông báo thành công, nếu thất bại, trả về lỗi.
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

	// sử dụng producer để mà viết 1 message
	err := kafkaProducer.WriteMessages(context.Background(), msg)

	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
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

	// Giải phóng kết nối sau khi hàm chạy
	defer reader.Close()

	fmt.Printf("Consumer(%d) Hong Phien ATC::\n", id)

	for {
		m, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Printf("Consumer(%d) error: %v", id, err)
		}

		fmt.Printf("Consumer(%d), hong topic: %v, partition: %v, offset:%v, time:%d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

// đọc 1m request mỗi giây
func main() {
	r := gin.Default()

	// Khai báo producer và lắng nghe nó dựa trên pointer
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	// Đăng kí 2 user để mua stock trong phiên ATC bao gồm ID 1, 2
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	go RegisterConsumerATC(4) // Không lấy những tin nhắn cũ

	r.Run(":8999")

}
