package jetstream

import (
	"github.com/nats-io/nats.go"
)

// Router TODO
type Router interface {
	HandleFunc(string, HandlerFunc)
}

// HandlerFunc TODO
type HandlerFunc func()

type router struct {
	nats.JetStream
}

// NewRouter TODO
func NewRouter(
	js nats.JetStream,
) Router {
	return &router{js}
}

// HandleFunc TODO
func (r *router) HandleFunc(
	subject string,
	handle HandlerFunc,
) {
	handle()
}
