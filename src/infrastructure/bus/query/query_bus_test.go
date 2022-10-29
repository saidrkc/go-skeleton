package query_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"

	metrics2 "go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/ping"
	"go-skeleton/src/infrastructure/bus/query"
)

func TestQueryBus_Exec(t *testing.T) {
	var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
	timer := prometheus.NewTimer(httpDuration.WithLabelValues("test"))
	ctrl := gomock.NewController(t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = &http.Request{
		Method: "test",
	}
	metrics := metrics2.NewMockMetricsInterface(ctrl)
	metrics.EXPECT().AddToResponseTime(gomock.Any()).Return()
	metrics.EXPECT().NewTimer(gomock.Any()).Return(timer)
	metrics.EXPECT().IncrementTotalRequests(gomock.Any()).Return()
	metrics.EXPECT().IncrementResponseStatus(gomock.Any()).Return()

	t.Run("Execute query bus searching for a good response", func(t *testing.T) {
		pongCommandHandler := ping.NewPingApplication(ctx, metrics)
		queryBus := query.NewQueryBus()
		queryBus.RegisterHandler(ping.PingQuery{}, pongCommandHandler)
		rsp, _ := queryBus.Exec(ping.PingQuery{})
		expected := ping.PingQueryResponse{Resp: "{'test'}"}
		assert.Equal(t, rsp, expected)
	})

}
