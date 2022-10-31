//go:build unit
// +build unit

package engine_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go-skeleton/cmd/engine"
	"go-skeleton/cmd/http"
	metrics2 "go-skeleton/infrastructure/metrics"
)

func TestNewEngine(t *testing.T) {
	t.Run(("Execute new engine"), func(t *testing.T) {
		ctrl := gomock.NewController(t)
		srv := http.NewMockHttpServer(ctrl)
		srv.EXPECT().BuildHttpServer(gomock.Any()).Return()
		metrics := metrics2.NewMockMetricsInterface(ctrl)
		engine := engine.NewEngine(srv, metrics)
		assert.Equal(t, engine.Server, srv)
		assert.Equal(t, engine.Metrics, metrics)
		cfg := engine.BuildEngine()
		expected := http.Config{
			DefaultPrometheusMetric: "metrics",
			AddressPort:             8080,
			AddressIp:               "localhost",
		}
		assert.Equal(t, cfg, expected)
	})
}
