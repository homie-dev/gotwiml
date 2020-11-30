package status

// CallbackEvent is callback event type
type CallbackEvent string

// CallBackEvent constants
const (
	CallbackInitiated CallbackEvent = "initiated"
	CallbackRinging   CallbackEvent = "ringing"
	CallbackAnswered  CallbackEvent = "answered"
	CallbackCompleted CallbackEvent = "completed"

	CallbackStart CallbackEvent = "start"
	CallbackEnd CallbackEvent = "end"
	CallbackJoin CallbackEvent = "join"
	CallbackLeave CallbackEvent = "leave"
	CallbackMute CallbackEvent = "mute"
	CallbackHold CallbackEvent = "hold"
	CallbackSpeaker CallbackEvent = "speaker"
)
