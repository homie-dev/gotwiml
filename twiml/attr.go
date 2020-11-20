package twiml

import (
	"strconv"

	"github.com/homie-dev/gotwiml/twiml/const/dial"
)

type Attr func(TwiML)

const (
	AttrVoice                         = "voice"
	AttrLanguage                      = "language"
	AttrLoop                          = "loop"
	AttrRecord                        = "record"
	AttrAction                        = "action"
	AttrAnswerOnBridge                = "answerOnBridge"
	AttrHangupOnStar                  = "hangupOnStar"
	AttrCallerID                      = "callerId"
	AttrMethod                        = "method"
	AttrRecordingStatusCallback       = "recordingStatusCallback"
	AttrRecordingStatusCallbackMethod = "recordingStatusCallbackMethod"
	AttrRecordingStatusCallbackEvent  = "recordingStatusCallbackEvent"
	AttrRingTone                      = "ringTone"
	AttrTimeLimit                     = "timeLimit"
	AttrTimeout                       = "timeout"
	AttrTrim                          = "trim"
	AttrURL                           = "url"
	AttrStatusCallback                = "statusCallback"
	AttrStatusCallbackMethod          = "statusCallbackMethod"
	AttrStatusCallbackEvent           = "statusCallbackEvent"
)

// Record sets record type of the call
func Record(r dial.RecordType) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrRecord, string(r))
	}
}

// Action is action url
func Action(url string) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrAction, url)
	}
}

// AnswerOnBridge preserve the ringing behavior of the inbound call until the Dialed call picks up
func AnswerOnBridge(b bool) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrAnswerOnBridge, strconv.FormatBool(b))
	}
}

// HangupOnStar Hangup call on star press
func HangupOnStar(b bool) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrHangupOnStar, strconv.FormatBool(b))
	}
}

// CallerID is Caller ID to display
func CallerID(id string) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrCallerID, id)
	}
}

// Method is http method with request action url
func Method(method string) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrMethod, method)
	}
}

// RecordingStatusCallback is Recording status callback URL
func RecordingStatusCallback(url string) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrRecordingStatusCallback, url)
	}
}

// RecordingStatusCallbackMethod is Recording status callback URL method
func RecordingStatusCallbackMethod(method string) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrRecordingStatusCallbackMethod, method)
	}
}

// RecordingStatusCallbackEvent is Recording status callback URL method
func RecordingStatusCallbackEvent(method string) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrRecordingStatusCallbackEvent, method)
	}
}

// RingTone allows you to override the ringback tone that Twilio will play back to the caller while executing the Dial
func RingTone(r dial.RingToneType) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrRingTone, string(r))
	}
}

// TimeLimit is Max time length
func TimeLimit(sec int) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrTimeLimit, strconv.Itoa(sec))
	}
}

// Timeout is Time to wait for answer
func Timeout(sec int) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrTimeout, strconv.Itoa(sec))
	}
}

// Trim is Time to wait for answer
func Trim(tt dial.TrimType) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrTrim, string(tt))
	}
}

// URL sets url
func URL(url string) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrURL, url)
	}
}

// StatusCallback sets status call back url
func StatusCallback(url string) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrStatusCallback, url)
	}
}
