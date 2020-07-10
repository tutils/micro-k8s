package monitor

import (
    "fmt"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

func StartMonitor() {
    http.Handle("/metrics", promhttp.Handler())
    if err := http.ListenAndServe(":9090", nil); err != nil {
        panic(fmt.Sprintf("ERROR: cannot init Prometheus%v\n", err))
    }
}
