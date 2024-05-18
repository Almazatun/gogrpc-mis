package router

import (
	"fmt"
	"net/http"

	handler "github.com/Almazatun/gogrpc-mis/gateway/pkg/handler/buzz"
)

type httpServer struct {
	port string
	buzz handler.BuzzHttp
}

func NewHttpServer(port string, buzz handler.BuzzHttp) *httpServer {
	return &httpServer{
		port: port,
		buzz: buzz,
	}
}

func (h *httpServer) Run() error {
	router := http.NewServeMux()

	// Buzz
	router.HandleFunc("/buzz", h.buzz.Ping)

	fmt.Println("Starting server on " + h.port)
	return http.ListenAndServe(h.port, router)
}
