package command

import (
	"fmt"
)

type Command interface {
	CommandID() string
	PingCommand()
}

type CommandHandler interface {
	Handle(Command) error
}

type CommandBus struct {
	handlersMap map[string]CommandHandler
}

func NewCommandBus() CommandBus {
	return CommandBus{handlersMap: make(map[string]CommandHandler)}
}

func (cb *CommandBus) RegisterHandler(c Command, ch CommandHandler) error {
	cmdID := c.CommandID()

	_, ok := cb.handlersMap[cmdID]
	if ok {
		return fmt.Errorf("the Command %s is already register", cmdID)
	}
	cb.handlersMap[cmdID] = ch
	return nil
}

func (cb CommandBus) Exec(c Command) error {
	cmdID := c.CommandID()

	ch, ok := cb.handlersMap[cmdID]
	if !ok {
		return fmt.Errorf("there not any CommandHandler associate to Command %s", cmdID)
	}
	return ch.Handle(c)
}
