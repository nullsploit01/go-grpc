package handler

import (
	"github.com/nullsploit01/go-microservice/kitchen/services/common/genproto/orders"
	"github.com/nullsploit01/go-microservice/kitchen/services/orders/types"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService() {
	// grpcHandler := &OrdersGrpcHandler{}
}
