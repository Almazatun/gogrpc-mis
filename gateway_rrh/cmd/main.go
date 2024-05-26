package main

import (
	"os"
	"strings"

	grpcServer "github.com/Almazatun/gogrpc-mis/gateway_rrh/pkg/grpc"
	handler "github.com/Almazatun/gogrpc-mis/gateway_rrh/pkg/handler/buzz"
	router "github.com/Almazatun/gogrpc-mis/gateway_rrh/pkg/http"
	"google.golang.org/grpc"
)

func main() {
	// grpc
	listConnGrpcServices := []*grpc.ClientConn{}
	addrs := strings.Split(os.Getenv("ADD_SERVICES"), ",")

	for _, addr := range addrs {
		conn := grpcServer.NewGRPCClient(addr)
		listConnGrpcServices = append(listConnGrpcServices, conn)
		defer conn.Close()
	}

	// list grpc handler
	listHandler := []handler.BuzzGrpc{}
	for _, conn := range listConnGrpcServices {
		buzzGRPCHandler := handler.NewBuzzGrpcHandler(conn)
		listHandler = append(listHandler, buzzGRPCHandler)
	}

	rrh := handler.NewRoundRobinGrpcHandler(listHandler)
	server := router.NewHttpServer(":3055", rrh)
	server.Run()
}
