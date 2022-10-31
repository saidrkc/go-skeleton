package http_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-skeleton/cmd/http"
)

func TestNewHttpServer(t *testing.T) {
	t.Run("Retrieves correct instance of http server", func(t *testing.T) {
		httpServer := http.NewHttpServer()
		assert.NotEmpty(t, httpServer.Gin())
		assert.NotEmpty(t, httpServer.Metrics)
		httpServer.BuildHttpServer(httpServer.Metrics)
	})
}
