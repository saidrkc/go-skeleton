package memory

import (
	"sort"
	"sync"

	"go-skeleton/src/domain"
)

type UserScoreInMemory struct {
	UserId int
	Total  int
}

type UserRepository struct {
	UsersScore []*UserScoreInMemory
	Mutex      sync.Mutex
}

func (u *UserRepository) AddAbsoluteScoreToUser(score domain.UserScore) {
	scoreUser, slicePos := u.FindUserScore(score)
	if scoreUser.UserId != 0 {
		scoreUser.Total = score.Total
		u.UpdateScoreToUsersInMemory(UserScoreInMemory{UserId: score.UserId, Total: scoreUser.Total}, slicePos)
		return
	}

	sc := UserScoreInMemory{
		UserId: score.UserId,
		Total:  score.Total,
	}

	u.AddScoreToUsersInMemory(sc)
}

func (u *UserRepository) AddRelativeScoreToUser(score domain.UserScore) {
	scoreUser, slicePos := u.FindUserScore(score)
	if score.UserId != 0 && slicePos >= 0 {
		scoreUser.Total = scoreUser.Total + score.Score
		if scoreUser.Total < 0 {
			scoreUser.Total = 0
		}
		u.UpdateScoreToUsersInMemory(UserScoreInMemory{UserId: score.UserId, Total: scoreUser.Total}, slicePos)
		return
	}

	total := func() int {
		total := scoreUser.Total + score.Score
		if total < 0 {
			return 0
		}
		return total
	}

	sc := UserScoreInMemory{
		UserId: score.UserId,
		Total:  total(),
	}

	u.AddScoreToUsersInMemory(sc)
}

func (u *UserRepository) UpdateScoreToUsersInMemory(user UserScoreInMemory, slicePos int) {
	u.Mutex.Lock()
	u.UsersScore[slicePos] = &user
	u.Mutex.Unlock()
}

func (u *UserRepository) AddScoreToUsersInMemory(user UserScoreInMemory) {
	u.Mutex.Lock()
	u.UsersScore = append(u.UsersScore, &user)
	u.Mutex.Unlock()
}

func (u *UserRepository) FindUserScore(score domain.UserScore) (*UserScoreInMemory, int) {
	for k, v := range u.UsersScore {
		if v.UserId == score.UserId {
			return v, k
		}
	}
	return &UserScoreInMemory{}, -1
}

func (u *UserRepository) AbsoluteRanking(ranking int) []domain.UserScoreResponse {
	usersScore := make([]domain.UserScoreResponse, 0)
	sort.Slice(u.UsersScore, func(i, j int) bool {
		return u.UsersScore[i].Total > u.UsersScore[j].Total
	})

	for k, v := range u.UsersScore {
		if k == ranking {
			return usersScore
		}
		usersScore = append(usersScore, domain.UserScoreResponse{
			UserId: v.UserId,
			Total:  v.Total,
		})
	}

	return usersScore
}

func (u *UserRepository) RelativeRanking(point int, around int) ([]domain.UserScoreResponse, error) {
	if len(u.UsersScore) == 0 {
		return []domain.UserScoreResponse{}, nil
	}

	usersScore := make([]domain.UserScoreResponse, 0)
	sort.Slice(u.UsersScore, func(i, j int) bool {
		return u.UsersScore[i].Total > u.UsersScore[j].Total
	})
	point = point - 1
	offset := point - around
	if offset <= 0 {
		offset = 0
	}

	end := point + around
	if end >= len(u.UsersScore) {
		end = len(u.UsersScore) - 1
	}

	score := u.UsersScore[offset:end]
	score = append(score, u.UsersScore[end])

	for _, v := range score {
		usersScore = append(usersScore, domain.UserScoreResponse{
			UserId: v.UserId,
			Total:  v.Total,
		})
	}

	return usersScore, nil
}

func NewUserRepository() UserRepository {
	return UserRepository{
		UsersScore: make([]*UserScoreInMemory, 0),
	}
}
