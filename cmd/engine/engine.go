package engine

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
	Server  http.HttpServer
	Metrics metrics.MetricsInterface
}

const defaultEnv = "etc/dev/env"
const httpServerAddress = "HTTP_SERVER_ADDRESS"
const httpServerPort = "HTTP_SERVER_PORT"
const defaultPrometheusUrl = "DEFAULT_PROMETHEUS_URL"

func NewEngine(server http.HttpServer, metrics metrics.MetricsInterface) Engine {
	return Engine{
		server,
		metrics,
	}
}

func (e *Engine) BuildEngine() http.Config {
	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/../../%s", pwd, defaultEnv)
	err := godotenv.Load(path)
	if err != nil {
		err := godotenv.Load(defaultEnv)
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port, _ := strconv.Atoi(os.Getenv(httpServerPort))
	var cfg = http.Config{
		DefaultPrometheusMetric: os.Getenv(defaultPrometheusUrl),
		AddressPort:             port,
		AddressIp:               os.Getenv(httpServerAddress),
	}
	e.Server.BuildHttpServer(e.Metrics)
	return cfg
}

func (e *Engine) RunEngine(port int) {
	e.Server.Gin().Run(fmt.Sprintf(": %d", port))
}
