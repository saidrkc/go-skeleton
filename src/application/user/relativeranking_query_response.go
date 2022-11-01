package user

import "go-skeleton/src/domain"

type RelativeUsersScoreResponse []UserScoreResponse

type RelativeUserScoreResponse struct {
	User  int `json:"user"`
	Total int `json:"total"`
}

type RelativeRankingQueryResponse struct {
	UsersScoreResponse `json:"user_score_response"`
}

func (p RelativeRankingQueryResponse) Response() {}

func NewRelativeRankingQueryResponse(domainResponse []domain.UserScoreResponse) RelativeRankingQueryResponse {
	userScoreResponse := buildRelativeUserScoreResponse(domainResponse)
	return RelativeRankingQueryResponse{
		userScoreResponse,
	}
}

func buildRelativeUserScoreResponse(domainResponse []domain.UserScoreResponse) []UserScoreResponse {
	userScoreResponse := make([]UserScoreResponse, 0)
	for _, v := range domainResponse {
		userScoreResponse = append(userScoreResponse, UserScoreResponse{User: v.UserId, Total: v.Total})
	}

	return userScoreResponse
}
