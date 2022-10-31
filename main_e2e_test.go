//go:build e2e
// +build e2e

package main_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	req "github.com/stretchr/testify/require"

	"go-skeleton/cmd/engine"
	http2 "go-skeleton/cmd/http"
)

func TestMain(m *testing.M) {
	srv := http2.NewHttpServer()
	eng := engine.NewEngine(srv, srv.Metrics)
	eng.BuildEngine()
	go eng.RunEngine(8080)

	ctx2, _ := context.WithTimeout(context.Background(), 5*time.Second)
Loop:
	for {
		select {
		case <-ctx2.Done():
			fmt.Println("server: cannot connect to server")
			os.Exit(1)
		default:
			if _, err := http.Get(getBaseUrl("ping")); err == nil {
				break Loop
			}
		}
	}

	code := m.Run()

	os.Exit(code)
}

func Test_PingEndpoint(t *testing.T) {
	t.Run("Test ping endpoint response", func(t *testing.T) {
		r := req.New(t)
		endpoint := getBaseUrl("ping")

		resp, _ := http.Get(endpoint)

		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		r.NoError(err)
		expected := "{\"Resp\":\"{'test'}\"}"
		r.JSONEq(string(expected), string(body))
	})
}

func Test_PongEndpoint(t *testing.T) {
	t.Run("Test pong endpoint response", func(t *testing.T) {
		r := req.New(t)
		endpoint := getBaseUrl("pong")
		request := `{}`

		resp, err := http.Post(endpoint, "application/json", strings.NewReader(request))
		r.NoError(err)

		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		r.NoError(err)
	})
}

func getBaseUrl(en string) string {
	const baseUrl = "http://127.0.0.1%s/%s"
	endpoint := fmt.Sprintf(baseUrl, ":8080", en)

	if v, e := os.LookupEnv("SERVER_LISTEN_ADDR"); e {
		endpoint = fmt.Sprintf(baseUrl, v, en)
	}
	return endpoint
}
