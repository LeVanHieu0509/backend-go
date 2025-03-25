package initialize

import "github.com/LeVanHieu0509/backend-go/internal/controller/ticket"

func InitPrometheus() {
	ticket.TicketItem.Init()
}
