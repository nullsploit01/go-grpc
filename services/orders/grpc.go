package main

import (
	"fmt"
	"log"
	"net"

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

	log.Println("Starting gRPC Server on %w", s.addr)

	return grpcServer.Serve(lis)
}
