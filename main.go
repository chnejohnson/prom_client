package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
)

func main() {
	flag.Parse()

	reg := prometheus.NewRegistry()

	if err := reg.Register(cpuTemp); err != nil {
		fmt.Println("cpu_temperature_celsius not registered:", err)
	} else {
		fmt.Println("cpu_temperature_celsius registered.")
	}

	cpuTemp.Set(65.3)

	http.Handle("/", http.HandlerFunc(handleTemp))

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	fmt.Printf("Server listening up on port %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleTemp(res http.ResponseWriter, req *http.Request) {
	cpuTemp.Add(10)
	res.Write([]byte("Hello World"))
}
