package messenger

// Generates a new envelope
func newEnvelope(id string, message interface{}) Envelope {
	return Envelope{
		Id:      id,
		Message: message,
	}
}

// An envelope includes the message, result and error if they exist. It responsibility of the testMiddleware to
// fill it correctly.
type Envelope struct {
	Id         string
	Message    interface{}
	LastResult interface{}
	LastError  error
}
