package ticket

import (
	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
	"github.com/gin-gonic/gin"
)

var TicketItem = new(cTicketItem)

type cTicketItem struct{}

func (p *cTicketItem) GetTicketItemById(ctx *gin.Context) {
	ticketItem, err := service.TicketItem().GetTicketItemById(ctx, 1)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	} else {
		response.SuccessResponse(ctx, response.ErrCodeSuccess, ticketItem)
	}
}
