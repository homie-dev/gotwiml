package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Gather is <Gather> twiml verb
	Gather interface {
		Say(message string, options ...attr.Option) Gather
		core.XMLer
		core.EmbedXMLer
	}

	gather struct {
		*core.XML
	}
)

// NewGather creates a <Gather> element
func NewGather(options ...attr.Option) Gather {
	t := core.NewCoreXML(verbGather)
	for _, o := range options {
		o(t)
	}
	return &gather{XML: t}
}

// GetEmbedXML returns embed xml
func (g *gather) GetEmbedXML() core.XMLer {
	return g.XML
}

// Say appends a <Say> element with options
func (g *gather) Say(message string, options ...attr.Option) Gather {
	g.Append(NewSay(message, options...))
	return g
}
