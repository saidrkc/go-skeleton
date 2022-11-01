package user

import "go-skeleton/src/domain"

type UsersScoreResponse []UserScoreResponse

type UserScoreResponse struct {
	User  int `json:"user"`
	Total int `json:"total"`
}

type AbsoluteRankingQueryResponse struct {
	UsersScoreResponse
}

func (p AbsoluteRankingQueryResponse) Response() {}

func NewAbsoluteRankingQueryResponse(domainResponse []domain.UserScoreResponse) AbsoluteRankingQueryResponse {
	userScoreResponse := buildUserScoreResponse(domainResponse)
	return AbsoluteRankingQueryResponse{
		userScoreResponse,
	}
}

func buildUserScoreResponse(domainResponse []domain.UserScoreResponse) []UserScoreResponse {
	userScoreResponse := make([]UserScoreResponse, 0)
	for _, v := range domainResponse {
		userScoreResponse = append(userScoreResponse, UserScoreResponse{User: v.UserId, Total: v.Total})
	}

	return userScoreResponse
}
