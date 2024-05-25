package router

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/Almazatun/gogrpc-mis/gateway_rrh/pkg/handler/buzz"
)

type httpServer struct {
	port    string
	handler *handler.RoundRobinGrpcHandleListener
}

func NewHttpServer(port string, h *handler.RoundRobinGrpcHandleListener) *httpServer {
	return &httpServer{
		port:    port,
		handler: h,
	}
}

func (h *httpServer) Run() {
	router := http.NewServeMux()
	// Buzz
	router.HandleFunc("/buzz", h.handler.HandleRequests)

	fmt.Println("Starting server on " + h.port)

	// Listener
	go h.handler.Run()

	err := http.ListenAndServe(h.port, router)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
