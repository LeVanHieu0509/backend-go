package service

import (
	"context"

	"github.com/LeVanHieu0509/backend-go/internal/model"
)

type (
	ITicketHome interface{}
	ITicketItem interface {
		GetTicketItemById(ctx context.Context, ticketId int) (out *model.TicketItemsOutput, err error)
	}
)

var (
	localTicketItem ITicketItem
	localTicketHome ITicketHome
)

func TicketHome() ITicketHome {
	if localTicketHome == nil {
		panic("implement ...")
	}

	return localTicketHome
}

func InitTicketHome(i ITicketHome) {
	localTicketHome = i
}

func TicketItem() ITicketItem {
	if localTicketItem == nil {
		panic("implement ...")
	}

	return localTicketItem
}

func InitTicketItem(i ITicketItem) {
	localTicketItem = i
}
