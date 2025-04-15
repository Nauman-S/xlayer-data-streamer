package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var EffectiveTCPBytes = promauto.NewCounter(prometheus.CounterOpts{
	Name: "relay_client_effective_tcp_bytes_total",
	Help: "The final, ordered, de-duplicated TCP stream from server",
})
