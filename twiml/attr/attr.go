package attr

import (
	"strconv"

	"github.com/homie-dev/gotwiml/twiml/const/http"
	"github.com/homie-dev/gotwiml/twiml/const/say"
	"github.com/homie-dev/gotwiml/twiml/const/status"

	"github.com/homie-dev/gotwiml/twiml/core"

	"github.com/homie-dev/gotwiml/twiml/const/dial"
)

// Option is function to set XML attribute
type Option func(core.XMLer)

const (
	participantIdentity           = "participantIdentity"
	voice                         = "voice"
	language                      = "language"
	loop                          = "loop"
	record                        = "record"
	action                        = "action"
	answerOnBridge                = "answerOnBridge"
	hangupOnStar                  = "hangupOnStar"
	callerID                      = "callerId"
	method                        = "method"
	recordingStatusCallback       = "recordingStatusCallback"
	recordingStatusCallbackMethod = "recordingStatusCallbackMethod"
	recordingStatusCallbackEvent  = "recordingStatusCallbackEvent"
	ringTone                      = "ringTone"
	timeLimit                     = "timeLimit"
	timeout                       = "timeout"
	trim                          = "trim"
	url                           = "url"
	statusCallback                = "statusCallback"
	statusCallbackMethod          = "statusCallbackMethod"
	statusCallbackEvent           = "statusCallbackEvent"
)

// Voice sets voice to use
func Voice(v say.Voice) Option {
	return func(t core.XMLer) {
		t.SetAttr(voice, string(v))
	}
}

// Language sets message language
func Language(v say.Language) Option {
	return func(t core.XMLer) {
		t.SetAttr(language, string(v))
	}
}

// Loop sets times to loop message
func Loop(v say.Language) Option {
	return func(t core.XMLer) {
		t.SetAttr(loop, string(v))
	}
}

// Record sets record type of the call
func Record(r dial.RecordType) Option {
	return func(t core.XMLer) {
		t.SetAttr(record, string(r))
	}
}

// Action is action url
func Action(url string) Option {
	return func(t core.XMLer) {
		t.SetAttr(action, url)
	}
}

// AnswerOnBridge preserve the ringing behavior of the inbound call until the Dialed call picks up
func AnswerOnBridge(b bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(answerOnBridge, strconv.FormatBool(b))
	}
}

// HangupOnStar Hangup call on star press
func HangupOnStar(b bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(hangupOnStar, strconv.FormatBool(b))
	}
}

// CallerID is Caller ID to display
func CallerID(id string) Option {
	return func(t core.XMLer) {
		t.SetAttr(callerID, id)
	}
}

// Method is http method with request action url
func Method(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(method, string(v))
	}
}

// RecordingStatusCallback is Recording status callback URL
func RecordingStatusCallback(url string) Option {
	return func(t core.XMLer) {
		t.SetAttr(recordingStatusCallback, url)
	}
}

// RecordingStatusCallbackMethod is Recording status callback URL method
func RecordingStatusCallbackMethod(method http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(recordingStatusCallbackMethod, string(method))
	}
}

// RecordingStatusCallbackEvent is Recording status callback URL method
func RecordingStatusCallbackEvent(method string) Option {
	return func(t core.XMLer) {
		t.SetAttr(recordingStatusCallbackEvent, method)
	}
}

// RingTone allows you to override the ringback tone that Twilio will play back to the caller while executing the Dial
func RingTone(r dial.RingToneType) Option {
	return func(t core.XMLer) {
		t.SetAttr(ringTone, string(r))
	}
}

// TimeLimit is Max time length
func TimeLimit(sec int) Option {
	return func(t core.XMLer) {
		t.SetAttr(timeLimit, strconv.Itoa(sec))
	}
}

// Timeout is Time to wait for answer
func Timeout(sec int) Option {
	return func(t core.XMLer) {
		t.SetAttr(timeout, strconv.Itoa(sec))
	}
}

// Trim is Time to wait for answer
func Trim(tt dial.TrimType) Option {
	return func(t core.XMLer) {
		t.SetAttr(trim, string(tt))
	}
}

// URL sets url
func URL(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(url, v)
	}
}

// StatusCallback sets status call back url
func StatusCallback(url string) Option {
	return func(t core.XMLer) {
		t.SetAttr(statusCallback, url)
	}
}

// StatusCallbackEvent sets status call back url
func StatusCallbackEvent(event status.CallbackEvent) Option {
	return func(t core.XMLer) {
		t.SetAttr(statusCallbackEvent, string(event))
	}
}

// StatusCallbackMethod sets status call back method
func StatusCallbackMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(statusCallbackMethod, string(v))
	}
}

// ParticipantIdentity sets a unique identity on the incoming caller
func ParticipantIdentity(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(participantIdentity, v)
	}
}
