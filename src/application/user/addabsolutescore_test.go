package user_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/user"
)

func TestAbsoluteScore_Handle(t *testing.T) {
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
	t.Run("Add absolute score to User ranking ", func(t *testing.T) {
		r := require.New(t)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		absolutescoreCommand := user.NewAbsoluteScoreCommand(ctx, mtrcs)
		absoluteScoreApplication := user.NewAbsoluteScoreApplication(ctx, mtrcs, usersRepository)
		err := absoluteScoreApplication.Handle(absolutescoreCommand)
		r.NoError(err)
	})
}
