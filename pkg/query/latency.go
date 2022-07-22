package query

import (
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	latencyCollector *prometheus.HistogramVec = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds_bucket",
	}, []string{"hostname", "ip"})

	latencyGoodEvents http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// request duration faster than 0.3s / 0.300 ms
		randLatency := rand.Float64() * 0.29
		latencyCollector.WithLabelValues(r.Host, r.RemoteAddr).Observe(randLatency)
	}

	latencyBadEvents http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// request duration slower than 0.3s / 0.300 ms
		randLatency := rand.Float64() + 0.30
		latencyCollector.WithLabelValues(r.Host, r.RemoteAddr).Observe(randLatency)
	}
)
