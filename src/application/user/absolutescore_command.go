package user

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
)

type AbsoluteScoreCommand struct {
	Context *gin.Context
	Metrics metrics.MetricsInterface
}

func (p AbsoluteScoreCommand) AbsoluteScoreCommand() {}

func (p AbsoluteScoreCommand) CommandID() string {
	return "gopher_AbsoluteScoreCommand"
}

func NewAbsoluteScoreCommand(c *gin.Context, metrics metrics.MetricsInterface) AbsoluteScoreCommand {
	return AbsoluteScoreCommand{
		Context: c,
		Metrics: metrics,
	}
}
