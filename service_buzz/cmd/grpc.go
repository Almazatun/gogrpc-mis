package main

import (
	"log"
	"net"

	handler "github.com/Almazatun/gogrpc-mis/service_buzz/pkg/handler"
	service "github.com/Almazatun/gogrpc-mis/service_buzz/pkg/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *GRPCServer {
	return &GRPCServer{addr: addr}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Printf("failed to listen: %v", err)

		return err
	}

	grpcServer := grpc.NewServer()

	// register our grpc services
	buzzService := service.NewBuzzService()
	handler.NewBuzzGrpcHandler(grpcServer, buzzService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
