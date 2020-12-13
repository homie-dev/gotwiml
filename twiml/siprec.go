package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Siprec is <Siprec> twiml verb
	Siprec interface {
		Parameter(...attr.Option) Siprec
		core.XMLer
		core.EmbedXMLer
	}

	siprec struct {
		*core.XML
	}
)

// NewSiprec creates <Siprec> element
func NewSiprec(options ...attr.Option) Siprec {
	c := core.NewCoreXML(verbSiprec)
	for _, o := range options {
		o(c)
	}
	return &siprec{XML: c}
}

// GetEmbedXML returns embed xml
func (e *siprec) GetEmbedXML() core.XMLer {
	return e.XML
}

// Parameter appends <Parameter> element with options
func (e *siprec) Parameter(attrs ...attr.Option) Siprec {
	x := core.NewCoreXML(nounParameter)
	for _, a := range attrs {
		a(x)
	}
	e.Append(x)
	return e
}
