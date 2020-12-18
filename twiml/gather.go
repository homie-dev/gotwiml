package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Gather is <Gather> twiml verb
	Gather interface {
		Say(message string, options ...attr.Option) Gather
		AppendSay(Say) Gather
		Pause(options ...attr.Option) Gather
		AppendPause(Pause) Gather
		Play(url string, options ...attr.Option) Gather
		AppendPlay(Play) Gather
		core.XMLer
		core.EmbedXMLer
	}

	gather struct {
		*core.XML
	}
)

// NewGather creates a <Gather> element
func NewGather(options ...attr.Option) Gather {
	t := core.NewCoreXML(tagGather)
	for _, o := range options {
		o(t)
	}
	return &gather{XML: t}
}

// GetEmbedXML returns embed xml
func (e *gather) GetEmbedXML() core.XMLer {
	return e.XML
}

// Pause appends a <Pause> element with options
func (e *gather) Pause(options ...attr.Option) Gather {
	e.Append(NewPause(options...))
	return e
}

// AppendPause appends <Pause> element
func (e *gather) AppendPause(s Pause) Gather {
	e.Append(s)
	return e
}

// Play appends a <Play> element with options
func (e *gather) Play(url string, options ...attr.Option) Gather {
	e.Append(NewPlay(url, options...))
	return e
}

// AppendPlay appends <Play> element
func (e *gather) AppendPlay(s Play) Gather {
	e.Append(s)
	return e
}

// Say appends a <Say> element with options
func (e *gather) Say(message string, options ...attr.Option) Gather {
	e.Append(NewSay(message, options...))
	return e
}

// AppendSay appends <Say> element
func (e *gather) AppendSay(s Say) Gather {
	e.Append(s)
	return e
}
