package query

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	uptimeCollector *prometheus.GaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "uptime_good",
	}, []string{"hostname", "ip"})

	uptimeGoodEvents http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		uptimeCollector.WithLabelValues(r.Host, r.RemoteAddr).Set(1)
	}
	uptimeBadEvents http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		uptimeCollector.WithLabelValues(r.Host, r.RemoteAddr).Set(0)
	}
)
