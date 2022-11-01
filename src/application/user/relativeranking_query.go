package user

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
)

type RelativeRankingQuery struct {
	Context *gin.Context
	Metrics metrics.MetricsInterface
}

func (p RelativeRankingQuery) QueryID() string {
	return "gopher_RelativeRankingQuery"
}

func NewRelativeRankingQuery(c *gin.Context, metrics metrics.MetricsInterface) RelativeRankingQuery {
	return RelativeRankingQuery{
		Context: c,
		Metrics: metrics,
	}
}
