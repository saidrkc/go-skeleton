package get

import (
	"time"

	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/ping"
	"go-skeleton/src/infrastructure/bus/query"
)

type PingHandler struct {
	metrics metrics.Metrics
}

func (h PingHandler) Ping(c *gin.Context, queryBus query.QueryBus) {
	time.Sleep(time.Second * 3)
	pingQuery := ping.NewQuery(c, h.metrics)
	rsp, _ := queryBus.Exec(pingQuery)
	c.JSON(200, rsp)
}

func NewPingHandler(metrics metrics.Metrics) PingHandler {
	return PingHandler{metrics: metrics}
}
