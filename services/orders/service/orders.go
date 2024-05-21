package service

import (
	"context"

	"github.com/nullsploit01/go-microservice/kitchen/services/common/genproto/orders"
)

var ordersList = make([]*orders.Order, 0)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersList = append(ordersList, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersList
}
