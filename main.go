package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})

	reqTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "web_request_total",
		Help: "Total webapp request count",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(reqTotal)
}

func main() {

	cpuTemp.Set(65.3)

	http.Handle("/", http.HandlerFunc(requestCount))
	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Server listening up on port %s\n", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestCount(res http.ResponseWriter, req *http.Request) {
	reqTotal.Add(1)
	res.Write([]byte("Count added!"))

}
