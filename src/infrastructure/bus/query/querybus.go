package query

import "fmt"

type Query interface {
	QueryID() string
	PingQuery()
}

type QueryHandler interface {
	Handle(Query) (QueryResponse, error)
}

type QueryResponse interface {
	Response()
}

type QueryBus struct {
	handlersMap map[string]QueryHandler
}

func NewQueryBus() QueryBus {
	return QueryBus{handlersMap: make(map[string]QueryHandler)}
}

func (cb *QueryBus) RegisterHandler(c Query, ch QueryHandler) error {
	cmdID := c.QueryID()

	_, ok := cb.handlersMap[cmdID]
	if ok {
		return fmt.Errorf("the Command %s is already registered", cmdID)
	}
	cb.handlersMap[cmdID] = ch
	return nil
}

func (cb QueryBus) Exec(c Query) (QueryResponse, error) {
	cmdID := c.QueryID()

	ch, ok := cb.handlersMap[cmdID]
	if !ok {
		return nil, fmt.Errorf("there not any QueryHandler associate to query %s", cmdID)
	}
	rsp, err := ch.Handle(c)
	return rsp, err
}
