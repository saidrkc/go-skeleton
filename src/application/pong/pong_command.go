package pong

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
)

type PongCommand struct {
	Context *gin.Context
	Metrics metrics.MetricsInterface
}

func (p PongCommand) PingCommand() {}

func (p PongCommand) CommandID() string {
	return "gopher_PongCommand"
}

func NewCommand(c *gin.Context, metrics metrics.MetricsInterface) PongCommand {
	return PongCommand{
		Context: c,
		Metrics: metrics,
	}
}
