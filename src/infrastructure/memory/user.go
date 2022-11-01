package memory

import (
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

func (u *UserRepository) FillUserScore(numberOfUsers int) {
	var i int
	for i = 0; i < numberOfUsers; i++ {
		u.AddScoreToUsersInMemory(UserScoreRandom(i))
	}
}

func NewUserRepository() UserRepository {
	return UserRepository{
		UsersScore: make([]*UserScoreInMemory, 0),
	}
}
