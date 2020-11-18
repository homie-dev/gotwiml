package twiml

import (
	"encoding/xml"
)

type (
	twiml struct {
		XMLName xml.Name
		Text string `xml:",chardata"`
		Attrs []xml.Attr `xml:",attr"`
		Verbs []twiml
	}

	option func(t *twiml)
)

const rootTagName = "Response"

func newTwiML(tagName string, options ...option) twiml {
	t := twiml{
		XMLName: xml.Name{Local: tagName},
	}
	for _, fn := range options {
		fn(&t)
	}
	return t
}

func (t *twiml) append(v twiml) {
	t.Verbs = append(t.Verbs, v)
}

func (t *twiml) nest(v twiml) twiml {
	t.append(v)
	return v
}

// ToXML returns XML String with XML declaration
func (t *twiml) ToXML() (string, error) {
	x, err := t.String()
	return xml.Header + x, err
}

// String returns XML String
func (t *twiml) String() (string, error) {
	//s, err := t.Marshal()
	s, err := t.MarshalIndent("", " ")
	return string(s), err
}

// Marshal returns XML Byte Slice
func (t *twiml) Marshal() ([]byte, error) {
	return xml.Marshal(t)
}

// MarshalIndent returns XML Byte Slice with prefix and indent
func (t *twiml) MarshalIndent(prefix string, indent string) ([]byte, error) {
	return xml.MarshalIndent(t, prefix, indent)
}

func text(v string) option {
	return func(t *twiml) {
		t.Text = v
	}
}

func attr(k string, v string) option {
	return func(t *twiml) {
		t.Attrs = append(t.Attrs, xml.Attr{
			Name: xml.Name{Local: k}, Value: v,
		})
	}
}

