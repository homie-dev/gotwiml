package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Stop is <Stop> twiml verb
	Stop interface {
		Siprec(...attr.Option) Stop
		AppendSiprec(Siprec) Stop
		Stream(...attr.Option) Stop
		AppendStream(Stream) Stop
		core.XMLer
		core.EmbedXMLer
	}

	stop struct {
		*core.XML
	}
)

// NewStop creates <Stop> element
func NewStop(options ...attr.Option) Stop {
	c := core.NewCoreXML(tagStop)
	for _, o := range options {
		o(c)
	}
	return &stop{XML: c}
}

// GetEmbedXML returns embed xml
func (e *stop) GetEmbedXML() core.XMLer {
	return e.XML
}

// AppendSiprec appends <Siprec> element with optons
func (e *stop) Siprec(options ...attr.Option) Stop {
	e.Append(NewSiprec(options...))
	return e
}

// AppendSiprec appends <Siprec> element
func (e *stop) AppendSiprec(s Siprec) Stop {
	e.Append(s)
	return e
}

// AppendStream appends <Stream> element with options
func (e *stop) Stream(options ...attr.Option) Stop {
	e.Append(NewStream(options...))
	return e
}

// AppendStream appends <Stream> element
func (e *stop) AppendStream(s Stream) Stop {
	e.Append(s)
	return e
}
