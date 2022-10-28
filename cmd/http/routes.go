package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application"
)

const DEFAULT_PING_URL = "/ping"

type Routes struct {
	gin *gin.Engine
}

func (g *Routes) BindRoutes(cfg Config, metrics metrics.Metrics) {
	g.gin.GET(DEFAULT_PING_URL, application.Ping(metrics))
	g.gin.GET("/"+cfg.DefaultPrometheusMetric, prometheusHandler())
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
