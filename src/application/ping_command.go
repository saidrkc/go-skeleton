package application

type PingCommand struct {
}

func (p PingCommand) CommandID() string {
	return "gopher_PingCommand"
}
