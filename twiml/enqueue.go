package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Enqueue is <Enqueue> twiml verb
	Enqueue interface {
		core.XMLer
		core.EmbedXMLer
	}

	enqueue struct {
		*core.XML
	}
)

// NewEnqueue creates a <Enqueue> element
func NewEnqueue(queueName string, options ...attr.Option) Enqueue {
	t := core.NewCoreXML(tagEnqueue)
	t.SetText(queueName)
	for _, o := range options {
		o(t)
	}
	return &enqueue{XML: t}
}

// GetEmbedXML returns embed xml
func (c *enqueue) GetEmbedXML() core.XMLer {
	return c.XML
}
