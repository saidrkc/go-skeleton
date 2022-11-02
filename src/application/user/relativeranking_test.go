package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/user"
	"go-skeleton/src/domain"
)

func TestRelativeRanking_Handle(t *testing.T) {
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
	t.Run("Trying to get users relative scoring but params wrong", func(t *testing.T) {
		r := require.New(t)
		ctrl := gomock.NewController(t)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		uri, _ := url.Parse("?around=2")
		ctx.Request = &http.Request{
			Method: "/test",
			URL:    uri,
		}
		userRepository := domain.NewMockUserRepositoryInterface(ctrl)
		relativeRankingQuery := user.NewRelativeRankingQuery(ctx, mtrcs)
		relativeRanking := user.NewRelativeRanking(ctx, mtrcs, userRepository)
		_, err := relativeRanking.Handle(relativeRankingQuery)
		r.Error(err)
		r.Equal(err, errors.New("Bad request, param point is mandatory"))
	})
	t.Run("Trying to get users relative scoring but params wrong", func(t *testing.T) {
		r := require.New(t)
		ctrl := gomock.NewController(t)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		uri, _ := url.Parse("?point=2")
		ctx.Request = &http.Request{
			Method: "/test",
			URL:    uri,
		}
		userRepository := domain.NewMockUserRepositoryInterface(ctrl)
		relativeRankingQuery := user.NewRelativeRankingQuery(ctx, mtrcs)
		relativeRanking := user.NewRelativeRanking(ctx, mtrcs, userRepository)
		_, err := relativeRanking.Handle(relativeRankingQuery)
		r.Error(err)
		r.Equal(err, errors.New("Bad request, param around is mandatory"))
	})

	t.Run("Retrieve relative scoring by point and around parameters", func(t *testing.T) {
		expected := []domain.UserScoreResponse{
			{
				UserId: 1,
				Total:  3,
			},
			{
				UserId: 2,
				Total:  3,
			},
			{
				UserId: 3,
				Total:  3,
			},
		}
		r := require.New(t)
		ctrl := gomock.NewController(t)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		uri, _ := url.Parse("?point=1&around=1")
		ctx.Request = &http.Request{
			Method: "/test",
			URL:    uri,
		}
		userRepository := domain.NewMockUserRepositoryInterface(ctrl)
		relativeRankingQuery := user.NewRelativeRankingQuery(ctx, mtrcs)
		userRepository.EXPECT().RelativeRanking(1, 1).Return(expected, nil)
		relativeRanking := user.NewRelativeRanking(ctx, mtrcs, userRepository)
		rsp, err := relativeRanking.Handle(relativeRankingQuery)
		r.NoError(err)
		transformation := buildRelativeUserResponse(expected)
		r.Equal(rsp, transformation)
	})
}

func buildRelativeUserResponse(domainResponse []domain.UserScoreResponse) user.RelativeRankingQueryResponse {
	userScoreResponse := make([]user.UserScoreResponse, 0)
	for _, v := range domainResponse {
		userScoreResponse = append(userScoreResponse, user.UserScoreResponse{User: v.UserId, Total: v.Total})
	}

	return user.RelativeRankingQueryResponse{
		UsersScoreResponse: userScoreResponse,
	}
}
