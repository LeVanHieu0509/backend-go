package main

import (
	"log"

	"github.com/LeVanHieu0509/backend-go/cmd/cli/actor/messages"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

func main() {
	// Tạo một hệ thống actor mới.
	system := actor.NewActorSystem()

	// Cấu hình cho việc kết nối từ xa với địa chỉ 127.0.0.1 và cổng ngẫu nhiên (0).
	config := remote.Configure("127.0.0.1", 0)

	// Khởi tạo một đối tượng từ xa với hệ thống actor và cấu hình.
	remoter := remote.NewRemote(system, config)

	// Bắt đầu hệ thống từ xa.
	remoter.Start()

	// Tạo một PID mới cho server chat với địa chỉ 127.0.0.1:8080 và tên chatserver.
	server := actor.NewPID("127.0.0.1:8080", "chatserver")

	// Định nghĩa bối cảnh gốc của hệ thống actor.
	rootContext := system.Root

	// Định nghĩa các thuộc tính của client actor dựa trên hàm xử lý các thông điệp.
	props := actor.PropsFromFunc(func(context actor.Context) {
		switch msg := context.Message().(type) {
		case *messages.Connected:
			log.Println(msg.Message)
		case *messages.SayResponse:
			log.Printf("%v: %v", msg.UserName, msg.Message)
		case *messages.NickResponse:
			log.Printf("%v is now known as %v", msg.OldUserName, msg.NewUserName)
		}
	})

	// Khởi tạo client actor và gửi yêu cầu kết nối đến server
	client := rootContext.Spawn(props)

	// Gửi yêu cầu kết nối đến server chat.
	rootContext.Send(server, &messages.Connect{
		Sender: client,
	})

	// Định nghĩa nickname
	nick := "Roger"

	// Khởi tạo console để đọc dòng lệnh từ người dùng và gửi yêu cầu nói chuyện đến server.
	cons := console.NewConsole(func(text string) {
		rootContext.Send(server, &messages.SayRequest{
			UserName: nick,
			Message:  text,
		})
	})

	// Định nghĩa lệnh /nick để thay đổi nickname.
	cons.Command("/nick", func(newNick string) {
		// Gửi yêu cầu đổi nickname đến server.
		rootContext.Send(server, &messages.NickRequest{
			OldUserName: nick,
			NewUserName: newNick,
		})
		nick = newNick
	})

	// Định nghĩa lệnh /nick để thay đổi nickname.
	cons.Command("/oderbook", func(newNick string) {
		// Gửi yêu cầu đổi nickname đến server.
		rootContext.Send(server, &messages.NickRequest{
			OldUserName: nick,
			NewUserName: newNick,
		})
		nick = newNick
	})
	cons.Run()
}

/*
Tổng kết:
	1, Khởi tạo hệ thống và cấu hình từ xa: Tạo hệ thống actor và cấu hình kết nối từ xa.
	2, Định nghĩa địa chỉ server chat: Tạo PID cho server chat.
	3, Định nghĩa root context: Định nghĩa bối cảnh gốc của hệ thống actor.
	4, Định nghĩa thuộc tính của client actor: Xử lý các thông điệp từ server.
	5, Khởi tạo client actor và gửi yêu cầu kết nối đến server: Tạo và khởi động client actor, gửi yêu cầu kết nối đến server chat.
	6, Định nghĩa nickname và khởi tạo console: Đặt tên nickname ban đầu và khởi tạo console để đọc dòng lệnh từ người dùng.
	7, Định nghĩa lệnh đổi nickname: Định nghĩa lệnh /nick để thay đổi nickname.
	8, Chạy console: Chạy console để bắt đầu đọc dòng lệnh từ người dùng.
*/
