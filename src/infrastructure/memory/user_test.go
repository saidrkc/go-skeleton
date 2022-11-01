package memory

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go-skeleton/src/domain"
)

func TestUserRepository_AddAbsoluteScoreToUser(t *testing.T) {
	userRepository := NewUserRepository()
	t.Run("Should add total to user id 1 ", func(t *testing.T) {
		expected := UserScoreInMemory{
			UserId: 1,
			Total:  3,
		}
		req := require.New(t)
		userScore := domain.UserScore{
			UserId: 1,
			Total:  3,
			Score:  0,
		}
		userRepository.AddAbsoluteScoreToUser(userScore)
		score, _ := userRepository.FindUserScore(domain.UserScore{UserId: 1})
		req.Equal(expected.UserId, score.UserId)
		req.Equal(expected.Total, score.Total)
	})
	t.Run("Should add score to user id 1 ", func(t *testing.T) {
		expected := UserScoreInMemory{
			UserId: 1,
			Total:  2,
		}
		req := require.New(t)
		userScore := domain.UserScore{
			UserId: 1,
			Total:  0,
			Score:  2,
		}
		userRepository.AddRelativeScoreToUser(userScore)
		score, _ := userRepository.FindUserScore(domain.UserScore{UserId: 1})
		req.Equal(expected.UserId, score.UserId)
		req.Equal(expected.Total, score.Total)
	})
	t.Run("Should add score to user id 1 but total never must be negative", func(t *testing.T) {
		expected := UserScoreInMemory{
			UserId: 1,
			Total:  0,
		}
		req := require.New(t)
		userScore := domain.UserScore{
			UserId: 1,
			Total:  0,
			Score:  -2,
		}
		userRepository.AddRelativeScoreToUser(userScore)
		score, _ := userRepository.FindUserScore(domain.UserScore{UserId: 1})
		req.Equal(expected.UserId, score.UserId)
		req.Equal(expected.Total, score.Total)
	})

	t.Run("Should add score to user id 1 but total never must be positive", func(t *testing.T) {
		expected := UserScoreInMemory{
			UserId: 1,
			Total:  322,
		}
		req := require.New(t)
		userScore := domain.UserScore{
			UserId: 1,
			Total:  0,
			Score:  322,
		}
		userRepository.AddRelativeScoreToUser(userScore)
		score, _ := userRepository.FindUserScore(domain.UserScore{UserId: 1})
		req.Equal(expected.UserId, score.UserId)
		req.Equal(expected.Total, score.Total)
	})
}
