package router

import (
	"fmt"
	"net/http"

	handlerB "github.com/Almazatun/gogrpc-mis/gateway/pkg/handler/buzz"
	handlerF "github.com/Almazatun/gogrpc-mis/gateway/pkg/handler/fuzz"
)

type httpServer struct {
	port string
	buzz handlerB.BuzzHttp
	fuzz handlerF.FuzzHttp
}

func NewHttpServer(port string, buzz handlerB.BuzzHttp, fuzz handlerF.FuzzHttp) *httpServer {
	return &httpServer{
		port: port,
		buzz: buzz,
		fuzz: fuzz,
	}
}

func (h *httpServer) Run() error {
	router := http.NewServeMux()

	// Buzz
	router.HandleFunc("/buzz", h.buzz.Ping)

	// Fuzz
	router.HandleFunc("/fuzz", h.fuzz.Ping)

	fmt.Println("Starting server on " + h.port)
	return http.ListenAndServe(h.port, router)
}
