package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/raymondwongso/gogox/stats"
	gogox_prom "github.com/raymondwongso/gogox/stats/prometheus"
)

var (
	a stats.Stats
)

func recordMetrics() {
	go func() {
		for {
			a.Add("gogox_add", 1.23, stats.Option{Tags: stats.Tags{"a": "b", "c": "d", "e": "f"}})
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	a = gogox_prom.New("", stats.Tags{"service": "gogox service"})

	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
