package user

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/user"
	"go-skeleton/src/infrastructure/bus/command"
)

type AbsoluteScoreHandler struct {
	metrics metrics.MetricsInterface
}

func (h AbsoluteScoreHandler) AbsoluteScore(c *gin.Context, commandbus command.CommandBus) {
	absoluteScoreCommand := user.NewAbsoluteScoreCommand(c, h.metrics)
	commandbus.Exec(absoluteScoreCommand)
	c.JSON(200, "{}")
}

func NewAbsoluteScoreHandler(metrics metrics.MetricsInterface) AbsoluteScoreHandler {
	return AbsoluteScoreHandler{metrics: metrics}
}
