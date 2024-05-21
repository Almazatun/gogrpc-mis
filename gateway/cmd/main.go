package main

import (
	"log"
	"os"

	grpcServer "github.com/Almazatun/gogrpc-mis/gateway/pkg/grpc"
	buzzHandler "github.com/Almazatun/gogrpc-mis/gateway/pkg/handler/buzz"
	fuzzHandler "github.com/Almazatun/gogrpc-mis/gateway/pkg/handler/fuzz"
	router "github.com/Almazatun/gogrpc-mis/gateway/pkg/http"
)

func main() {
	// grpc
	conn_buzz := grpcServer.NewGRPCClient(os.Getenv("BUZZ_SERVICE_ADDR"))
	defer conn_buzz.Close()

	conn_fuzz := grpcServer.NewGRPCClient(os.Getenv("FUZZ_SERVICE_ADDR"))
	defer conn_fuzz.Close()

	// grpc handler
	buzzGRPCHandler := buzzHandler.NewBuzzGrpcHandler(conn_buzz)
	fuzzGRPCHandler := fuzzHandler.NewFuzzGrpcHandler(conn_fuzz)

	// buzz
	buzzHandler := buzzHandler.NewBuzzHttpHandler(buzzGRPCHandler)
	// fuzz
	fuzzHandler := fuzzHandler.NewFuzzHttpHandler(fuzzGRPCHandler)
	// http
	server := router.NewHttpServer(":3000", buzzHandler, fuzzHandler)
	err := server.Run()

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
