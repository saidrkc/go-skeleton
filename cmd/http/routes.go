package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application"
	"go-skeleton/src/domain"
	"go-skeleton/src/infrastructure/http/get"
)

const DEFAULT_PING_URL = "/ping"

type Routes struct {
	gin     *gin.Engine
	Metrics metrics.Metrics
}

func (g *Routes) BindRoutes(cfg Config) {
	g.gin.GET(DEFAULT_PING_URL, g.buildHandlersMapping)
	g.gin.GET("/"+cfg.DefaultPrometheusMetric, prometheusHandler())
}

func (g *Routes) buildHandlersMapping(c *gin.Context) {
	pingCommandHandler := application.NewPingApplication(c.Request.URL.Query())
	cbManager := domain.NewCommandBus()
	cbManager.RegisterHandler(application.PingCommand{}, pingCommandHandler)
	pingController := get.NewPingHandler(g.Metrics)
	pingController.Ping(c, cbManager)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
