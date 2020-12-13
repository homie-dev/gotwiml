package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Parameter is <Parameter> twiml verb
	Parameter interface {
		core.XMLer
		core.EmbedXMLer
	}

	parameter struct {
		*core.XML
	}
)

// NewParameter creates a <Parameter> element
func NewParameter(options ...attr.Option) Parameter {
	t := core.NewCoreXML(nounParameter)
	for _, o := range options {
		o(t)
	}
	return &parameter{XML: t}
}

// GetEmbedXML returns embed xml
func (e *parameter) GetEmbedXML() core.XMLer {
	return e.XML
}
