package main

import (
	"context"

	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
)

type PaymentsService interface {
	CreatePayment(context.Context, *pb.Order) (string, error)
}
