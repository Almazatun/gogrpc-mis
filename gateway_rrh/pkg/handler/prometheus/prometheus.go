package handler

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	UpgradeGateway()
}

type MetricsProm struct {
	upgrades *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer) *MetricsProm {
	m := &MetricsProm{
		upgrades: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "gateway",
			Name:      "upgrades_total",
			Help:      "Counter of upgrades gateway_rrh",
		}, []string{"type"}),
	}

	reg.MustRegister(m.upgrades)

	return m
}

func (m *MetricsProm) UpgradeGateway() {
	// Increment metric
	m.upgrades.With(prometheus.Labels{"type": "router"}).Inc()
}
