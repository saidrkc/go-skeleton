package app

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"

	"go-skeleton/cmd/http"
	"go-skeleton/infrastructure/metrics"
)

type Engine struct {
	Server  http.Server
	Metrics metrics.Metrics
}

const defaultEnv = "etc/dev/env"
const httpServerAddress = "HTTP_SERVER_ADDRESS"
const httpServerPort = "HTTP_SERVER_PORT"
const defaultPrometheusUrl = "DEFAULT_PROMETHEUS_URL"

func NewEngine() Engine {
	err := godotenv.Load(defaultEnv)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, _ := strconv.Atoi(os.Getenv(httpServerPort))
	var cfg = http.Config{
		DefaultPrometheusMetric: os.Getenv(defaultPrometheusUrl),
		AddressPort:             port,
		AddressIp:               os.Getenv(httpServerAddress),
	}

	mtrcs := metricsRegister()
	srv := http.NewHttpServer(cfg, mtrcs)

	srv.GinEngine.Run(fmt.Sprintf(": %d", cfg.AddressPort))

	return Engine{
		Server:  srv,
		Metrics: mtrcs,
	}
}

func metricsRegister() metrics.Metrics {
	var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
	var totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path"},
	)
	var responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_status",
			Help: "Status of HTTP response",
		},
		[]string{"status"},
	)
	prometheus.MustRegister(httpDuration)
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(responseStatus)
	return metrics.NewMetrics(httpDuration, totalRequests, responseStatus)
}
