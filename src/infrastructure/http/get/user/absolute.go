package user

import (
	"net/http"

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
	rsp, err := queryBus.Exec(absoluteRankingQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, rsp)
}

func NewAbsoluteRankingHandler(metrics metrics.MetricsInterface) AbsoluteRankingHandler {
	return AbsoluteRankingHandler{metrics: metrics}
}
