package http_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"go-skeleton/cmd/http"
	"go-skeleton/infrastructure/metrics"
)

func TestRoutes_BindRoutes(t *testing.T) {
	var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
	var totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path"},
	)
	var responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_status",
			Help: "Status of HTTP response",
		},
		[]string{"status"},
	)
	mtrcs := metrics.NewMetrics(httpDuration, totalRequests, responseStatus)
	t.Run("Binding Routes ", func(t *testing.T) {
		routes := http.Routes{Metrics: mtrcs, Gin: gin.Default()}
		routes.BindRoutes()
	})
}
