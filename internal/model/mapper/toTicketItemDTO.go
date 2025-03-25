package mapper

import (
	"github.com/LeVanHieu0509/backend-go/internal/database"
	"github.com/LeVanHieu0509/backend-go/internal/model"
)

func ToTicketItemDTO(ticketItem database.GetTicketItemByIdRow) model.TicketItemsOutput {
	return model.TicketItemsOutput{
		TicketId:       int(ticketItem.ID),
		TicketName:     ticketItem.Name,
		StockInitial:   int(ticketItem.StockInitial),
		StockAvailable: int(ticketItem.StockAvailable),
	}
}
