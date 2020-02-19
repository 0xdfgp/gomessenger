package messenger

import (
	"reflect"
)

// Generates a new Middleware to handle messages
func NewMessageHandlers() MessageHandlers {
	return MessageHandlers{Handlers: map[string]Handler{}}
}

// A message handlers handles messages as a Middleware. It has to be placed at the end of the chain because
// it's not going to send to any other testMiddleware.
type MessageHandlers struct {
	MiddlewareImpl
	Handlers Handlers
}

// Collections of handlers
type Handlers map[string]Handler

// Adds a new handler to the collection.
func (c Handlers) Add(handler Handler) {
	c[reflect.TypeOf(handler.Command()).Name()] = handler
}

// Contract that has to be implemented by an application message handler.
type Handler interface {
	Command() interface{}
}

// Message handler that doesn't return any result.
type VoidHandler interface {
	Handler
	Handle(msg interface{}) error
}

// Message handler that returns a result.
type HandlerWithResult interface {
	Handler
	Handle(msg interface{}) (interface{}, error)
}

// Adds a new handler to the Middleware.
func (c *MessageHandlers) Add(handler Handler) {
	if c.Handlers == nil {
		c.Handlers = map[string]Handler{}
	}

	c.Handlers.Add(handler)
}

// Checks it if the is any handler to the given message and executes it.
func (c MessageHandlers) Handle(env Envelope) Envelope {
	cmdName := reflect.TypeOf(env.Message).Name()
	if _, exists := c.Handlers[cmdName]; !exists {
		return env
	}

	var err error
	var result interface{}

	switch handler := c.Handlers[cmdName].(type) {
	case VoidHandler:
		err = handler.Handle(env.Message)

	case HandlerWithResult:
		result, err = handler.Handle(env.Message)
	}

	env.LastResult = result
	env.LastError = err

	return env
}
