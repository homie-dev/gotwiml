package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Start is <Start> twiml verb
	Start interface {
		Siprec(...attr.Option) Start
		Stream(...attr.Option) Start
		core.XMLer
		core.EmbedXMLer
	}

	start struct {
		*core.XML
	}
)

// NewStart creates <Start> element
func NewStart(options ...attr.Option) Start {
	c := core.NewCoreXML(verbStart)
	for _, o := range options {
		o(c)
	}
	return &start{XML: c}
}

// GetEmbedXML returns embed xml
func (e *start) GetEmbedXML() core.XMLer {
	return e.XML
}

func (e *start) Siprec(options ...attr.Option) Start {
	e.Append(NewSiprec(options...))
	return e
}

func (e *start) Stream(options ...attr.Option) Start {
	e.Append(NewStream(options...))
	return e
}
