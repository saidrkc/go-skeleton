package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/ping"
	"go-skeleton/src/application/pong"
	"go-skeleton/src/infrastructure/bus/command"
	"go-skeleton/src/infrastructure/bus/query"
	"go-skeleton/src/infrastructure/http/get"
	"go-skeleton/src/infrastructure/http/post"
)

const DEFAULT_PING_URL = "/ping"
const DEFAULT_PONG_URL = "/pong"
const DEFAULT_PROMETHEUS_METRICS = "/metrics"

type Routes struct {
	gin     *gin.Engine
	Metrics metrics.Metrics
}

func (g *Routes) BindRoutes() {
	g.gin.POST(DEFAULT_PONG_URL, g.buildPongHandlersMapping)
	g.gin.GET(DEFAULT_PING_URL, g.buildPingHandlersMapping)
	g.gin.GET(DEFAULT_PROMETHEUS_METRICS, prometheusHandler())
}

func (g *Routes) buildPingHandlersMapping(c *gin.Context) {
	pingQueryHandler := ping.NewPingApplication(c, g.Metrics)
	qbManager := query.NewQueryBus()
	qbManager.RegisterHandler(ping.PingQuery{}, pingQueryHandler)
	pingController := get.NewPingHandler(g.Metrics)
	pingController.Ping(c, qbManager)
}

func (g *Routes) buildPongHandlersMapping(c *gin.Context) {
	pongCommandHandler := pong.NewPongApplication(c, g.Metrics)
	cbManager := command.NewCommandBus()
	cbManager.RegisterHandler(pong.PongCommand{}, pongCommandHandler)
	pongController := post.NewPongHandler(g.Metrics)
	pongController.Pong(c, cbManager)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
