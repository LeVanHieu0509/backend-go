package symbol

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/binance/event"
	"github.com/asynkron/protoactor-go/actor"
)

type Symbol struct {
	Symbol string
}

func New(symbol string) actor.Producer {
	return func() actor.Actor {
		return &Symbol{
			Symbol: symbol,
		}
	}
}

//
// Create Alert => [cond_1,cond_2, cond_3] => price > x, Volume > 0
// => Postgres
// => Postgres => Alert => Redis

// Receive implements actor.Actor
// 1. When are we going to send the data in case of a WS stream
// AND when are we going to apply / exec the trigger?

func (s *Symbol) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *event.MarketPrice:
		fmt.Printf("Received MarketPrice: %s - %f\n", msg.Symbol, msg.Price)
	case *event.Trade:
		fmt.Printf("Received Trade: Symbol: %s - Price: %f - Qty: %f\n", msg.Symbol, msg.Price, msg.Qty)
	case *actor.Started:
		_ = msg
	}
}

func (a *Symbol) start() {
	fmt.Println("symbol actor started", a.Symbol)
}
