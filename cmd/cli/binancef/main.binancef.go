package main

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/binance/actors/consumer"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

func main() {
	// tạo một hệ thống Actor mới. Đây là điểm khởi đầu của hệ thống, nơi tất cả các Actor sẽ được quản lý.
	system := actor.NewActorSystem()

	// cấu hình một remote để kết nối hệ thống Actor với địa chỉ IP 127.0.0.1 (localhost) và cổng 3000.
	// Đây là cấu hình để giao tiếp giữa các hệ thống Actor qua mạng.
	remote := remote.NewRemote(system, remote.Configure("127.0.0.1", 3000))
	remote.Start()

	// Khởi Tạo Actor Binancef: tạo một Actor mới với tên "binancef" từ producer
	pid, _ := system.Root.SpawnNamed(actor.PropsFromProducer(consumer.NewBinancef()), "binancef")

	//  in ra PID (Process ID) của actor vừa được tạo. PID là một đối tượng đại diện cho actor và được sử dụng để gửi thông điệp đến actor đó.
	fmt.Println("pid", pid)
	select {}
}
