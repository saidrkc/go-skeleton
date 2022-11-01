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

type AbsoluteRanking struct {
	Context   *gin.Context
	Metrics   metrics.MetricsInterface
	UserScore domain.UserRepositoryInterface
}

func (p AbsoluteRanking) Handle(query query.Query) (query.QueryResponse, error) {
	p.Metrics.AddToResponseTime(p.Context.Request.Method)
	timer := p.Metrics.NewTimer(p.Context.Request.Method)
	defer timer.ObserveDuration()
	p.Metrics.IncrementTotalRequests(p.Context.Request.Method)
	p.Metrics.IncrementResponseStatus(p.Context.Writer.Status())

	params := p.Context.Request.URL.Query()
	if top := params.Get("top"); top == "" {
		p.Context.AbortWithStatus(http.StatusBadRequest)
		return NewAbsoluteRankingQueryResponse([]domain.UserScoreResponse{}), errors.New("Bad request, param top is mandatory")
	}

	top := params.Get("top")
	tp, _ := strconv.Atoi(top)
	usersScore := p.UserScore.AbsoluteRanking(tp)

	return NewAbsoluteRankingQueryResponse(usersScore), nil
}

func NewAbsoluteRanking(context *gin.Context, metrics metrics.MetricsInterface, userScore domain.UserRepositoryInterface) AbsoluteRanking {
	return AbsoluteRanking{
		Context:   context,
		Metrics:   metrics,
		UserScore: userScore,
	}
}
