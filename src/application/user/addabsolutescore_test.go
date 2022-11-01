package user_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/user"
	"go-skeleton/src/domain"
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
	t.Run("Add absolute score to User ranking Error Bad request (none filled all mandatory params) ", func(t *testing.T) {
		r := require.New(t)
		ctrl := gomock.NewController(t)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		ctx.Request = &http.Request{
			Method: "/test",
		}
		userRepository := domain.NewMockUserRepositoryInterface(ctrl)
		absolutescoreCommand := user.NewAbsoluteScoreCommand(ctx, mtrcs)
		absoluteScoreApplication := user.NewAbsoluteScoreApplication(ctx, mtrcs, userRepository)
		err := absoluteScoreApplication.Handle(absolutescoreCommand)
		r.Error(err)
	})

	t.Run("Add absolute score to User Bad Request (total and score filled)", func(t *testing.T) {
		r := require.New(t)
		ctrl := gomock.NewController(t)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		req, _ := json.Marshal(user.AbsoluteScoreRequest{
			User:  1,
			Total: 2,
			Score: 2,
		})
		ctx.Request = &http.Request{
			Method: "/test",
			Body:   ioutil.NopCloser(bytes.NewReader([]byte(req))),
		}
		userRepository := domain.NewMockUserRepositoryInterface(ctrl)
		absolutescoreCommand := user.NewAbsoluteScoreCommand(ctx, mtrcs)
		absoluteScoreApplication := user.NewAbsoluteScoreApplication(ctx, mtrcs, userRepository)
		err := absoluteScoreApplication.Handle(absolutescoreCommand)
		r.Error(err)
		r.Equal(err, errors.New("if wants to change total with score, total must be 0 and vice-versa"))
	})

	t.Run("Add absolute score to User good Request", func(t *testing.T) {
		r := require.New(t)
		ctrl := gomock.NewController(t)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		req, _ := json.Marshal(user.AbsoluteScoreRequest{
			User:  1,
			Total: 2,
			Score: 0,
		})
		ctx.Request = &http.Request{
			Method: "/test",
			Body:   ioutil.NopCloser(bytes.NewReader([]byte(req))),
		}
		userRepository := domain.NewMockUserRepositoryInterface(ctrl)
		userRepository.EXPECT().AddAbsoluteScoreToUser(gomock.Any()).Return()
		absolutescoreCommand := user.NewAbsoluteScoreCommand(ctx, mtrcs)
		absoluteScoreApplication := user.NewAbsoluteScoreApplication(ctx, mtrcs, userRepository)
		err := absoluteScoreApplication.Handle(absolutescoreCommand)
		r.NoError(err)
	})
}
