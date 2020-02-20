package messenger

import (
	"github.com/isd4n/gomessenger/pkg/uuid"
)

// Generates a default bus.
func DefaultBus() Bus {
	return Bus{
		CommandHandlers: NewMessageHandlers(),
	}
}

// Bus is used as transport for messages
type Bus struct {
	Middleware      *Middleware
	CommandHandlers MessageHandlers
}

// Adds a new middleware setting it as "next" from the existing middleware, if it exists.
func (b *Bus) AddMiddleware(m Middleware) {
	if b.Middleware == nil {
		b.Middleware = &m
		return
	}

	(*b.Middleware).Last().SetNext(m)
}

// Handlers are the final receiver of a message and it's used to place the application logic.
func (b *Bus) AddHandler(h Handler) {
	b.CommandHandlers.Add(h)
}

// Dispatches a message. It generates a new trace id.
func (b Bus) Dispatch(msg interface{}) Envelope {
	return b.DispatchWithId(uuid.New(), msg)
}

// Dispatches a message. It's used if you want to mark with an existing trace id.
func (b Bus) DispatchWithId(id string, msg interface{}) Envelope {
	e := newEnvelope(id, msg)

	var tempMiddleware Middleware
	if b.Middleware == nil {
		tempMiddleware = &b.CommandHandlers
	} else {
		tempMiddleware = *b.Middleware
		tempMiddleware.SetNext(&b.CommandHandlers)
	}

	e = tempMiddleware.Handle(e)

	return e
}
