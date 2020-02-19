package main

import (
	"fmt"

	"github.com/isd4n/gomessenger/pkg/messenger"

	"github.com/isd4n/gomessenger/pkg/errors"
	"github.com/isd4n/gomessenger/pkg/middleware"
)

func main() {
	// We create a new bus
	commandBus := messenger.DefaultBus()

	// Adding a middleware to catch errors and send them to sentry
	commandBus.AddMiddleware(middleware.NewSentry("local", "dns-url"))

	// Adding new message handlers
	commandBus.AddHandler(CreateUserHandler{})
	commandBus.AddHandler(GetUserHandler{})

	// Dispatching the first message without expecting any result
	createUserCmd := CreateUser{Id: "123abc"}
	env := commandBus.Dispatch(createUserCmd)
	if env.LastError != nil {
		_ = fmt.Errorf("error on get user: %v", env.LastError)
	}

	// Dispatching a message with an expected result
	env = commandBus.Dispatch(GetUser{Id: "123abc"})
	if env.LastResult != nil {
		fmt.Printf("Output received: %v", env.LastResult)
	}
}

func (g CreateUserHandler) Handle(msg interface{}) error {
	return UserNotCreated{errors.New("User not created")}
}

type CreateUser struct {
	Id string
}

type CreateUserHandler struct {
}

type GetUser struct {
	Id string
}

type GetUserHandler struct {
}

func (g GetUserHandler) Command() interface{} {
	return GetUser{}
}

func (g GetUserHandler) Handle(msg interface{}) (interface{}, error) {
	query := msg.(GetUser)

	user := struct{ Id string }{Id: query.Id}

	return user, nil
}

func (g CreateUserHandler) Command() interface{} {
	return CreateUser{}
}

type UserNotCreated struct {
	errors.Base
}
