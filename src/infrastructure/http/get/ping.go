package get

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application"
	"go-skeleton/src/domain"
)

type PingHandler struct {
	metrics metrics.Metrics
}

func (h PingHandler) Ping(c *gin.Context, commandBus domain.CommandBus) {
	h.metrics.HttpResponseTime.WithLabelValues(c.Request.Method)
	timer := prometheus.NewTimer(h.metrics.HttpResponseTime.WithLabelValues(c.Request.Method))
	defer timer.ObserveDuration()
	h.metrics.TotalRequest.WithLabelValues(c.Request.Method).Inc()
	c.JSON(500, "{'test'}")
	h.metrics.ResponseStatus.WithLabelValues(strconv.Itoa(c.Writer.Status())).Inc()
	commandBus.Exec(application.PingCommand{})
	time.Sleep(time.Second * 3)
}

func NewPingHandler(metrics metrics.Metrics) PingHandler {
	return PingHandler{metrics: metrics}
}
