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

func Test_AbsoluteScoreEndpoint(t *testing.T) {
	t.Run("Test absolute score endpoint bad request response", func(t *testing.T) {
		r := req.New(t)
		endpoint := getBaseUrl("score")
		request := `{"usr": 1}`

		resp, err := http.Post(endpoint, "application/json", strings.NewReader(request))
		r.NoError(err)

		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test absolute score endpoint adding user's and check the response", func(t *testing.T) {
		r := req.New(t)
		endpoint := getBaseUrl("score")
		request := `{"user": 1, "total": 1}`

		resp, err := http.Post(endpoint, "application/json", strings.NewReader(request))
		r.NoError(err)

		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		expected, err := ioutil.ReadFile(`src/infrastructure/test/fixture/response_1_golden.json`)
		r.NoError(err)
		resp2, err := http.Get(getBaseUrl("ranking?top=10"))
		body, err := ioutil.ReadAll(resp2.Body)
		r.NotEmpty(body)
		r.NoError(err)
		r.JSONEq(string(expected), string(body))
	})

	t.Run("Test absolute score endpoint adding only score to the user (total must be modified)", func(t *testing.T) {
		r := req.New(t)
		endpoint := getBaseUrl("score")
		request1 := `{"user": 1, "score": 1}`
		request2 := `{"user": 1, "score": -1}`

		resp, err := http.Post(endpoint, "application/json", strings.NewReader(request1))
		r.NoError(err)

		resp, err = http.Post(endpoint, "application/json", strings.NewReader(request2))
		r.NoError(err)

		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		expected, err := ioutil.ReadFile(`src/infrastructure/test/fixture/response_2_golden.json`)
		r.NoError(err)
		resp2, err := http.Get(getBaseUrl("ranking?top=10"))
		body, err := ioutil.ReadAll(resp2.Body)
		r.NotEmpty(body)
		r.NoError(err)
		r.JSONEq(string(expected), string(body))
	})
}

func Test_RankingQueryEndpoint(t *testing.T) {
	t.Run("Test ranking query score endpoint bad request response", func(t *testing.T) {
		r := req.New(t)
		endpoint := getBaseUrl("ranking")

		resp, err := http.Get(endpoint)
		r.NoError(err)

		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test ranking query with top 3 values", func(t *testing.T) {
		r := req.New(t)
		endpoint := getBaseUrl("score")
		for i := 1; i < 5; i++ {
			request := fmt.Sprintf(`{"user": %d, "total": %d}`, i, i)
			_, err := http.Post(endpoint, "application/json", strings.NewReader(request))
			r.NoError(err)
		}

		expected, err := ioutil.ReadFile(`src/infrastructure/test/fixture/response_3_golden.json`)
		r.NoError(err)
		resp2, err := http.Get(getBaseUrl("ranking?top=3"))
		body, err := ioutil.ReadAll(resp2.Body)
		r.NotEmpty(body)
		r.NoError(err)
		r.JSONEq(string(expected), string(body))
	})
}

func Test_RelativeRankingQueryEndpoint(t *testing.T) {
	t.Run("Test relative ranking query score endpoint bad request response (no params)", func(t *testing.T) {
		r := req.New(t)
		endpoint := getBaseUrl("relative")

		resp, err := http.Get(endpoint)
		r.NoError(err)

		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test relative ranking query score endpoint point 2, 1 around", func(t *testing.T) {
		r := req.New(t)
		endpoint := getBaseUrl("score")
		for i := 1; i < 10; i++ {
			request := fmt.Sprintf(`{"user": %d, "total": %d}`, i, i)
			_, err := http.Post(endpoint, "application/json", strings.NewReader(request))
			r.NoError(err)
		}

		expected, err := ioutil.ReadFile(`src/infrastructure/test/fixture/response_4_golden.json`)
		r.NoError(err)
		resp2, err := http.Get(getBaseUrl("relative?point=2&around=1"))
		body, err := ioutil.ReadAll(resp2.Body)
		r.NotEmpty(body)
		r.NoError(err)
		r.JSONEq(string(expected), string(body))
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
