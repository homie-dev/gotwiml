package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Sip is <Sip> twiml verb
	Sip interface {
		core.XMLer
		core.EmbedXMLer
	}

	sip struct {
		*core.XML
	}
)

// NewSip creates a <Sip> element
func NewSip(uri string, options ...attr.Option) Sip {
	t := core.NewCoreXML(tagSip)
	t.SetText(uri)
	for _, o := range options {
		o(t)
	}
	return &sip{XML: t}
}

// GetEmbedXML returns embed xml
func (c *sip) GetEmbedXML() core.XMLer {
	return c.XML
}
