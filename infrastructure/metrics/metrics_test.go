package metrics_test

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"go-skeleton/infrastructure/metrics"
)

func TestNewMetrics(t *testing.T) {
	t.Run("Test add response variable", func(t *testing.T) {
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
		mtrcs.NewTimer("test")
		mtrcs.AddToResponseTime("test")
		mtrcs.IncrementTotalRequests("test")
		mtrcs.IncrementResponseStatus(200)
	})
}
