package user

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/user"
	"go-skeleton/src/infrastructure/bus/query"
)

type AbsoluteRankingHandler struct {
	metrics metrics.MetricsInterface
}

func (h AbsoluteRankingHandler) AbsoluteRanking(c *gin.Context, queryBus query.QueryBus) {
	absoluteRankingQuery := user.NewAbsoluteRankingQuery(c, h.metrics)
	rsp, _ := queryBus.Exec(absoluteRankingQuery)
	c.JSON(200, rsp)
}

func NewAbsoluteRankingHandler(metrics metrics.MetricsInterface) AbsoluteRankingHandler {
	return AbsoluteRankingHandler{metrics: metrics}
}
