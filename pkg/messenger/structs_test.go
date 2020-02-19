package messenger_test

import (
	"github.com/isd4n/gomessenger/pkg/errors"
	"github.com/isd4n/gomessenger/pkg/messenger"
)

type successVoidCommand struct {
}

type successVoidHandler struct {
}

func (successVoidHandler) Command() interface{} {
	return successVoidCommand{}
}

func (successVoidHandler) Handle(msg interface{}) error {
	return nil
}

type errorVoidCommand struct {
}

type errorVoidHandler struct {
}

func (errorVoidHandler) Command() interface{} {
	return errorVoidCommand{}
}

func (errorVoidHandler) Handle(msg interface{}) error {
	return errors.New("Error!!")
}

type successCommand struct {
}

type successHandler struct {
}

func (successHandler) Command() interface{} {
	return successCommand{}
}

func (successHandler) Handle(msg interface{}) (interface{}, error) {
	result := struct{ Id string }{Id: "entity-id"}

	return result, nil
}

type errorCommand struct {
}

type errorHandler struct {
}

func (errorHandler) Command() interface{} {
	return errorCommand{}
}

func (errorHandler) Handle(msg interface{}) (interface{}, error) {
	return nil, errors.New("Error!!!")
}

type middlewareTest struct {
	messenger.MiddlewareImpl
}

func (m middlewareTest) Handle(e messenger.Envelope) messenger.Envelope {
	e.Id = e.Id + " modified"
	return e
}
