package pong

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/infrastructure/bus/command"
)

type Pong struct {
	Context *gin.Context
	Metrics metrics.MetricsInterface
}

func (p Pong) Handle(command command.Command) error {
	p.Metrics.AddToResponseTime(p.Context.Request.Method)
	timer := p.Metrics.NewTimer(p.Context.Request.Method)
	defer timer.ObserveDuration()
	p.Metrics.IncrementTotalRequests(p.Context.Request.Method)
	p.Metrics.IncrementResponseStatus(p.Context.Writer.Status())
	return nil
}

func NewPongApplication(context *gin.Context, metrics metrics.MetricsInterface) Pong {
	return Pong{
		Context: context,
		Metrics: metrics,
	}
}
