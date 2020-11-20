package twiml

type (
	// VoiceResponse is <Response> TwiML for voice
	VoiceResponse interface {
		Say(message string, attrs ...Attr) VoiceResponse
		Dial(number string, attrs ...Attr) VoiceResponse
		NestDial(attrs ...Attr) Dial
		TwiML
		EmbedTwiML
	}
	voiceResponse struct {
		*twiML
	}
)

// NewVoiceResponse creates a new voice response instance
func NewVoiceResponse() VoiceResponse {
	return &voiceResponse{twiML: new(tagRoot)}
}

func (v *voiceResponse) getTwiML() TwiML {
	return v.twiML
}

// Say appends a <Say> element
func (v *voiceResponse) Say(message string, options ...Attr) VoiceResponse {
	t := New(verbSay).SetText(message)
	for _, o := range options {
		o(t)
	}
	v.Append(t)
	return v
}

// Dial creates a <Dial> element
func (v *voiceResponse) Dial(number string, options ...Attr) VoiceResponse {
	t := New(verbDial).SetText(number)
	for _, o := range options {
		o(t)
	}
	v.Append(t)
	return v
}
func (v *voiceResponse) NestDial(attrs ...Attr) Dial {
	t := NewDial(attrs...)
	v.Append(t)
	return t
}
