package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/alexandreLamarre/metric-mock-server/pkg/query"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartInstrumentationServer(ctx context.Context) (int, chan bool) {
	// lg := e.Logger
	port := 8080
	mux := http.NewServeMux()

	for queryName, queryObj := range query.AvailableQueries {
		// register each prometheus collector
		prometheus.MustRegister(queryObj.GetCollector())

		// create an endpoint simulating good events
		mux.HandleFunc(fmt.Sprintf("/%s/%s", queryName, "good"), queryObj.GetGoodEventGenerator())
		// create an endpoint simulating bad events
		mux.HandleFunc(fmt.Sprintf("/%s/%s", queryName, "bad"), queryObj.GetBadEventGenerator())

	}
	// expose prometheus metrics
	mux.Handle("/metrics", promhttp.Handler())

	autoInstrumentationServer := &http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%d", port),
		Handler:        mux,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	done := make(chan bool)

	go func() {
		err := autoInstrumentationServer.ListenAndServe()
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()
	go func() {
		defer autoInstrumentationServer.Shutdown(context.Background())
		select {
		case <-ctx.Done():
		case <-done:
		}
	}()

	return port, done
}

func main() {
	port, _ := StartInstrumentationServer(context.Background())
	fmt.Println("Instrumentation server started on port", port)
	select {}
}
