package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Number is <Number> twiml verb
	Number interface {
		core.XMLer
		core.EmbedXMLer
	}

	number struct {
		*core.XML
	}
)

// NewNumber creates a <Number> element
func NewNumber(phoneNumber string, options ...attr.Option) Number {
	t := core.NewCoreXML(tagNumber)
	t.SetText(phoneNumber)
	for _, o := range options {
		o(t)
	}
	return &number{XML: t}
}

// GetEmbedXML returns embed xml
func (c *number) GetEmbedXML() core.XMLer {
	return c.XML
}
