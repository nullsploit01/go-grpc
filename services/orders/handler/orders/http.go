package handler

import (
	"net/http"

	"github.com/nullsploit01/go-microservice/kitchen/services/common/genproto/orders"
	"github.com/nullsploit01/go-microservice/kitchen/services/common/util"
	"github.com/nullsploit01/go-microservice/kitchen/services/orders/types"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		orderService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest

	err := util.ParseJSON(r, &req)

	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerId(),
		ProductID:  req.GetProductId(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)

	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "Success!"}
	util.WriteJSON(w, http.StatusOK, res)
}
