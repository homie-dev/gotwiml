package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// VoiceResponse is <Response> XML for voice
	VoiceResponse interface {
		AppendConnect(Connect) VoiceResponse
		Dial(number string, attrs ...attr.Option) VoiceResponse
		AppendDial(Dial) VoiceResponse
		Say(message string, attrs ...attr.Option) VoiceResponse
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

// GetEmbedXML returns embed xml
func (v *voiceResponse) GetEmbedXML() core.XMLer {
	return v.XML
}

// AppendConnect appends a <Conenct> element
func (v *voiceResponse) AppendConnect(c Connect) VoiceResponse {
	v.Append(c)
	return v
}

// Say appends a <Say> element and applies attributes
func (v *voiceResponse) Say(message string, options ...attr.Option) VoiceResponse {
	t := core.NewXML(verbSay).SetText(message)
	for _, o := range options {
		o(t)
	}
	v.Append(t)
	return v
}

// Dial appends a <Dial> element and applies attributes
func (v *voiceResponse) Dial(number string, options ...attr.Option) VoiceResponse {
	t := core.NewXML(verbDial).SetText(number)
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
