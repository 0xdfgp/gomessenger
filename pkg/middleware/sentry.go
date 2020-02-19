package middleware

import (
	"fmt"
	"reflect"
	"time"

	"github.com/isd4n/gomessenger/pkg/messenger"

	"github.com/getsentry/sentry-go"
)

func NewSentry(environment string, dsn string) messenger.Middleware {
	return &Sentry{
		Environment: environment,
		Dsn:         dsn,
	}
}

type Sentry struct {
	messenger.MiddlewareImpl
	Environment string
	Dsn         string
}

func (s Sentry) Handle(env messenger.Envelope) messenger.Envelope {
	defer sentry.Flush(2 * time.Second)

	s.initialize()

	if env.LastError != nil {
		sentry.CaptureException(env.LastError)
		return env
	}

	if s.Next() != nil {
		next := *s.Next()
		env = next.Handle(env)
		if env.LastError != nil {
			fmt.Println(reflect.TypeOf(env.LastError).String())
			sentry.CaptureException(env.LastError)
		}
	}

	return env
}

func (s Sentry) initialize() {
	_ = sentry.Init(sentry.ClientOptions{
		Dsn:   s.Dsn,
		Debug: false,
	})

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag("type", "MyError")
		scope.SetTag("environment", s.Environment)
	})
}
