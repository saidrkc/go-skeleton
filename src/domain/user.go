//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package domain

type UserRepositoryInterface interface {
	AddAbsoluteScoreToUser(score UserScore)
	AddRelativeScoreToUser(score UserScore)
	AbsoluteRanking(ranking int) []UserScoreResponse
	RelativeRanking(point int, around int) []UserScoreResponse
}

type UserScore struct {
	UserId int
	Total  int
	Score  int
}

type UserScoreResponse struct {
	UserId int
	Total  int
}
