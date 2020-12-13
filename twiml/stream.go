package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Stream is <Stream> twiml verb
	Stream interface {
		core.XMLer
		core.EmbedXMLer
	}

	stream struct {
		*core.XML
	}
)

// NewStream creates a <Stream> element
func NewStream(options ...attr.Option) Stream {
	t := core.NewCoreXML(verbStream)
	for _, o := range options {
		o(t)
	}
	return &stream{XML: t}
}

// GetEmbedXML returns embed xml
func (c *stream) GetEmbedXML() core.XMLer {
	return c.XML
}
