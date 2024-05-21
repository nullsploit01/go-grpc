package main

import (
	"fmt"
	"log"
	"net"

	handler "github.com/nullsploit01/go-microservice/kitchen/services/orders/handler/orders"
	"github.com/nullsploit01/go-microservice/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		return fmt.Errorf("err starting on grpc server: %w", err)
	}

	grpcServer := grpc.NewServer()

	orderServcice := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderServcice)

	log.Println("Starting gRPC Server on %w", s.addr)

	return grpcServer.Serve(lis)
}
