package inmemory

import (
	"go-skeleton/src/domain"
)

type UsersScoreInMemory struct {
	UsersScore []UserScoreInMemory
}

type UserScoreInMemory struct {
	UserId int
	Total  float32
}

type UserRepository struct {
	UsersScoreMemory UsersScoreInMemory
}

func (u *UserRepository) AddAbsoluteScoreToUser(score domain.UserScore) {
	scoreUser := u.FindUserScore(score)
	if scoreUser.UserId != 0 {
		scoreUser.Total = scoreUser.Total + score.Total
		return
	}

	sc := UserScoreInMemory{
		UserId: score.UserId,
		Total:  score.Total,
	}

	u.AddScoreToUsersInMemory(sc)
}

func (u *UserRepository) AddRelativeScoreToUser(score domain.UserScore) {
	scoreUser := u.FindUserScore(score)
	if score.UserId != 0 {
		scoreUser.Total = scoreUser.Total + score.Score
	}

	sc := UserScoreInMemory{
		UserId: score.UserId,
		Total:  score.Total,
	}

	u.AddScoreToUsersInMemory(sc)
}

func (u *UserRepository) AddScoreToUsersInMemory(user UserScoreInMemory) {
	u.UsersScoreMemory.UsersScore = append(u.UsersScoreMemory.UsersScore, user)
}

func (u *UserRepository) FindUserScore(score domain.UserScore) UserScoreInMemory {
	for _, v := range u.UsersScoreMemory.UsersScore {
		if v.UserId == score.UserId {
			return v
		}
	}

	return UserScoreInMemory{}
}

func (u *UserRepository) FillUserScore(numberOfUsers int) {
	var i int
	for i = 0; i < numberOfUsers; i++ {
		u.AddScoreToUsersInMemory(UserScoreRandom(i))
	}
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}
