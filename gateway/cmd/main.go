package main

import (
	grpcServer "github.com/Almazatun/gogrpc-mis/gateway/pkg/grpc"
	buzzHandler "github.com/Almazatun/gogrpc-mis/gateway/pkg/handler/buzz"
	router "github.com/Almazatun/gogrpc-mis/gateway/pkg/http"
)

func main() {
	// grpc
	conn := grpcServer.NewGRPCClient(":5001")
	defer conn.Close()

	// grpc handler
	buzzGRPCHandler := buzzHandler.NewBuzzGrpcHandler(conn)
	// buzz
	buzzHandler := buzzHandler.NewBuzzHttpHandler(buzzGRPCHandler)
	// http
	server := router.NewHttpServer(":3000", buzzHandler)
	server.Run()
}
