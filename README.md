[![Build Status](https://travis-ci.com/isd4n/gomessenger.svg?branch=master)](https://travis-ci.com/isd4n/gomessenger)
[![codecov](https://codecov.io/gh/isd4n/gomessenger/branch/master/graph/badge.svg)](https://codecov.io/gh/isd4n/gomessenger)

# GoMessenger

Simple implementation of a message bus to handle commands and queries easily but with the potential of middlewares.

You can go to the examples to see how it works: [examples](examples)

## Installation
```bash
go get github.com/isd4n/gomessenger
```

## Usage
```bash
import (
    "github.com/isd4n/gomessenger/pkg/messenger"
    "github.com/isd4n/gomessenger/pkg/errors"
    "github.com/isd4n/gomessenger/pkg/middleware"
)

func main() {
    // We create a new bus
    commandBus := messenger.DefaultBus()

    // Adding a middleware to catch errors and send them to sentry (optional)
    commandBus.AddMiddleware(middleware.NewSentry("local", "dsn"))

    // Adding new message handler
    commandBus.AddHandler(MyHandler{})

    // Dispatching the first message without expecting any result
    myMessage := MyMessage{Id: "123abc"}
    env := commandBus.Dispatch(myMessage)
    if env.LastError != nil {
        // Error has occurried
    }

    if env.LastResult != nil {
        // Result received
    }
}
```