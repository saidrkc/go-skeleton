package ping

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
)

type PingQuery struct {
	Context *gin.Context
	Metrics metrics.Metrics
}

func (p PingQuery) PingQuery() {}

func (p PingQuery) QueryID() string {
	return "gopher_PingQuery"
}

func NewQuery(c *gin.Context, metrics metrics.Metrics) PingQuery {
	return PingQuery{
		Context: c,
		Metrics: metrics,
	}
}
