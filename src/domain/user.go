//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package domain

type UserRepositoryInterface interface {
	AddAbsoluteScoreToUser(score UserScore)
	AddRelativeScoreToUser(score UserScore)
}

type UserScore struct {
	UserId int
	Total  int
	Score  int
}
