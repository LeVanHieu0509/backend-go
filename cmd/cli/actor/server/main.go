package main

import (
	"fmt"
	"log"

	"github.com/LeVanHieu0509/backend-go/cmd/cli/actor/messages"
	console "github.com/asynkron/goconsole"    // Thư viện để đọc dòng lệnh từ console.
	"github.com/asynkron/protoactor-go/actor"  // Thư viện ProtoActor để tạo và quản lý actors
	"github.com/asynkron/protoactor-go/remote" // Thư viện ProtoActor để cấu hình và khởi động các actors từ xa.
)

// Hàm này gửi một thông điệp (message) đến tất cả các clients trong clients.
// Bối cảnh hiện tại của actor.
// Tập hợp các Process IDs (PIDs) của các clients.
// Thông điệp cần gửi.
func notifyAll(context actor.Context, clients *actor.PIDSet, message interface{}) {
	fmt.Println("Server send to client: ", message)
	for _, client := range clients.Values() {
		context.Send(client, message)
	}
}

func main() {
	// Tạo một hệ thống actor mới.
	system := actor.NewActorSystem()

	// Cấu hình cho việc kết nối từ xa với địa chỉ 127.0.0.1 và cổng 8080.
	config := remote.Configure("127.0.0.1", 8080)

	// Khởi tạo một đối tượng từ xa với hệ thống actor và cấu hình.
	remoter := remote.NewRemote(system, config)

	// Bắt đầu hệ thống từ xa.
	remoter.Start()

	// Tạo một tập hợp các PIDs của các clients.
	clients := actor.NewPIDSet()

	// Định nghĩa các thuộc tính của actor dựa trên hàm xử lý các thông điệp.
	props := actor.PropsFromFunc(func(context actor.Context) {

		// Kiểm tra loại thông điệp nhận được và xử lý tương ứng
		switch msg := context.Message().(type) {

		// Xử lý thông điệp kết nối từ client.
		case *messages.Connect:
			log.Printf("Client %v connected", msg.Sender)
			clients.Add(msg.Sender)
			context.Send(msg.Sender, &messages.Connected{Message: "Welcome!"})

		// Xử lý yêu cầu nói chuyện từ client.
		case *messages.SayRequest:
			notifyAll(context, clients, &messages.SayResponse{
				UserName: msg.UserName,
				Message:  msg.Message,
			})

		// Xử lý yêu cầu đổi tên từ client.
		case *messages.NickRequest:
			notifyAll(context, clients, &messages.NickResponse{
				OldUserName: msg.OldUserName,
				NewUserName: msg.NewUserName,
			})
		}
	})

	// Khởi tạo và chạy actor, Tạo và khởi động actor với tên chatserver.
	_, _ = system.Root.SpawnNamed(props, "chatserver")

	// Đọc một dòng từ console để giữ chương trình chạy.
	_, _ = console.ReadLine()
}

/*
Tổng kết
	1, Khởi tạo hệ thống và cấu hình: Tạo hệ thống actor và cấu hình kết nối từ xa.
	2, Khởi tạo tập hợp clients: Tạo tập hợp lưu trữ các clients kết nối.
	3, Định nghĩa thuộc tính actor: Xử lý các thông điệp khác nhau từ clients.
	4, Khởi tạo và chạy actor: Tạo và chạy actor chatserver, đồng thời giữ chương trình chạy với console.ReadLine().
*/
