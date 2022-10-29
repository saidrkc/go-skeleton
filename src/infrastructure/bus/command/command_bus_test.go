//go:build unit
// +build unit

package command_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	metrics2 "go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/pong"
	"go-skeleton/src/infrastructure/bus/command"
)

func TestCommandBus_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	metrics := metrics2.NewMockMetricsInterface(ctrl)
	metrics.EXPECT().AddToResponseTime(gomock.Any()).Return()
	metrics.EXPECT().IncrementTotalRequests(gomock.Any()).Return()
	metrics.EXPECT().IncrementResponseStatus(gomock.Any()).Return()
	pongCommandHandler := pong.NewPongApplication(ctx, metrics)
	commandbus := command.NewCommandBus()
	commandbus.RegisterHandler(pong.PongCommand{}, pongCommandHandler)
	commandbus.Exec(pong.PongCommand{})
}
