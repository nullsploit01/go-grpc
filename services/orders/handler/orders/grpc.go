package handler

import (
	"context"

	"github.com/nullsploit01/go-microservice/kitchen/services/common/genproto/orders"
	"github.com/nullsploit01/go-microservice/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, orderService types.OrderService) {
	grpcHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}

	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	ordersList := h.orderService.GetOrders(ctx)

	res := &orders.GetOrderResponse{
		Orders: ordersList,
	}

	return res, nil
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    1,
		CustomerID: 2,
		ProductID:  3,
		Quantity:   69,
	}

	err := h.orderService.CreateOrder(ctx, order)

	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "Success!",
	}

	return res, nil
}
