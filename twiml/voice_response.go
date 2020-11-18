package twiml

import (
	"encoding/xml"
)

type (
	// VoiceResponse is <Response> TwiML for voice
	VoiceResponse struct {
		XMLName xml.Name `xml:"Response"`
		TwiML
	}
)

// NewVoiceResponse creates a new voice response instance
func NewVoiceResponse() *VoiceResponse {
	return &VoiceResponse{
		TwiML: New("Response"),
	}
}

// Say appends a <Say> element
func (v *VoiceResponse) Say(message string, options ...SayOption) *VoiceResponse {
	t := New("Say").SetText(message)
	for _, o := range options {
		o(t)
	}
	v.Append(t)
	return v
}

// Dial creates a <Dial> element
func (v *VoiceResponse) Dial(options ...DialOption) *VoiceResponse {
	t := New("Dial")
	for _, o := range options {
		o(t)
	}
	v.Append(t)
	return v
}
