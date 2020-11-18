package twiml

import (
	"encoding/xml"
)

type (
	// VoiceResponse is <Response> TwiML for voice
	VoiceResponse struct {
		XMLName 	xml.Name `xml:"Response"`
		twiml
	}
)

// NewVoiceResponse creates a new voice response instance
func NewVoiceResponse() *VoiceResponse {
	return &VoiceResponse{
		twiml: newTwiML(rootTagName),
	}
}

// Say creates a <Say> element
func (v *VoiceResponse) Say(message string) *VoiceResponse {
	t := newTwiML("Say", text(message), attr("key", "value"))
	v.append(t)
	return v
}
