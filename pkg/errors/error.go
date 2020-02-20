package errors

import "runtime"

// Creates a new error that includes the stacktrace.
func New(msg string) Base {
	return Base{
		Message: msg,
		Stack:   generateStackTrace(),
	}
}

type Base struct {
	Message string
	Stack   []uintptr
}

func (u Base) Error() string {
	return u.Message
}

func (u Base) StackTrace() []uintptr {
	return u.Stack
}

func generateStackTrace() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[0:n]
}
