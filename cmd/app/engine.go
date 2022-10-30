package app

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"go-skeleton/cmd/http"
	"go-skeleton/infrastructure/metrics"
)

type EngineInterface interface {
	NewEngine() Engine
}

type Engine struct {
	Server  http.Server
	Metrics metrics.Metrics
}

const defaultEnv = "etc/dev/env"
const httpServerAddress = "HTTP_SERVER_ADDRESS"
const httpServerPort = "HTTP_SERVER_PORT"
const defaultPrometheusUrl = "DEFAULT_PROMETHEUS_URL"

func NewEngine(server http.Server, metrics metrics.Metrics) Engine {
	return Engine{
		server,
		metrics,
	}
}

func (e *Engine) RunEngine() {
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
	e.Server.BuildHttpServer(e.Metrics)
	e.Server.GinEngine.Run(fmt.Sprintf(": %d", cfg.AddressPort))
}
