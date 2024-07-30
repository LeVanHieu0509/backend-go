package main

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/binance/actors/consumer"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

func main() {
	system := actor.NewActorSystem()
	remote := remote.NewRemote(system, remote.Configure("127.0.0.1", 3000))
	remote.Start()
	pid, _ := system.Root.SpawnNamed(actor.PropsFromProducer(consumer.NewBinancef()), "binancef")
	fmt.Println(pid)
	select {}
}
