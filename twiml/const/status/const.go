package status

// CallbackEvent is callback event type
type CallbackEvent string

// CallBackEvent constants
const (
	CallbackInitiated CallbackEvent = "initiated"
	CallbackRinging   CallbackEvent = "ringing"
	CallbackAnswered  CallbackEvent = "answered"
	CallbackCompleted CallbackEvent = "completed"
)
