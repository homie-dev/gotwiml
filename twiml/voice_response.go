package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// VoiceResponse is <Response> XML for voice
	VoiceResponse interface {
		Connect(options ...attr.Option) VoiceResponse
		AppendConnect(Connect) VoiceResponse
		Dial(number string, options ...attr.Option) VoiceResponse
		AppendDial(Dial) VoiceResponse
		Echo() VoiceResponse
		AppendEcho(Echo) VoiceResponse
		Enqueue(queueName string, options ...attr.Option) VoiceResponse
		AppendEnqueue(Enqueue) VoiceResponse
		Gather(options ...attr.Option) VoiceResponse
		AppendGather(Gather) VoiceResponse
		Hangup() VoiceResponse
		AppendHangup(Hangup) VoiceResponse
		Leave() VoiceResponse
		AppendLeave(Leave) VoiceResponse
		Pause(options ...attr.Option) VoiceResponse
		AppendPause(Pause) VoiceResponse
		Play(url string, options ...attr.Option) VoiceResponse
		AppendPlay(Play) VoiceResponse
		Queue(queueName string, options ...attr.Option) VoiceResponse
		AppendQueue(Queue) VoiceResponse
		Record(options ...attr.Option) VoiceResponse
		AppendRecord(Record) VoiceResponse
		Redirect(url string, options ...attr.Option) VoiceResponse
		AppendRedirect(Redirect) VoiceResponse
		Reject(options ...attr.Option) VoiceResponse
		AppendReject(Reject) VoiceResponse
		Say(message string, options ...attr.Option) VoiceResponse
		AppendSay(Say) VoiceResponse
		Pay(options ...attr.Option) VoiceResponse
		AppendPay(Pay) VoiceResponse
		Prompt(options ...attr.Option) VoiceResponse
		AppendPrompt(Prompt) VoiceResponse
		Start(options ...attr.Option) VoiceResponse
		AppendStart(Start) VoiceResponse
		SMS(message string, options ...attr.Option) VoiceResponse
		AppendSMS(SMS) VoiceResponse
		Stop(options ...attr.Option) VoiceResponse
		AppendStop(Stop) VoiceResponse
		Refer(options ...attr.Option) VoiceResponse
		AppendRefer(Refer) VoiceResponse
		core.XMLer
		core.EmbedXMLer
	}

	voiceResponse struct {
		*core.XML
	}
)

// NewVoiceResponse creates a new voice response instance
func NewVoiceResponse() VoiceResponse {
	return &voiceResponse{XML: core.NewCoreXML(tagResponse)}
}

// GetEmbedXML returns embed xml
func (v *voiceResponse) GetEmbedXML() core.XMLer {
	return v.XML
}

// Connect appends a <Connect> element and applies attributes
func (v *voiceResponse) Connect(options ...attr.Option) VoiceResponse {
	v.Append(NewConnect(options...))
	return v
}

// AppendConnect appends a <Connect> element
func (v *voiceResponse) AppendConnect(c Connect) VoiceResponse {
	v.Append(c)
	return v
}

// Echo appends a <Echo> element and applies attributes
func (v *voiceResponse) Echo() VoiceResponse {
	v.Append(NewEcho())
	return v
}

// AppendEcho appends a <Echo> element
func (v *voiceResponse) AppendEcho(c Echo) VoiceResponse {
	v.Append(c)
	return v
}

// Enqueue appends a <Enqueue> element and applies attributes
func (v *voiceResponse) Enqueue(queueName string, options ...attr.Option) VoiceResponse {
	v.Append(NewEnqueue(queueName, options...))
	return v
}

// AppendEnqueue appends a <Enqueue> element
func (v *voiceResponse) AppendEnqueue(c Enqueue) VoiceResponse {
	v.Append(c)
	return v
}

// Gather appends a <Gather> element and applies attributes
func (v *voiceResponse) Gather(options ...attr.Option) VoiceResponse {
	v.Append(NewGather(options...))
	return v
}

// AppendGather appends a <Gather> element
func (v *voiceResponse) AppendGather(c Gather) VoiceResponse {
	v.Append(c)
	return v
}

// Hangup appends a <Hangup> element and applies attributes
func (v *voiceResponse) Hangup() VoiceResponse {
	v.Append(NewHangup())
	return v
}

// AppendHangup appends a <Hangup> element
func (v *voiceResponse) AppendHangup(c Hangup) VoiceResponse {
	v.Append(c)
	return v
}

// Leave appends a <Leave> element and applies attributes
func (v *voiceResponse) Leave() VoiceResponse {
	v.Append(NewLeave())
	return v
}

// AppendLeave appends a <Leave> element
func (v *voiceResponse) AppendLeave(c Leave) VoiceResponse {
	v.Append(c)
	return v
}

// Pause appends a <Pause> element and applies attributes
func (v *voiceResponse) Pause(options ...attr.Option) VoiceResponse {
	v.Append(NewPause(options...))
	return v
}

// AppendPause appends a <Pause> element
func (v *voiceResponse) AppendPause(c Pause) VoiceResponse {
	v.Append(c)
	return v
}

// Play appends a <Play> element and applies attributes
func (v *voiceResponse) Play(url string, options ...attr.Option) VoiceResponse {
	v.Append(NewPlay(url, options...))
	return v
}

// AppendPlay appends a <Play> element
func (v *voiceResponse) AppendPlay(c Play) VoiceResponse {
	v.Append(c)
	return v
}

// Queue appends a <Queue> element and applies attributes
func (v *voiceResponse) Queue(queueName string, options ...attr.Option) VoiceResponse {
	v.Append(NewQueue(queueName, options...))
	return v
}

// AppendQueue appends a <Queue> element
func (v *voiceResponse) AppendQueue(c Queue) VoiceResponse {
	v.Append(c)
	return v
}

// Record appends a <Record> element and applies attributes
func (v *voiceResponse) Record(options ...attr.Option) VoiceResponse {
	v.Append(NewRecord(options...))
	return v
}

// AppendRecord appends a <Record> element
func (v *voiceResponse) AppendRecord(c Record) VoiceResponse {
	v.Append(c)
	return v
}

// Redirect appends a <Redirect> element and applies attributes
func (v *voiceResponse) Redirect(url string, options ...attr.Option) VoiceResponse {
	v.Append(NewRedirect(url, options...))
	return v
}

// AppendRedirect appends a <Redirect> element
func (v *voiceResponse) AppendRedirect(c Redirect) VoiceResponse {
	v.Append(c)
	return v
}

// Reject appends a <Reject> element and applies attributes
func (v *voiceResponse) Reject(options ...attr.Option) VoiceResponse {
	v.Append(NewReject(options...))
	return v
}

// AppendReject appends a <Reject> element
func (v *voiceResponse) AppendReject(c Reject) VoiceResponse {
	v.Append(c)
	return v
}

// Say appends a <Say> element and applies attributes
func (v *voiceResponse) Say(message string, options ...attr.Option) VoiceResponse {
	v.Append(NewSay(message, options...))
	return v
}

// AppendSay appends a <Say> element
func (v *voiceResponse) AppendSay(c Say) VoiceResponse {
	v.Append(c)
	return v
}

// Pay appends a <Pay> element and applies attributes
func (v *voiceResponse) Pay(options ...attr.Option) VoiceResponse {
	v.Append(NewPay(options...))
	return v
}

// AppendPay appends a <Pay> element
func (v *voiceResponse) AppendPay(c Pay) VoiceResponse {
	v.Append(c)
	return v
}

// Prompt appends a <Prompt> element and applies attributes
func (v *voiceResponse) Prompt(options ...attr.Option) VoiceResponse {
	v.Append(NewPrompt(options...))
	return v
}

// AppendPrompt appends a <Prompt> element
func (v *voiceResponse) AppendPrompt(c Prompt) VoiceResponse {
	v.Append(c)
	return v
}

// Start appends a <Start> element and applies attributes
func (v *voiceResponse) Start(options ...attr.Option) VoiceResponse {
	v.Append(NewStart(options...))
	return v
}

// AppendStart appends a <Start> element
func (v *voiceResponse) AppendStart(c Start) VoiceResponse {
	v.Append(c)
	return v
}

// SMS appends a <Sms> element and applies attributes
func (v *voiceResponse) SMS(message string, options ...attr.Option) VoiceResponse {
	v.Append(NewSMS(message, options...))
	return v
}

// AppendSMS appends a <SMS> element
func (v *voiceResponse) AppendSMS(c SMS) VoiceResponse {
	v.Append(c)
	return v
}

// Stop appends a <Stop> element and applies attributes
func (v *voiceResponse) Stop(options ...attr.Option) VoiceResponse {
	v.Append(NewStop(options...))
	return v
}

// AppendStop appends a <Stop> element
func (v *voiceResponse) AppendStop(c Stop) VoiceResponse {
	v.Append(c)
	return v
}

// Refer appends a <Refer> element and applies attributes
func (v *voiceResponse) Refer(options ...attr.Option) VoiceResponse {
	v.Append(NewRefer(options...))
	return v
}

// AppendRefer appends a <Refer> element
func (v *voiceResponse) AppendRefer(c Refer) VoiceResponse {
	v.Append(c)
	return v
}

// Dial appends a <Dial> element and applies attributes
func (v *voiceResponse) Dial(number string, options ...attr.Option) VoiceResponse {
	t := core.NewXML(tagDial).SetText(number)
	for _, o := range options {
		o(t)
	}
	v.Append(t)
	return v
}

// AppendDial appends a <Dial> element
func (v *voiceResponse) AppendDial(d Dial) VoiceResponse {
	v.Append(d)
	return v
}
