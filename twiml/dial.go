package twiml

import (
	"encoding/xml"
	"strconv"
)

type (
	RecordType                       string
	RecordingStatusCallbackEventType string
	RingToneType                     string
	TrimType                         string

	DialOption   func(t TwiML)
	ClientOption func(t TwiML)

	Dial struct {
		XMLName xml.Name `xml:"Dial"`
		TwiML
	}
)

const (
	DoNotRecord           RecordType = "do-not-record"
	RecordFromAnswer      RecordType = "record-from-answer"
	RecordFromRinging     RecordType = "record-from-ringing"
	RecordFromAnswerDual  RecordType = "record-from-answer-dual"
	RecordFromRingingDual RecordType = "record-from-ringing-dual"

	RecordingStatusCallbackEventInProgress RecordingStatusCallbackEventType = "in-progress"
	RecordingStatusCallbackEventCompleted  RecordingStatusCallbackEventType = "completed"
	RecordingStatusCallbackEventAbsent     RecordingStatusCallbackEventType = "absent"

	RingToneAt    = RingToneType("at")
	RingToneAu    = RingToneType("au")
	RingToneBg    = RingToneType("bg")
	RingToneBr    = RingToneType("br")
	RingToneBe    = RingToneType("be")
	RingToneCh    = RingToneType("ch")
	RingToneCl    = RingToneType("cl")
	RingToneCn    = RingToneType("cn")
	RingToneCz    = RingToneType("cz")
	RingToneDe    = RingToneType("de")
	RingToneDk    = RingToneType("dk")
	RingToneEe    = RingToneType("ee")
	RingToneEs    = RingToneType("es")
	RingToneFi    = RingToneType("fi")
	RingToneFr    = RingToneType("fr")
	RingToneGr    = RingToneType("gr")
	RingToneHu    = RingToneType("hu")
	RingToneIl    = RingToneType("il")
	RingToneIn    = RingToneType("in")
	RingToneIt    = RingToneType("it")
	RingToneLt    = RingToneType("lt")
	RingToneJp    = RingToneType("jp")
	RingToneMx    = RingToneType("mx")
	RingToneMy    = RingToneType("my")
	RingToneNl    = RingToneType("nl")
	RingToneNo    = RingToneType("no")
	RingToneNz    = RingToneType("nz")
	RingTonePh    = RingToneType("ph")
	RingTonePl    = RingToneType("pl")
	RingTonePt    = RingToneType("pt")
	RingToneRu    = RingToneType("ru")
	RingToneSe    = RingToneType("se")
	RingToneSg    = RingToneType("sg")
	RingToneTh    = RingToneType("th")
	RingToneUk    = RingToneType("uk")
	RingToneUs    = RingToneType("us")
	RingToneUsOld = RingToneType("us-old")
	RingToneTw    = RingToneType("tw")
	RingToneVe    = RingToneType("ve")
	RingToneZa    = RingToneType("za")

	TrimSilence    = TrimType("trim-silence")
	TrimNotSilence = TrimType("do-not-trim")
)

// number is phone number
func Number(n string) DialOption {
	return func(t TwiML) {
		n := New("Number").SetText(n)
		t.Append(n)
	}
}

// Record the call
func Record(r RecordType) DialOption {
	return func(t TwiML) {
		t.SetOption("record", string(r))
	}
}

// Action is action url
func Action(url string) DialOption {
	return func(t TwiML) {
		t.SetOption("action", url)
	}
}

// AnswerOnBridge preserve the ringing behavior of the inbound call until the Dialed call picks up
func AnswerOnBridge(b bool) DialOption {
	return func(t TwiML) {
		t.SetOption("answerOnBridge", strconv.FormatBool(b))
	}
}

// HangupOnStar Hangup call on star press
func HangupOnStar(b bool) DialOption {
	return func(t TwiML) {
		t.SetOption("hangupOnStar", strconv.FormatBool(b))
	}
}

// CallerID is Caller ID to display
func CallerID(id string) DialOption {
	return func(t TwiML) {
		t.SetOption("callerId", id)
	}
}

// Method is http method with request action url
func Method(method string) DialOption {
	return func(t TwiML) {
		t.SetOption("method", method)
	}
}

// RecordingStatusCallback is Recording status callback URL
func RecordingStatusCallback(url string) DialOption {
	return func(t TwiML) {
		t.SetOption("recordingStatusCallback", url)
	}
}

// RecordingStatusCallbackMethod is Recording status callback URL method
func RecordingStatusCallbackMethod(method string) DialOption {
	return func(t TwiML) {
		t.SetOption("recordingStatusCallbackMethod", method)
	}
}

// RecordingStatusCallbackEvent is Recording status callback URL method
func RecordingStatusCallbackEvent(method string) DialOption {
	return func(t TwiML) {
		t.SetOption("recordingStatusCallbackEvent", method)
	}
}

// RingTone allows you to override the ringback tone that Twilio will play back to the caller while executing the Dial
func RingTone(r RingToneType) DialOption {
	return func(t TwiML) {
		t.SetOption("ringTone", string(r))
	}
}

// TimeLimit is Max time length
func TimeLimit(sec int) DialOption {
	return func(t TwiML) {
		t.SetOption("timeLimit", strconv.Itoa(sec))
	}
}

// Timeout is Time to wait for answer
func Timeout(sec int) DialOption {
	return func(t TwiML) {
		t.SetOption("timeout", strconv.Itoa(sec))
	}
}

// Trim is Time to wait for answer
func Trim(tt TrimType) DialOption {
	return func(t TwiML) {
		t.SetOption("trim", string(tt))
	}
}

// NewDial creates a <Dial> element
func NewDial(options ...DialOption) *Dial {
	t := New("Dial")
	for _, o := range options {
		o(t)
	}
	return &Dial{TwiML: t}
}

// Client appends <Client> element
func (d *Dial) Client(options ...ClientOption) *Dial {
	client := New("Client")
	for _, o := range options {
		o(client)
	}
	d.Append(client)
	return d
}

func (d *Dial) ClientURL(url string) ClientOption {
	return func(t TwiML) {
		t.SetOption("url", url)
	}
}
