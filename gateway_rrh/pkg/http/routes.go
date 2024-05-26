package router

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/Almazatun/gogrpc-mis/gateway_rrh/pkg/handler/buzz"
	handler_metrics "github.com/Almazatun/gogrpc-mis/gateway_rrh/pkg/handler/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	// Metrics
	regMetrics := prometheus.NewRegistry()
	metrics := handler_metrics.NewMetrics(regMetrics)

	// Buzz
	buzzHandler := http.HandlerFunc(h.handler.HandleRequests)
	router.Handle("/buzz", MiddlewareMetrics(buzzHandler, metrics))

	promHandler := promhttp.HandlerFor(regMetrics, promhttp.HandlerOpts{})
	// lock point
	router.Handle("/metrics", promHandler)

	fmt.Println("Starting server on " + h.port)

	// Listener
	go h.handler.Run()

	err := http.ListenAndServe(h.port, router)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
