package main

import (
	"context"
	"log"

	"github.com/LeVanHieu0509/backend-go/microservice/micro/common"
	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
)

type Service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *Service {
	return &Service{store}
}

func (s *Service) CreateOrder(context.Context) error {
	return nil
}

func (s *Service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Item, error) {
	if len(p.Items) == 0 {
		return nil, common.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(p.Items)

	// validate with the stock service
	// inStock, items, err := s.gateway.CheckIfItemIsInStock(ctx, p.CustomerID, mergedItems)
	// if err != nil {
	// 	return nil, err
	// }
	// if !inStock {
	// 	return items, common.ErrNoStock
	// }

	log.Printf("mergedItems: ", mergedItems)

	return nil, nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	merged := make([]*pb.ItemsWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}
