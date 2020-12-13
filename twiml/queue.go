package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Queue is <Queue> twiml verb
	Queue interface {
		core.XMLer
		core.EmbedXMLer
	}

	queue struct {
		*core.XML
	}
)

// NewQueue creates a <Queue> element
func NewQueue(queueName string, options ...attr.Option) Queue {
	t := core.NewCoreXML(tagQueue)
	t.SetText(queueName)
	for _, o := range options {
		o(t)
	}
	return &queue{XML: t}
}

// GetEmbedXML returns embed xml
func (c *queue) GetEmbedXML() core.XMLer {
	return c.XML
}
