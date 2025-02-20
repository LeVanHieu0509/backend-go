package main

import (
	"context"

	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
)

type OrdersService interface {
	CreateOrder(context.Context) error
	ValidateOrder(context.Context, *pb.CreateOrderRequest) ([]*pb.Item, error)
	// GetOrder(context.Context, *pb.GetOrderRequest) (*pb.Order, error)
	// UpdateOrder(context.Context, *pb.Order) (*pb.Order, error)
}

type OrdersStore interface {
	Create(context.Context) error
	// Get(ctx context.Context, id, customerID string) (*Order, error)
	// Update(ctx context.Context, id string, o *pb.Order) error
}

type Order struct {
	// ID          primitive.ObjectID `bson:"_id,omitempty"`
	// CustomerID  string             `bson:"customerID,omitempty"`
	// Status      string             `bson:"status,omitempty"`
	// PaymentLink string             `bson:"paymentLink,omitempty"`
	// Items       []*pb.Item         `bson:"items,omitempty"`
}

// func (o *Order) ToProto() *pb.Order {
// 	return &pb.Order{
// 		ID:          o.ID.Hex(),
// 		CustomerID:  o.CustomerID,
// 		Status:      o.Status,
// 		PaymentLink: o.PaymentLink,
// 	}
// }
