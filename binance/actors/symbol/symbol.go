package symbol

import (
	"fmt"

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

// Receive implements actor.Actor.
func (s *Symbol) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		_ = msg
	}
}

func (a *Symbol) start() {
	fmt.Println("symbol actor started", a.Symbol)
}
