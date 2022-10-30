//go:build unit
// +build unit

package command_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"

	metrics2 "go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/pong"
	"go-skeleton/src/infrastructure/bus/command"
)

func Test_ExecCommandBus(t *testing.T) {
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
	pongCommandHandler := pong.NewPongApplication(ctx, metrics)
	commandbus := command.NewCommandBus()
	commandbus.RegisterHandler(pong.PongCommand{}, pongCommandHandler)
	commandbus.Exec(pong.PongCommand{})
	assert.Equal(pong.Pong{}, pong.Pong{})
}
