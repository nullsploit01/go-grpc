package main

import (
	"context"
	"log"
	"net/http"
	"text/template"
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

		ordersList, err := c.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerId: 22,
		})

		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, ordersList.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}

	})

	log.Println("server staring on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<title>Kitchen Orders</title>
	</head>
	<body>
		<h1>Orders List</h1>
		<table border="1">
			<tr>
				<th>Order ID</th>
				<th>Customer ID</th>
				<th>Quantity</th>
			</tr>
			{{range .}}
			<tr>
				<td>{{.OrderID}}</td>
				<td>{{.CustomerID}}</td>
				<td>{{.Quantity}}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>`
