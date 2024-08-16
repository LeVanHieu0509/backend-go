package main

import (
	"fmt"
	"time"

	"github.com/LeVanHieu0509/backend-go/plugins/cypto-exchange/client"
	"github.com/LeVanHieu0509/backend-go/plugins/cypto-exchange/server"
)

const ethPrice = 1281

func main() {
	go server.StartServer()
	time.Sleep(1 * time.Second)
	c := client.NewClient()

	go makeMarketSimple(c)

}

func seedMarket(c *client.Client) {
	currentPrice := ethPrice // async call to fetch the price

	ask := orderbook.NewOrder(true, 10, 8)
	resp, err := c.PlaceLimitOrder()
}

func makeMarketSimple(c *client.Client) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		bestAsk, err := c.GetBestAsk()
		if err != nil {
			panic(err)
		}

		bestBid, err := c.GetBestBid()
		if err != nil {
			panic(err)
		}

		if bestAsk == 0 && bestBid == 0 {
			seedMarket(c)
		}

		fmt.Println("Best Ask:", bestAsk)
		fmt.Println("Best Bid:", bestBid)

		<-ticker.C
	}
}
