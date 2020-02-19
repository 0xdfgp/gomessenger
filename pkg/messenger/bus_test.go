package messenger_test

import (
	"testing"

	"github.com/isd4n/gomessenger/pkg/messenger"
)

func TestBus_AddMiddleware(t *testing.T) {
	t.Run("It should add a testMiddleware and sets it as next from the existing element", func(t *testing.T) {
		bus := messenger.DefaultBus()
		m1 := &middlewareTest{}
		m2 := &middlewareTest{}
		m3 := &middlewareTest{}

		bus.AddMiddleware(m1)
		if bus.Middleware == nil {
			t.Error("Middleware should be added")
		}

		bus.AddMiddleware(m2)
		if (*m1).Next() == nil {
			t.Error("Middleware should be set as next")
		}

		bus.AddMiddleware(m3)
		if (*m2).Next() == nil {
			t.Error("Middleware should be set as next")
		}

		if (*m3).Next() != nil {
			t.Error("The last middleware shouldn't have next")
		}
	})
}

func TestBus_Dispatch(t *testing.T) {
	bus := messenger.Bus{}
	bus.AddHandler(successVoidHandler{})
	bus.AddHandler(errorVoidHandler{})
	bus.AddHandler(successHandler{})
	bus.AddHandler(errorHandler{})

	t.Run("It should run the void handler without error", func(t *testing.T) {
		command := successVoidHandler{}
		envelope := bus.DispatchWithId("my-id", command)
		expectedError := successVoidHandler{}.Handle(command)

		if envelope.Id != "my-id" {
			t.Error("Id has not been propagated")
		}
		if envelope.LastError != expectedError {
			t.Error("The error wasn't expected")
		}
		if envelope.LastResult != nil {
			t.Error("The result wasn't expected")
		}
	})

	t.Run("It should run the void handler with error", func(t *testing.T) {
		command := errorVoidCommand{}
		envelope := bus.Dispatch(command)

		if envelope.Id == "" {
			t.Error("Id has not been set")
		}
		if envelope.LastError == nil {
			t.Error("The error was expected")
		}
		if envelope.LastResult != nil {
			t.Error("The result wasn't expected")
		}
	})

	t.Run("It should run the handler and return the result", func(t *testing.T) {
		command := successCommand{}
		envelope := bus.Dispatch(command)
		expectedResult, _ := successHandler{}.Handle(command)

		if envelope.Id == "" {
			t.Error("Id has not been propagated")
		}
		if envelope.LastError != nil {
			t.Error("Unexpected error from handler")
		}
		if envelope.LastResult != expectedResult {
			t.Error("Unexpected result from handler")
		}
	})

	t.Run("It should run the handler and return the error", func(t *testing.T) {
		command := errorCommand{}
		envelope := bus.Dispatch(command)
		expectedResult, _ := errorHandler{}.Handle(command)

		if envelope.Id == "" {
			t.Error("Id has not been propagated")
		}
		if envelope.LastError == nil {
			t.Error("Unexpected error from handler")
		}
		if envelope.LastResult != expectedResult {
			t.Error("Unexpected result from handler")
		}
	})

	t.Run("It should pass trough the middleware", func(t *testing.T) {
		middleware := middlewareTest{}
		bus.AddMiddleware(&middleware)
		envelope := bus.DispatchWithId("my id", successVoidCommand{})

		if envelope.Id != "my id modified" {
			t.Error("Id has not been propagated")
		}
	})
}
