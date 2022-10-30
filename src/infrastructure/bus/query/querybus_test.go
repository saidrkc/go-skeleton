//go:build unit
// +build unit

package query_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"

	metrics2 "go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/ping"
	"go-skeleton/src/infrastructure/bus/query"
)

func TestQueryBus_Exec(t *testing.T) {

	t.Run("Execute query bus searching for a good response", func(t *testing.T) {
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
		pingQueryHandler := ping.NewPingApplication(ctx, metrics)
		queryBus := query.NewQueryBus()
		queryBus.RegisterHandler(ping.PingQuery{}, pingQueryHandler)
		rsp, _ := queryBus.Exec(ping.PingQuery{})
		expected := ping.PingQueryResponse{Resp: "{'test'}"}
		assert.Equal(t, rsp, expected)
	})

	t.Run("Execute query bus without register a query handler", func(t *testing.T) {
		queryBus := query.NewQueryBus()
		_, err := queryBus.Exec(ping.PingQuery{})
		expected := errors.New("there not any QueryHandler associate to query gopher_PingQuery")
		assert.Equal(t, err.Error(), expected.Error())
	})

	t.Run("Register same handler two times", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		metrics := metrics2.NewMockMetricsInterface(ctrl)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		pingQueryHandler := ping.NewPingApplication(ctx, metrics)
		queryBus := query.NewQueryBus()
		queryBus.RegisterHandler(ping.PingQuery{}, pingQueryHandler)
		err := queryBus.RegisterHandler(ping.PingQuery{}, pingQueryHandler)
		expected := errors.New("the Command gopher_PingQuery is already registered")
		assert.Equal(t, err.Error(), expected.Error())
	})

}
