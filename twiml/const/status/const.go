package status

type (
	CallbackEvent string
)

const (
	CallbackInitiated CallbackEvent = "initiated"
	CallbackRinging   CallbackEvent = "ringing"
	CallbackAnswered  CallbackEvent = "answered"
	CallbackCompleted CallbackEvent = "completed"
)
