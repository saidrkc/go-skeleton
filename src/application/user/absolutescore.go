package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/domain"
	"go-skeleton/src/infrastructure/bus/command"
)

type AbsoluteScoreRequest struct {
	User  int `json:"user" binding:"required"`
	Total int `json:"total"`
	Score int `json:"score"`
}

type AbsoluteScore struct {
	Context   *gin.Context
	Metrics   metrics.MetricsInterface
	UserScore domain.UserRepositoryInterface
}

func (p AbsoluteScore) Handle(command command.Command) error {
	p.Metrics.AddToResponseTime(p.Context.Request.Method)
	timer := p.Metrics.NewTimer(p.Context.Request.Method)
	defer timer.ObserveDuration()
	p.Metrics.IncrementTotalRequests(p.Context.Request.Method)

	var expectedRequest AbsoluteScoreRequest
	if err := p.Context.ShouldBindJSON(&expectedRequest); err != nil {
		return errors.New("bad Request, parameters are not well")
	}

	if expectedRequest.Score != 0 && expectedRequest.Total != 0 {
		return errors.New("if wants to change total with score, total must be 0 and vice-versa")
	}

	if expectedRequest.Score != 0 {
		userScore := domain.UserScore{
			UserId: expectedRequest.User,
			Score:  expectedRequest.Score,
		}

		p.UserScore.AddRelativeScoreToUser(userScore)
		return nil
	}

	userScore := domain.UserScore{
		UserId: expectedRequest.User,
		Total:  expectedRequest.Total,
	}

	p.UserScore.AddAbsoluteScoreToUser(userScore)
	p.Metrics.IncrementResponseStatus(p.Context.Writer.Status())
	return nil
}

func NewAbsoluteScoreApplication(context *gin.Context, metrics metrics.MetricsInterface, userScore domain.UserRepositoryInterface) AbsoluteScore {
	return AbsoluteScore{
		Context:   context,
		Metrics:   metrics,
		UserScore: userScore,
	}
}
