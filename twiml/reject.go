package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Reject is <Reject> twiml verb
	Reject interface {
		core.XMLer
		core.EmbedXMLer
	}

	reject struct {
		*core.XML
	}
)

// NewReject creates a <Reject> element
func NewReject(options ...attr.Option) Reject {
	t := core.NewCoreXML(tagReject)
	for _, o := range options {
		o(t)
	}
	return &reject{XML: t}
}

// GetEmbedXML returns embed xml
func (e *reject) GetEmbedXML() core.XMLer {
	return e.XML
}
