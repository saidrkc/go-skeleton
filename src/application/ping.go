package application

import (
	"errors"

	"go-skeleton/src/domain"
)

type Ping struct {
	PingRequest map[string][]string
}

func (p Ping) Handle(c domain.Command) error {
	cmd, ok := c.(domain.Command)
	if !ok {
		return errors.New("invalid command")
	}

	cmd.CommandID()
	return nil
}

func NewPingApplication(pingReq map[string][]string) Ping {
	return Ping{PingRequest: pingReq}
}
