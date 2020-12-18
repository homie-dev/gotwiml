package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// SMS is <Sms> twiml verb
	SMS interface {
		core.XMLer
		core.EmbedXMLer
	}

	sms struct {
		*core.XML
	}
)

// NewSMS creates <Sms> element
func NewSMS(message string, options ...attr.Option) SMS {
	c := core.NewCoreXML(tagSMS)
	c.SetText(message)
	for _, o := range options {
		o(c)
	}
	return &sms{XML: c}
}

// GetEmbedXML returns embed xml
func (e *sms) GetEmbedXML() core.XMLer {
	return e.XML
}
