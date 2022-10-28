package http

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
)

type Config struct {
	DefaultPrometheusMetric string
	AddressPort             int
	AddressIp               string
}

type Server struct {
	GinEngine *gin.Engine
}

func NewHttpServer(cfg Config, metrics metrics.Metrics) Server {
	r := gin.New()
	routes := Routes{gin: r, Metrics: metrics}
	routes.BindRoutes(cfg)

	return Server{
		GinEngine: r,
	}
}
