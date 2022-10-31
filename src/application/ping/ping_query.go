package ping

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
)

type PingQuery struct {
	Context *gin.Context
	Metrics metrics.MetricsInterface
}

func (p PingQuery) PingQuery() {}

func (p PingQuery) QueryID() string {
	return "gopher_PingQuery"
}

func NewQuery(c *gin.Context, metrics metrics.MetricsInterface) PingQuery {
	return PingQuery{
		Context: c,
		Metrics: metrics,
	}
}
