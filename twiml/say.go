package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Say is <Say> twiml verb
	Say interface {
		core.XMLer
		core.EmbedXMLer
	}

	say struct {
		*core.XML
	}
)

// NewSay creates a <Say> element
func NewSay(message string, options ...attr.Option) Say {
	t := core.NewCoreXML(tagSay)
	t.SetText(message)
	for _, o := range options {
		o(t)
	}
	return &say{XML: t}
}

// GetEmbedXML returns embed xml
func (c *say) GetEmbedXML() core.XMLer {
	return c.XML
}
