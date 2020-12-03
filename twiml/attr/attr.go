package attr

import (
	"strconv"
	"strings"

	"github.com/homie-dev/gotwiml/twiml/attr/const/record"
	"github.com/homie-dev/gotwiml/twiml/attr/const/region"
	"github.com/homie-dev/gotwiml/twiml/attr/const/ring"

	"github.com/homie-dev/gotwiml/twiml/attr/const/beep"
	"github.com/homie-dev/gotwiml/twiml/attr/const/http"
	"github.com/homie-dev/gotwiml/twiml/attr/const/jitter"
	"github.com/homie-dev/gotwiml/twiml/attr/const/say"
	"github.com/homie-dev/gotwiml/twiml/attr/const/status"
	"github.com/homie-dev/gotwiml/twiml/attr/const/trim"
	"github.com/homie-dev/gotwiml/twiml/core"
)

// Option is function to set XML attribute
type Option func(core.XMLer)

const (
	_action                        = "action"
	_answerOnBridge                = "answerOnBridge"
	_beep                          = "beep"
	_byoc                          = "byoc"
	_callerID                      = "callerId"
	_coach                         = "coach"
	_endConferenceOnExit           = "endConferenceOnExit"
	_hangupOnStar                  = "hangupOnStar"
	_jitterBufferSize              = "jitterBufferSize"
	_language                      = "language"
	_loop                          = "loop"
	_maxParticipants               = "maxParticipants"
	_method                        = "method"
	_muted                         = "muted"
	_name                          = "name"
	_participantIdentity           = "participantIdentity"
	_participantLabel              = "participantLabel"
	_postWorkActivitySID           = "postWorkActivitySid"
	_record                        = "record"
	_recordingStatusCallback       = "recordingStatusCallback"
	_recordingStatusCallbackEvent  = "recordingStatusCallbackEvent"
	_recordingStatusCallbackMethod = "recordingStatusCallbackMethod"
	_region                        = "region"
	_reservationSID                = "reservationSid"
	_ringTone                      = "ringTone"
	_sendDigits                    = "sendDigits"
	_startConferenceOnEnter        = "startConferenceOnEnter"
	_statusCallback                = "statusCallback"
	_statusCallbackEvent           = "statusCallbackEvent"
	_statusCallbackMethod          = "statusCallbackMethod"
	_timeLimit                     = "timeLimit"
	_timeout                       = "timeout"
	_trim                          = "trim"
	_url                           = "url"
	_value                         = "value"
	_voice                         = "voice"
	_waitMethod                    = "waitMethod"
	_waitURL                       = "waitUrl"
	_waitURLMethod                 = "waitUrlMethod"
	_workflowSID                   = "workflowSid"
)

// Action is action url
func Action(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_action, v)
	}
}

// AnswerOnBridge preserve the ringing behavior of the inbound call until the Dialed call picks up
func AnswerOnBridge(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_answerOnBridge, strconv.FormatBool(v))
	}
}

// Beep sets to Play beep when joining
func Beep(v beep.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_beep, string(v))
	}
}

// BYOC sets byoc trunk SID
func BYOC(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_byoc, v)
	}
}

// CallerID is Caller ID to display
func CallerID(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_callerID, v)
	}
}

// Coach sets to accept a call SID
func Coach(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_coach, v)
	}
}

// EndConferenceOnExit sets to accept a call SID
func EndConferenceOnExit(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_endConferenceOnExit, strconv.FormatBool(v))
	}
}

// HangupOnStar Hangup call on star press
func HangupOnStar(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_hangupOnStar, strconv.FormatBool(v))
	}
}

// JitterBufferSize sets  jitter buffer behavior for a conference participant
func JitterBufferSize(v jitter.BufferSize) Option {
	return func(t core.XMLer) {
		t.SetAttr(_jitterBufferSize, string(v))
	}
}

// Language sets message language
func Language(v say.Language) Option {
	return func(t core.XMLer) {
		t.SetAttr(_language, string(v))
	}
}

// Loop sets times to loop message
func Loop(v say.Language) Option {
	return func(t core.XMLer) {
		t.SetAttr(_loop, string(v))
	}
}

// MaxParticipants sets maximum number of participants
func MaxParticipants(v int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_maxParticipants, strconv.Itoa(v))
	}
}

// Method sets http method with request action url
func Method(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_method, string(v))
	}
}

// Muted sets to join the conference muted
func Muted(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_muted, strconv.FormatBool(v))
	}
}

// Name sets the name of the custom parameter
func Name(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_name, v)
	}
}

// PostWorkActivitySID sets TaskRouter Activity SID
func PostWorkActivitySID(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_postWorkActivitySID, v)
	}
}

// ParticipantIdentity sets a unique identity on the incoming caller
func ParticipantIdentity(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_participantIdentity, v)
	}
}

// ParticipantLabel sets a label for the conference participant
func ParticipantLabel(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_participantLabel, v)
	}
}

// Record sets record type of the call
func Record(r record.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_record, string(r))
	}
}

// RecordingStatusCallback is Recording status callback URL
func RecordingStatusCallback(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_recordingStatusCallback, v)
	}
}

// RecordingStatusCallbackEvent is Recording status callback URL method
func RecordingStatusCallbackEvent(v record.StatusCallbackEventType) Option {
	return func(t core.XMLer) {
		t.SetAttr(_recordingStatusCallbackEvent, string(v))
	}
}

// RecordingStatusCallbackMethod is Recording status callback URL method
func RecordingStatusCallbackMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_recordingStatusCallbackMethod, string(v))
	}
}

// Region sets the region where Twilio should mix the conference
func Region(v region.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_region, string(v))
	}
}

// ReservationSID sets TaskRouter Reservation SID
func ReservationSID(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_reservationSID, v)
	}
}

// RingTone sets to allow you to override the ringback tone that Twilio will play back to the caller while executing the Dial
func RingTone(v ring.ToneType) Option {
	return func(t core.XMLer) {
		t.SetAttr(_ringTone, string(v))
	}
}

// SendDigits sets DTMF tones to play when the call is answered
func SendDigits(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_sendDigits, v)
	}
}

// StartConferenceOnEnter sets to start the conference on enter
func StartConferenceOnEnter(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_startConferenceOnEnter, strconv.FormatBool(v))
	}
}

// StatusCallback sets status call back url
func StatusCallback(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_statusCallback, v)
	}
}

// StatusCallbackEvent sets status call back url
func StatusCallbackEvent(vv ...status.CallbackEvent) Option {
	ss := make([]string, len(vv))
	for i, v := range vv {
		ss[i] = string(v)
	}
	return func(t core.XMLer) {
		t.SetAttr(_statusCallbackEvent, strings.Join(ss, " "))
	}
}

// StatusCallbackMethod sets status call back method
func StatusCallbackMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_statusCallbackMethod, string(v))
	}
}

// TimeLimit is Max time length
func TimeLimit(sec int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_timeLimit, strconv.Itoa(sec))
	}
}

// Timeout is Time to wait for answer
func Timeout(sec int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_timeout, strconv.Itoa(sec))
	}
}

// Trim is Time to wait for answer
func Trim(v trim.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_trim, string(v))
	}
}

// URL sets url
func URL(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_url, v)
	}
}

// Value sets the value of the custom parameter
func Value(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_value, v)
	}
}

// Voice sets voice to use
func Voice(v say.Voice) Option {
	return func(t core.XMLer) {
		t.SetAttr(_voice, string(v))
	}
}

// WaitMethod sets wait URL method
func WaitMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_waitMethod, string(v))
	}
}

// WaitURL sets wait URL
func WaitURL(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_waitURL, string(v))
	}
}

// WaitURLMethod sets wait URL method
func WaitURLMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_waitURLMethod, string(v))
	}
}

// WorkflowSID sets wait URL method
func WorkflowSID(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_workflowSID, v)
	}
}
