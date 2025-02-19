package user

import (
	"github.com/LeVanHieu0509/backend-go/internal/controller/ticket"
	"github.com/gin-gonic/gin"
)

type TicketRouter struct {
}

func (pr *TicketRouter) InitTicketRouter(Router *gin.RouterGroup) {
	ticketRouterPublic := Router.Group("/ticket")
	{
		// ticketRouterPublic.GET("/search")
		ticketRouterPublic.GET("/item/:id", ticket.TicketItem.GetTicketItemById)

	}
}
