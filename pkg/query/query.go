package query

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var AvailableQueries = map[string]Query{}

func init() {
	AvailableQueries["uptime"] = &MockMetricApp{
		Collector:          uptimeCollector,
		GoodEventGenerator: uptimeGoodEvents,
		BadEventGenerator:  uptimeBadEvents,
	}

	AvailableQueries["http-availability"] = &MockMetricApp{
		Collector:          availabilityCollector,
		GoodEventGenerator: availabilityGoodEvents,
		BadEventGenerator:  availabilityBadEvents,
	}

	AvailableQueries["http-latency"] = &MockMetricApp{
		Collector:          latencyCollector,
		GoodEventGenerator: latencyGoodEvents,
		BadEventGenerator:  latencyBadEvents,
	}
}

type Query interface {
	GetCollector() prometheus.Collector
	GetGoodEventGenerator() http.HandlerFunc
	GetBadEventGenerator() http.HandlerFunc
}

type MockMetricApp struct {
	prometheus.Collector
	GoodEventGenerator http.HandlerFunc
	BadEventGenerator  http.HandlerFunc
}

func (m MockMetricApp) GetCollector() prometheus.Collector {
	return m.Collector
}

func (m MockMetricApp) GetGoodEventGenerator() http.HandlerFunc {
	return m.GoodEventGenerator
}

func (m MockMetricApp) GetBadEventGenerator() http.HandlerFunc {
	return m.BadEventGenerator
}
