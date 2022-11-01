package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/domain"
	"go-skeleton/src/infrastructure/bus/command"
	"go-skeleton/src/infrastructure/inmemory"
)

type AbsoluteScoreRequest struct {
	User  int     `json:"user" binding:"required"`
	Total float32 `json:"total" binding:"required"`
}

type AbsoluteScore struct {
	Context   *gin.Context
	Metrics   metrics.MetricsInterface
	UserScore inmemory.UserRepository
}

func (p AbsoluteScore) Handle(command command.Command) error {
	p.Metrics.AddToResponseTime(p.Context.Request.Method)
	timer := p.Metrics.NewTimer(p.Context.Request.Method)
	defer timer.ObserveDuration()
	p.Metrics.IncrementTotalRequests(p.Context.Request.Method)

	var expectedRequest AbsoluteScoreRequest
	if err := p.Context.ShouldBindJSON(&expectedRequest); err != nil {
		p.Context.AbortWithStatus(http.StatusBadRequest)
	}
	userScore := domain.UserScore{
		UserId: expectedRequest.User,
		Total:  expectedRequest.Total,
		Score:  0,
	}

	p.UserScore.AddAbsoluteScoreToUser(userScore)
	p.Metrics.IncrementResponseStatus(p.Context.Writer.Status())
	return nil
}

func NewAbsoluteScoreApplication(context *gin.Context, metrics metrics.MetricsInterface, userScore inmemory.UserRepository) AbsoluteScore {
	return AbsoluteScore{
		Context:   context,
		Metrics:   metrics,
		UserScore: userScore,
	}
}
