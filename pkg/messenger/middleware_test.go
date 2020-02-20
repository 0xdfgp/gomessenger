package messenger_test

import (
	"testing"

	"github.com/isd4n/gomessenger/pkg/messenger"
)

func TestMiddlewareImpl_SetNext(t *testing.T) {
	t.Run("It should set the next middleware", func(t *testing.T) {
		m1 := middlewareTest{}
		m2 := middlewareTest{}

		m1.SetNext(&m2)

		if m1.Next() == nil {
			t.Error("Error no set next testMiddleware")
		}
	})
}

func TestMiddlewareImpl_Handle(t *testing.T) {
	m := struct {
		messenger.MiddlewareImpl
	}{}

	t.Run("It should panic if the middleware doesn't override the handle method", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Panic expected")
			}
		}()
		m.Handle(messenger.Envelope{})

		t.Error("Panic expected")
	})
}
