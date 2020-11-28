package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// VoiceResponse is <Response> XML for voice
	VoiceResponse interface {
		Say(message string, attrs ...attr.Option) VoiceResponse
		Dial(number string, attrs ...attr.Option) VoiceResponse
		NestDial(attrs ...attr.Option) Dial
		core.XMLer
		core.EmbedXMLer
	}
	voiceResponse struct {
		*core.XML
	}
)

// NewVoiceResponse creates a new voice response instance
func NewVoiceResponse() VoiceResponse {
	return &voiceResponse{XML: core.NewCoreXML(tagRoot)}
}

func (v *voiceResponse) GetEmbedXML() core.XMLer {
	return v.XML
}

// Say appends a <Say> element
func (v *voiceResponse) Say(message string, options ...attr.Option) VoiceResponse {
	t := core.NewXML(verbSay).SetText(message)
	for _, o := range options {
		o(t)
	}
	v.Append(t)
	return v
}

// Dial creates a <Dial> element
func (v *voiceResponse) Dial(number string, options ...attr.Option) VoiceResponse {
	t := core.NewXML(verbDial).SetText(number)
	for _, o := range options {
		o(t)
	}
	v.Append(t)
	return v
}
func (v *voiceResponse) NestDial(attrs ...attr.Option) Dial {
	t := NewDial(attrs...)
	v.Append(t)
	return t
}
