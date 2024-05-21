package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/nullsploit01/go-microservice/kitchen/services/common/genproto/orders"
	"github.com/nullsploit01/go-microservice/kitchen/services/common/util"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":3000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerId: 2,
			ProductId:  22,
			Quantity:   2,
		})

		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}
	})

	log.Println("server staring on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
