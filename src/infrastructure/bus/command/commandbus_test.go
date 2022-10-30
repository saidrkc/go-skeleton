//go:build unit
// +build unit

package command_test

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
	"go-skeleton/src/application/pong"
	"go-skeleton/src/infrastructure/bus/command"
)

func Test_ExecCommandBus(t *testing.T) {
	t.Run("Execute command bus", func(t *testing.T) {
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
	})

	t.Run("Without register command handler", func(t *testing.T) {
		commandbus := command.NewCommandBus()
		err := commandbus.Exec(pong.PongCommand{})
		expected := errors.New("there not any CommandHandler associate to Command gopher_PongCommand")
		assert.Equal(t, err.Error(), expected.Error())
	})

	t.Run("Register same handler two times", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		metrics := metrics2.NewMockMetricsInterface(ctrl)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		pongCommandHandler := pong.NewPongApplication(ctx, metrics)
		commandbus := command.NewCommandBus()
		commandbus.RegisterHandler(pong.PongCommand{}, pongCommandHandler)
		err := commandbus.RegisterHandler(pong.PongCommand{}, pongCommandHandler)
		expected := errors.New("the Command gopher_PongCommand is already registered")
		assert.Equal(t, err.Error(), expected.Error())
	})
}
