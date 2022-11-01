package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/domain"
	"go-skeleton/src/infrastructure/bus/query"
)

type RelativeRanking struct {
	Context   *gin.Context
	Metrics   metrics.MetricsInterface
	UserScore domain.UserRepositoryInterface
}

func (p RelativeRanking) Handle(query query.Query) (query.QueryResponse, error) {
	p.Metrics.AddToResponseTime(p.Context.Request.Method)
	timer := p.Metrics.NewTimer(p.Context.Request.Method)
	defer timer.ObserveDuration()
	p.Metrics.IncrementTotalRequests(p.Context.Request.Method)
	p.Metrics.IncrementResponseStatus(p.Context.Writer.Status())

	params := p.Context.Request.URL.Query()
	if point := params.Get("point"); point == "" {
		p.Context.AbortWithStatus(http.StatusBadRequest)
		return NewAbsoluteRankingQueryResponse([]domain.UserScoreResponse{}), errors.New("Bad request, param point is mandatory")
	}

	if around := params.Get("around"); around == "" {
		p.Context.AbortWithStatus(http.StatusBadRequest)
		return NewAbsoluteRankingQueryResponse([]domain.UserScoreResponse{}), errors.New("Bad request, param around is mandatory")
	}

	point := params.Get("point")
	pt, _ := strconv.Atoi(point)
	around := params.Get("around")
	ar, _ := strconv.Atoi(around)
	usersScore := p.UserScore.RelativeRanking(pt, ar)

	return NewAbsoluteRankingQueryResponse(usersScore), nil
}

func NewRelativeRanking(context *gin.Context, metrics metrics.MetricsInterface, userScore domain.UserRepositoryInterface) RelativeRanking {
	return RelativeRanking{
		Context:   context,
		Metrics:   metrics,
		UserScore: userScore,
	}
}
