package router

import (
	"net/http"

	handler "github.com/Almazatun/gogrpc-mis/gateway_rrh/pkg/handler/prometheus"
)

func MiddlewareMetrics(handler http.Handler, m handler.Metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Process metrics
		m.UpgradeGateway()

		// Process request
		handler.ServeHTTP(w, r)
	})
}
