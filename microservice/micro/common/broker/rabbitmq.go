package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

const MaxRetryCount = 3
const DLQ = "dlq_main"

// Hàm này thiết lập kết nối đến RabbitMQ và tạo các Exchange cần thiết cho ứng dụng
func Connect(user, pass, host, port string) (*amqp.Channel, func() error) {
	address := fmt.Sprintf("amqp://%s:%s@%s:%s", user, pass, host, port)

	conn, err := amqp.Dial(address)

	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	//  dùng để phân phối các tin nhắn cho các consumers khác nhau
	err = ch.ExchangeDeclare(OrderCreatedEvent, "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.ExchangeDeclare(OrderPaidEvent, "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// DLQ (Dead Letter Queue) cho các tin nhắn bị lỗi (sẽ được lưu trữ lại sau khi vượt qua số lần thử lại tối đa).
	err = createDLQAndDLX(ch)
	if err != nil {
		log.Fatal(err)
	}

	return ch, conn.Close
}

// Hàm này xử lý các tin nhắn bị lỗi và thử lại (retry)
func HandleRetry(ch *amqp.Channel, d *amqp.Delivery) error {
	if d.Headers == nil {
		d.Headers = amqp.Table{}
	}

	// Nếu x-retry-count không có trong header của tin nhắn, nó sẽ khởi tạo giá trị retry count là 0.
	retryCount, ok := d.Headers["x-retry-count"].(int64)
	if !ok {
		retryCount = 0
	}

	// Tin nhắn sẽ được gửi lại với số lần retry tăng dần
	retryCount++
	d.Headers["x-retry-count"] = retryCount

	log.Printf("Retrying message %s, retry count: %d", d.Body, retryCount)

	if retryCount >= MaxRetryCount {
		log.Printf("Moving message to DLQ %s", DLQ)

		return ch.PublishWithContext(context.Background(), "", DLQ, false, false, amqp.Publishing{
			ContentType:  "application/json",
			Headers:      d.Headers,
			Body:         d.Body,
			DeliveryMode: amqp.Persistent,
		})
	}

	// Sau mỗi lần retry, chương trình sẽ ngủ (time.Sleep) một khoảng thời gian
	// tùy theo số lần retry, nhằm giảm tải cho hệ thống.
	time.Sleep(time.Second * time.Duration(retryCount))

	// Nếu số lần retry đạt đến MaxRetryCount (mặc định là 3),
	// tin nhắn sẽ được chuyển sang Dead Letter Queue (DLQ) để xử lý sau hoặc kiểm tra lại.
	return ch.PublishWithContext(
		context.Background(),
		d.Exchange,
		d.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Headers:      d.Headers,
			Body:         d.Body,
			DeliveryMode: amqp.Persistent,
		},
	)
}

// Hàm này tạo một Dead Letter Exchange (DLX) và Dead Letter Queue (DLQ).
func createDLQAndDLX(ch *amqp.Channel) error {
	// main_queue (hàng đợi chính) được liên kết với DLX,
	// điều này có nghĩa là nếu có vấn đề với tin nhắn trong main_queue,
	// nó sẽ bị chuyển đến dlx_main.
	q, err := ch.QueueDeclare(
		"main_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	// Declare DLX
	// DLX là nơi các tin nhắn bị lỗi hoặc không thể được xử lý sẽ được chuyển tới.
	dlx := "dlx_main"
	err = ch.ExchangeDeclare(
		dlx,      // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}

	// Bind main queue to DLX
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		dlx,    // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Declare DLQ
	// Tin nhắn sau khi vượt quá số lần retry sẽ bị chuyển đến DLQ,
	// đây là nơi lưu trữ các tin nhắn lỗi để có thể xử lý lại hoặc phân tích sau.
	_, err = ch.QueueDeclare(
		DLQ,   // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	return err
}

// Đây là một cấu trúc dữ liệu (map) dùng để chứa các header của tin nhắn trong RabbitMQ.
type AmqpHeaderCarrier map[string]interface{}

func (a AmqpHeaderCarrier) Get(k string) string {
	value, ok := a[k]
	if !ok {
		return ""
	}

	return value.(string)
}

func (a AmqpHeaderCarrier) Set(k string, v string) {
	a[k] = v
}

func (a AmqpHeaderCarrier) Keys() []string {
	keys := make([]string, len(a))
	i := 0

	for k := range a {
		keys[i] = k
		i++
	}

	return keys
}

// Đồng thời, mã này cũng sử dụng OpenTelemetry để theo dõi và chèn thông tin header vào các tin nhắn RabbitMQ.

// Hàm InjectAMQPHeaders sử dụng OpenTelemetry để "inject" các thông tin liên quan
// đến ngữ cảnh theo dõi vào header của tin nhắn trước khi gửi qua RabbitMQ.
func InjectAMQPHeaders(ctx context.Context) map[string]interface{} {
	carrier := make(AmqpHeaderCarrier)
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	return carrier
}

// Hàm ExtractAMQPHeader giúp trích xuất thông tin ngữ cảnh từ header của tin nhắn
// và khôi phục lại ngữ cảnh OpenTelemetry từ đó, giúp theo dõi và giám sát quá trình xử lý tin nhắn.
func ExtractAMQPHeader(ctx context.Context, headers map[string]interface{}) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, AmqpHeaderCarrier(headers))
}

/*
- OpenTelemetry giúp giám sát và theo dõi các hệ thống phân tán,
- và việc tích hợp OpenTelemetry vào RabbitMQ sẽ giúp bạn theo dõi
- các hoạt động và xử lý tin nhắn trong hệ thống.
*/
