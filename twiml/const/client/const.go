package client

type (
	StatusCallbackEvent string
)

const (
	StatusCallbackInitiated StatusCallbackEvent = "initiated"
	StatusCallbackRinging   StatusCallbackEvent = "ringing"
	StatusCallbackAnswered  StatusCallbackEvent = "answered"
	StatusCallbackCompleted StatusCallbackEvent = "completed"
)
