package ping

type PingQueryResponse struct {
	Resp string
}

func (p PingQueryResponse) Response() {}

func NewPingResponse(queryResponse string) PingQueryResponse {
	return PingQueryResponse{
		Resp: queryResponse,
	}
}
