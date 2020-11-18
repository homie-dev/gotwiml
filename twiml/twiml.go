package twiml

import (
	"encoding/xml"
)

type (
	TwiML interface {
		Append(TwiML)
		Nest(TwiML) TwiML
		ToXML() (string, error)
		SetText(text string) TwiML
		SetOption(key, value string) TwiML
	}
	// TwiML is base struct
	twiML struct {
		XMLName xml.Name
		Text    string     `xml:",chardata"`
		Attrs   []xml.Attr `xml:",attr"`
		Verbs   []TwiML
	}
)

// New creates a new TwiML instance
func New(tagName string) TwiML {
	t := &twiML{
		XMLName: xml.Name{Local: tagName},
	}
	return t
}

// Append a child element to this element
func (t *twiML) Append(v TwiML) {
	t.Verbs = append(t.Verbs, v)
}

// Nest a child element to this element and returning the newly created element
func (t *twiML) Nest(v TwiML) TwiML {
	t.Append(v)
	return v
}

// ToXML returns XML String with XML declaration
func (t *twiML) ToXML() (string, error) {
	x, err := t.String()
	return xml.Header + x, err
}

// String returns XML String
func (t *twiML) String() (string, error) {
	//s, err := t.Marshal()
	s, err := t.MarshalIndent("", " ")
	return string(s), err
}

// Marshal returns XML Byte Slice
func (t *twiML) Marshal() ([]byte, error) {
	return xml.Marshal(t)
}

// MarshalIndent returns XML Byte Slice with prefix and indent
func (t *twiML) MarshalIndent(prefix string, indent string) ([]byte, error) {
	return xml.MarshalIndent(t, prefix, indent)
}

// SetOption adds key-value attributes to the generated xml
func (t *twiML) SetOption(key, value string) TwiML {
	t.Attrs = append(t.Attrs, xml.Attr{
		Name: xml.Name{Local: key}, Value: value,
	})
	return t
}

// SetText sets body text of xml element
func (t *twiML) SetText(text string) TwiML {
	t.Text = text
	return t
}
