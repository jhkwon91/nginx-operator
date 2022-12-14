package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	ReconcilesTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "reconciles_total",
			Help: "Number of total reconcilation attempts",
		},
	)
)

func init() {
	metrics.Registry.Register(ReconcilesTotal)
}
