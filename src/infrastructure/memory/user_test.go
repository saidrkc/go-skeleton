//go:build unit
// +build unit

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
}

func TestUserRepository_AddRelativeScoreToUser(t *testing.T) {
	userRepository := NewUserRepository()
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

func TestUserRepository_AbsoluteRanking(t *testing.T) {
	userRepository := NewUserRepository()
	t.Run("Should return absolute ranking sorted by totals", func(t *testing.T) {
		req := require.New(t)
		usersScore := providerUsersScoreSort()
		for _, v := range usersScore {
			userRepository.AddAbsoluteScoreToUser(v)
		}

		expectedUserScoreSort := expectedUserScoreSort()
		sortedUserScore := userRepository.AbsoluteRanking(5)
		req.Equal(sortedUserScore, expectedUserScoreSort)
	})

	t.Run("Should return absolute ranking sorted by totals, but ranking relative top 1", func(t *testing.T) {
		req := require.New(t)
		usersScore := providerUsersScoreSort()
		for _, v := range usersScore {
			userRepository.AddAbsoluteScoreToUser(v)
		}

		expectedUserScoreSort := expectedUserScoreSortWithRanked()
		sortedUserScore := userRepository.AbsoluteRanking(1)
		req.Equal(sortedUserScore, expectedUserScoreSort)
	})
}

func TestUserRepository_RelativeRanking(t *testing.T) {
	t.Run("Should return relative ranking sorted by totals around position 100/3", func(t *testing.T) {
		userRepository := NewUserRepository()
		req := require.New(t)
		usersScore := providerUsersScoreRelative()
		for _, v := range usersScore {
			userRepository.AddAbsoluteScoreToUser(v)
		}

		expectedUserRelativeRanking := expectedUserRelativeRanking()
		sortedUserScore, err := userRepository.RelativeRanking(100, 3)
		req.NoError(err)
		req.Equal(sortedUserScore, expectedUserRelativeRanking)
	})
	t.Run("Should return relative ranking sorted by totals around position 106/3 (only returning pos 106 and 3 before)", func(t *testing.T) {
		userRepository := NewUserRepository()
		req := require.New(t)
		usersScore := providerUsersScoreRelative()
		for _, v := range usersScore {
			userRepository.AddAbsoluteScoreToUser(v)
		}

		expectedUserRelativeRankingLimitAtTheTop := expectedUserRelativeRankingLimitAtTheTop()
		sortedUserScore, err := userRepository.RelativeRanking(106, 3)
		req.NoError(err)
		req.Equal(sortedUserScore, expectedUserRelativeRankingLimitAtTheTop)
	})
	t.Run("Should return relative ranking sorted by totals around position 1/3 (only returning pos 1 and 3 after)", func(t *testing.T) {
		userRepository := NewUserRepository()
		req := require.New(t)
		usersScore := providerUsersScoreRelative()
		for _, v := range usersScore {
			userRepository.AddAbsoluteScoreToUser(v)
		}

		expectedUserRelativeRankingLimitAtTheBottom := expectedUserRelativeRankingLimitAtTheBottom()
		sortedUserScore, err := userRepository.RelativeRanking(1, 3)
		req.NoError(err)
		req.Equal(sortedUserScore, expectedUserRelativeRankingLimitAtTheBottom)
		req.Equal(len(sortedUserScore), 4)
	})

}

func expectedUserScoreSort() []domain.UserScoreResponse {
	return []domain.UserScoreResponse{
		{
			UserId: 3,
			Total:  21,
		},
		{
			UserId: 2,
			Total:  20,
		},
		{
			UserId: 1,
			Total:  3,
		},
	}
}

func expectedUserScoreSortWithRanked() []domain.UserScoreResponse {
	return []domain.UserScoreResponse{
		{
			UserId: 3,
			Total:  21,
		},
	}
}

func expectedUserRelativeRanking() []domain.UserScoreResponse {
	return []domain.UserScoreResponse{
		{
			UserId: 10,
			Total:  10,
		},
		{
			UserId: 9,
			Total:  9,
		},
		{
			UserId: 8,
			Total:  8,
		},
		{
			UserId: 7,
			Total:  7,
		},
		{
			UserId: 6,
			Total:  6,
		},
		{
			UserId: 5,
			Total:  5,
		},
		{
			UserId: 4,
			Total:  4,
		},
	}
}

func expectedUserRelativeRankingLimitAtTheBottom() []domain.UserScoreResponse {
	return []domain.UserScoreResponse{
		{
			UserId: 106,
			Total:  106,
		},
		{
			UserId: 105,
			Total:  105,
		},
		{
			UserId: 104,
			Total:  104,
		},
		{
			UserId: 103,
			Total:  103,
		},
	}
}

func expectedUserRelativeRankingLimitAtTheTop() []domain.UserScoreResponse {
	return []domain.UserScoreResponse{
		{
			UserId: 4,
			Total:  4,
		},
		{
			UserId: 3,
			Total:  3,
		},
		{
			UserId: 2,
			Total:  2,
		},
		{
			UserId: 1,
			Total:  1,
		},
		{
			UserId: 0,
			Total:  0,
		},
	}
}

func providerUsersScoreSort() []domain.UserScore {
	return []domain.UserScore{
		{
			UserId: 1,
			Total:  3,
		},
		{
			UserId: 2,
			Total:  20,
		},
		{
			UserId: 3,
			Total:  21,
		},
	}
}

func providerUsersScoreRelative() []domain.UserScore {
	usersScore := make([]domain.UserScore, 0)
	for i := 0; i <= 106; i++ {
		user := domain.UserScore{
			UserId: i,
			Total:  i,
		}
		usersScore = append(usersScore, user)
	}

	return usersScore

}
