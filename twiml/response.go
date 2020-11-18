package twiml

import (
	"encoding/xml"
)

type (
	// VoiceResponse is <VoiceResponse> TwiML for voice
	VoiceResponse struct {
		XMLName xml.Name `xml:"VoiceResponse"`
		*TwiML
	}
)

// NewVoiceResponse creates a new voice response instance
func NewVoiceResponse() *VoiceResponse {
	return &VoiceResponse{
		TwiML: New("VoiceResponse"),
	}
}

// Say creates a <Say> element
func (v *VoiceResponse) Say(message string, options ...SayOption) *VoiceResponse {
	t := New("Say").SetText(message)
	for _, o := range options {
		o(t)
	}
	v.Append(t)
	return v
}
