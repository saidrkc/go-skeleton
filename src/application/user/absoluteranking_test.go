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

func TestAbsoluteRanking_Handle(t *testing.T) {
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
	t.Run("Trying to get users scoring absolute but params wrong", func(t *testing.T) {
		r := require.New(t)
		ctrl := gomock.NewController(t)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		uri, _ := url.Parse("")
		ctx.Request = &http.Request{
			Method: "/test",
			URL:    uri,
		}
		userRepository := domain.NewMockUserRepositoryInterface(ctrl)
		absoluteRankingQuery := user.NewAbsoluteRankingQuery(ctx, mtrcs)
		absoluteRanking := user.NewAbsoluteRanking(ctx, mtrcs, userRepository)
		_, err := absoluteRanking.Handle(absoluteRankingQuery)
		r.Error(err)
		r.Equal(err, errors.New("Bad request, param top is mandatory"))
	})

	t.Run("Retrieve ordered by totals with the top variable", func(t *testing.T) {
		expected := []domain.UserScoreResponse{
			{
				UserId: 1,
				Total:  3,
			},
		}
		r := require.New(t)
		ctrl := gomock.NewController(t)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		uri, _ := url.Parse("?top=3")
		ctx.Request = &http.Request{
			Method: "/test",
			URL:    uri,
		}
		userRepository := domain.NewMockUserRepositoryInterface(ctrl)
		absoluteRankingQuery := user.NewAbsoluteRankingQuery(ctx, mtrcs)
		userRepository.EXPECT().AbsoluteRanking(3).Return(expected)
		absoluteRanking := user.NewAbsoluteRanking(ctx, mtrcs, userRepository)
		rsp, err := absoluteRanking.Handle(absoluteRankingQuery)
		r.NoError(err)
		transformation := buildUserResponse(expected)
		r.Equal(rsp, transformation)
	})
}

func buildUserResponse(domainResponse []domain.UserScoreResponse) user.AbsoluteRankingQueryResponse {
	userScoreResponse := make([]user.UserScoreResponse, 0)
	for _, v := range domainResponse {
		userScoreResponse = append(userScoreResponse, user.UserScoreResponse{User: v.UserId, Total: v.Total})
	}

	return user.AbsoluteRankingQueryResponse{
		UsersScoreResponse: userScoreResponse,
	}
}