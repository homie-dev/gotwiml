package twiml

import (
	"encoding/xml"
)

type (
	// TwiML is base struct
	TwiML struct {
		XMLName xml.Name
		Text    string     `xml:",chardata"`
		Attrs   []xml.Attr `xml:",attr"`
		Verbs   []*TwiML
	}

	option func(t *TwiML)
)

const rootTagName = "VoiceResponse"

// New creates a new TwiML instance
func New(tagName string, options ...option) *TwiML {
	t := &TwiML{
		XMLName: xml.Name{Local: tagName},
	}
	for _, fn := range options {
		fn(t)
	}
	return t
}

// Append a child element to this element
func (t *TwiML) Append(v *TwiML) {
	t.Verbs = append(t.Verbs, v)
}

// Nest a child element to this element and returning the newly created element
func (t *TwiML) Nest(v *TwiML) *TwiML {
	t.Append(v)
	return v
}

// ToXML returns XML String with XML declaration
func (t *TwiML) ToXML() (string, error) {
	x, err := t.String()
	return xml.Header + x, err
}

// String returns XML String
func (t *TwiML) String() (string, error) {
	//s, err := t.Marshal()
	s, err := t.MarshalIndent("", " ")
	return string(s), err
}

// Marshal returns XML Byte Slice
func (t *TwiML) Marshal() ([]byte, error) {
	return xml.Marshal(t)
}

// MarshalIndent returns XML Byte Slice with prefix and indent
func (t *TwiML) MarshalIndent(prefix string, indent string) ([]byte, error) {
	return xml.MarshalIndent(t, prefix, indent)
}

// SetOption adds key-value attributes to the generated xml
func (t *TwiML) SetOption(key, value string) *TwiML {
	t.Attrs = append(t.Attrs, xml.Attr{
		Name: xml.Name{Local: key}, Value: value,
	})
	return t
}

// SetText sets body text of xml element
func (t *TwiML) SetText(text string) *TwiML {
	t.Text = text
	return t
}
