package impl

import (
	"context"
	"fmt"

	"github.com/LeVanHieu0509/backend-go/internal/database"
	"github.com/LeVanHieu0509/backend-go/internal/model"
)

type sTicketItem struct {
	// implementation interface here
	r *database.Queries
}

func NewTicketItemImpl(r *database.Queries) *sTicketItem {
	return &sTicketItem{
		r: r,
	}
}

func (s *sTicketItem) GetTicketItemById(ctx context.Context, ticketId int) (out *model.TicketItemsOutput, err error) {
	fmt.Println("CALL SERVICE GetTicketItemById")

	ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketId))
	if err != nil {
		return out, err
	}

	//mapper

	return &model.TicketItemsOutput{
		TicketId:       int(ticketItem.ID),
		TicketName:     ticketItem.Name,
		StockAvailable: int(ticketItem.StockAvailable),
		StockInitial:   int(ticketItem.StockInitial),
	}, nil
}
