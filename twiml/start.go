package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Start is <Start> twiml verb
	Start interface {
		Siprec(...attr.Option) Start
		AppendSiprec(Siprec) Start
		Stream(...attr.Option) Start
		AppendStream(Stream) Start
		core.XMLer
		core.EmbedXMLer
	}

	start struct {
		*core.XML
	}
)

// NewStart creates <Start> element
func NewStart(options ...attr.Option) Start {
	c := core.NewCoreXML(tagStart)
	for _, o := range options {
		o(c)
	}
	return &start{XML: c}
}

// GetEmbedXML returns embed xml
func (e *start) GetEmbedXML() core.XMLer {
	return e.XML
}

// Siprec appends <Siprec> element with optons
func (e *start) Siprec(options ...attr.Option) Start {
	e.Append(NewSiprec(options...))
	return e
}

// AppendSiprec appends <Siprec> element
func (e *start) AppendSiprec(s Siprec) Start {
	e.Append(s)
	return e
}

// Stream appends <Stream> element with options
func (e *start) Stream(options ...attr.Option) Start {
	e.Append(NewStream(options...))
	return e
}

// AppendStream appends <Stream> element
func (e *start) AppendStream(s Stream) Start {
	e.Append(s)
	return e
}
