package user

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
)

type AbsoluteRankingQuery struct {
	Context *gin.Context
	Metrics metrics.MetricsInterface
}

func (p AbsoluteRankingQuery) QueryID() string {
	return "gopher_AbsoluteRankingQuery"
}

func NewAbsoluteRankingQuery(c *gin.Context, metrics metrics.MetricsInterface) AbsoluteRankingQuery {
	return AbsoluteRankingQuery{
		Context: c,
		Metrics: metrics,
	}
}
