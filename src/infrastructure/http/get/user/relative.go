package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/user"
	"go-skeleton/src/infrastructure/bus/query"
)

type RelativeRankingHandler struct {
	metrics metrics.MetricsInterface
}

func (h RelativeRankingHandler) RelativeRanking(c *gin.Context, queryBus query.QueryBus) {
	relativeRatingQuery := user.NewRelativeRankingQuery(c, h.metrics)
	rsp, err := queryBus.Exec(relativeRatingQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, rsp)
}

func NewRelativeRankingHandler(metrics metrics.MetricsInterface) RelativeRankingHandler {
	return RelativeRankingHandler{metrics: metrics}
}
