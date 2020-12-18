package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Pay is <Pay> twiml verb
	Pay interface {
		Prompt(options ...attr.Option) Pay
		AppendPrompt(Prompt) Pay
		Parameter(options ...attr.Option) Pay
		AppendParameter(Parameter) Pay
		core.XMLer
		core.EmbedXMLer
	}

	pay struct {
		*core.XML
	}
)

// NewPay creates a <Pay> element
func NewPay(options ...attr.Option) Pay {
	t := core.NewCoreXML(tagPay)
	for _, o := range options {
		o(t)
	}
	return &pay{XML: t}
}

// GetEmbedXML returns embed xml
func (e *pay) GetEmbedXML() core.XMLer {
	return e.XML
}

// Prompt appends <Prompt> element with options
func (e *pay) Prompt(options ...attr.Option) Pay {
	e.Append(NewPrompt(options...))
	return e
}

// AppendPrompt appends <Prompt> element
func (e *pay) AppendPrompt(p Prompt) Pay {
	e.Append(p)
	return e
}

// Parameter appends <Parameter> element with options
func (e *pay) Parameter(options ...attr.Option) Pay {
	e.Append(NewParameter(options...))
	return e
}

// AppendParameter appends <Parameter> element
func (e *pay) AppendParameter(p Parameter) Pay {
	e.Append(p)
	return e
}
