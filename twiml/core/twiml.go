package core

import (
	"encoding/xml"
)

type (
	// XMLer is base interface
	XMLer interface {
		Append(XMLer) XMLer
		Nest(XMLer) XMLer
		Marshal() ([]byte, error)
		String() (string, error)
		ToXML() (string, error)
		MarshalIndent(prefix string, indent string) ([]byte, error)
		StringIndent(prefix string, indent string) (string, error)
		ToXMLIndent(prefix string, indent string) (string, error)
		SetText(text string) XMLer
		SetAttr(key, value string) XMLer
	}

	// EmbedXMLer is Embed Interface
	EmbedXMLer interface {
		GetEmbedXML() XMLer
	}

	// XML is base xml
	XML struct {
		XMLName xml.Name
		Text    string     `xml:",chardata"`
		Attrs   []xml.Attr `xml:",attr"`
		Verbs   []XMLer
	}
)

func NewCoreXML(tagName string) *XML {
	return &XML{
		XMLName: xml.Name{Local: tagName},
	}
}

// New creates a new XMLer instance
func NewXML(tagName string) XMLer {
	return NewCoreXML(tagName)
}

// Append a child element to this element
func (t *XML) Append(v XMLer) XMLer {
	if ex, ok := v.(EmbedXMLer); ok {
		t.Verbs = append(t.Verbs, ex.GetEmbedXML())
	}
	if _, ok := v.(*XML); ok {
		t.Verbs = append(t.Verbs, v)
	}
	return t
}

// Nest a child element to this element and returning the newly created element
func (t *XML) Nest(v XMLer) XMLer {
	t.Append(v)
	return v
}

// ToXML returns XML String with XML declaration
func (t *XML) ToXML() (string, error) {
	x, err := t.String()
	return xml.Header + x, err
}

func (t *XML) ToXMLIndent(prefix string, indent string) (string, error) {
	x, err := t.StringIndent(prefix, indent)
	return xml.Header + x, err
}

// String returns XML String
func (t *XML) String() (string, error) {
	s, err := t.Marshal()
	return string(s), err
}
func (t *XML) StringIndent(prefix string, indent string) (string, error) {
	s, err := t.MarshalIndent(prefix, indent)
	return string(s), err
}

// Marshal returns XML Byte Slice
func (t *XML) Marshal() ([]byte, error) {
	return xml.Marshal(t)
}

// MarshalIndent returns XML Byte Slice with prefix and indent
func (t *XML) MarshalIndent(prefix string, indent string) ([]byte, error) {
	return xml.MarshalIndent(t, prefix, indent)
}

// SetAttr adds key-value attributes to the generated xml
func (t *XML) SetAttr(key, value string) XMLer {
	t.Attrs = append(t.Attrs, xml.Attr{
		Name: xml.Name{Local: key}, Value: value,
	})
	return t
}

// SetText sets body text of xml element
func (t *XML) SetText(text string) XMLer {
	t.Text = text
	return t
}
