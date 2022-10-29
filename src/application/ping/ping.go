package ping

import (
	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/infrastructure/bus/query"
)

type Ping struct {
	Context *gin.Context
	Metrics metrics.MetricsInterface
}

func (p Ping) Handle(query query.Query) (query.QueryResponse, error) {
	p.Metrics.AddToResponseTime(p.Context.Request.Method)
	timer := p.Metrics.NewTimer(p.Context.Request.Method)
	defer timer.ObserveDuration()
	p.Metrics.IncrementTotalRequests(p.Context.Request.Method)
	p.Metrics.IncrementResponseStatus(p.Context.Writer.Status())
	return NewPingResponse("{'test'}"), nil
}

func NewPingApplication(context *gin.Context, metrics metrics.MetricsInterface) Ping {
	return Ping{
		Context: context,
		Metrics: metrics,
	}
}
