package messenger

type Middleware interface {
	Handle(envelope Envelope) Envelope
	Next() *Middleware
	Last() Middleware
	SetNext(middleware Middleware)
}

type MiddlewareImpl struct {
	next *Middleware
}

func (m MiddlewareImpl) Handle(env Envelope) Envelope {
	panic("handle method not implemented")
}

func (m MiddlewareImpl) Next() *Middleware {
	return m.next
}

func (m *MiddlewareImpl) Last() Middleware {
	var nextMiddleware Middleware = m

	for ok := nextMiddleware.Next() != nil; ok; ok = nextMiddleware.Next() != nil {
		next := nextMiddleware.Next()
		nextMiddleware = *next
	}

	return nextMiddleware
}

func (m *MiddlewareImpl) SetNext(middleware Middleware) {
	m.next = &middleware
}
