package application

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"go-skeleton/infrastructure/metrics"
)

func Ping(metrics metrics.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics.HttpResponseTime.WithLabelValues(c.Request.Method)
		timer := prometheus.NewTimer(metrics.HttpResponseTime.WithLabelValues(c.Request.Method))
		metrics.TotalRequest.WithLabelValues(c.Request.Method).Inc()
		c.JSON(500, "{'test'}")
		metrics.ResponseStatus.WithLabelValues(strconv.Itoa(c.Writer.Status())).Inc()
		time.Sleep(time.Second * 3)
		defer timer.ObserveDuration()
	}
}
