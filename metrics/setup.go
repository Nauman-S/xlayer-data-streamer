package metrics

import (
	"github.com/0xPolygonHermez/zkevm-data-streamer/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func Setup(address string) {
	if address == "" {
		log.Fatalf("Metrics enabled but address is empty")
	}
	log.Infof("Starting Prometheus server on %s", address)

	http.Handle("/debug/metrics/prometheus", promhttp.Handler())

	go func() {
		err := http.ListenAndServe(address, nil)
		if err != nil {
			log.Errorf("Error starting Prometheus server: %v", err)
			return
		}
	}()

	log.Infof("Prometheus server listening on %s", address)
}
