package query

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	availabilityCollector *prometheus.CounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_duration_seconds_count",
	},
		[]string{"code"},
	)
	availabilityGoodEvents http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		randomStatusInt1 := rand.Intn(199)

		// anything between 200-399, and yes http status codes dont'work like this
		randomStatusCode := 200 + randomStatusInt1

		availabilityCollector.WithLabelValues(fmt.Sprintf("%d", randomStatusCode)).Inc()
	}

	availabilityBadEvents http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		randomStatusInt1 := rand.Intn(199)
		// anything between 400-599, and yes http status codes dont'work like this
		randomStatusCode := 400 + randomStatusInt1
		availabilityCollector.WithLabelValues(fmt.Sprintf("%d", randomStatusCode)).Inc()
	}
)
