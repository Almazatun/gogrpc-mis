package main

import (
	"os"
)

func main() {
	add := os.Getenv("ADD_SERVICE_BUZZ")
	grpcServer := NewGRPCServer(add)
	err := grpcServer.Run()

	if err != nil {
		panic(err)
	}
}
