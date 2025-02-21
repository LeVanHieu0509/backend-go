package main

import (
	"encoding/json"
	"log"

	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
	"github.com/LeVanHieu0509/backend-go/microservice/micro/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Đây là struct chính của consumer, nó chứa một trường service kiểu OrdersService.
// Đây là dịch vụ sẽ được sử dụng để cập nhật đơn hàng khi một sự kiện "OrderPaid" được nhận.
type Consumer struct {
}

// Hàm này khởi tạo một consumer mới với dịch vụ OrdersService
func NewConsumer() *Consumer {
	return &Consumer{}
}

// Dây là hàm chính lắng nghe các tin nhắn từ RabbitMQ
func (c *Consumer) Listen(ch *amqp.Channel) {
	log.Println("Consumer started listening for messages")
	//  Một hàng đợi tạm thời ("" tức là hàng đợi tự động tạo ra tên ngẫu nhiên) được khai báo với các tham số
	q, err := ch.QueueDeclare(broker.OrderPaidRoutingKey, true, false, true, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Consumer này sẽ lắng nghe sự kiện OrderPaidEvent từ một Exchange được định nghĩa sẵn (trong broker.OrderPaidEvent).
	err = ch.QueueBind(q.Name, "", broker.OrderPaidEvent, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Sau khi binding thành công, consumer sẽ bắt đầu nhận các tin nhắn từ hàng đợi thông qua phương thức Consume.
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	// Consumer xử lý tin nhắn trong một goroutine mới để không chặn chương trình chính
	go func() {
		for d := range msgs {
			log.Printf("Received message: %s", d.Body)

			// Tin nhắn nhận được (d.Body) sẽ được giải mã thành một đối tượng Order của gRPC.
			o := &pb.Order{}
			if err := json.Unmarshal(d.Body, o); err != nil {
				d.Nack(false, false)
				log.Printf("failed to unmarshal order: %v", err)
				continue
			}
			log.Printf("UpdateOrder Service Success After payment")

			// Sau khi xử lý tin nhắn thành công, consumer sẽ gọi Ack để xác nhận rằng tin nhắn đã được xử lý thành công và có thể loại bỏ khỏi hàng đợi.
			d.Ack(false)
		}
	}()

	<-forever
}
