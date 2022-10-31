package post

import (
	"time"

	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/pong"
	"go-skeleton/src/infrastructure/bus/command"
)

type PongHandler struct {
	metrics metrics.MetricsInterface
}

func (h PongHandler) Pong(c *gin.Context, commandbus command.CommandBus) {
	time.Sleep(time.Second * 1)
	pongCommand := pong.NewCommand(c, h.metrics)
	commandbus.Exec(pongCommand)
	c.JSON(200, "{'ok'}")
}

func NewPongHandler(metrics metrics.MetricsInterface) PongHandler {
	return PongHandler{metrics: metrics}
}
