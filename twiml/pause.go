package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Pause is <Pause> twiml verb
	Pause interface {
		core.XMLer
		core.EmbedXMLer
	}

	pause struct {
		*core.XML
	}
)

// NewPause creates a <Pause> element
func NewPause(options ...attr.Option) Pause {
	t := core.NewCoreXML(verbPause)
	for _, o := range options {
		o(t)
	}
	return &pause{XML: t}
}

// GetEmbedXML returns embed xml
func (g *pause) GetEmbedXML() core.XMLer {
	return g.XML
}
