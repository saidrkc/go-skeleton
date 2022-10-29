package pong

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
)

type PongCommand struct {
	Context *gin.Context
	Metrics metrics.Metrics
}

func (p PongCommand) PingCommand() {}

func (p PongCommand) CommandID() string {
	return "gopher_PingCommand"
}

func NewCommand(c *gin.Context, metrics metrics.Metrics) PongCommand {
	return PongCommand{
		Context: c,
		Metrics: metrics,
	}
}
