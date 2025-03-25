package ticket

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "req_get_ticket_item_by_id",
		Help: "Number of pings received.",
	},
)

func (p *cTicketItem) Init() {
	prometheus.MustRegister(pingCounter)
}

var TicketItem = new(cTicketItem)

type cTicketItem struct{}

func (p *cTicketItem) GetTicketItemById(ctx *gin.Context) {
	pingCounter.Inc()

	ticket_item := ctx.Param("id")
	// Convert the string parameter to an integer.
	idInt, err := strconv.Atoi(ticket_item)
	if err != nil {
		// Handle the conversion error.  This is crucial!
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket item ID format"})
		return
	}
	// call implementation
	ticketItem, err := service.TicketItem().GetTicketItemById(ctx, idInt)
	if err != nil {
		if errors.Is(err, response.CouldNotGetTicketErr) {
			fmt.Println("4004???")
		}

		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())

	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, ticketItem)
}
