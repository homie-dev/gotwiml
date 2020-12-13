package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Pay is <Pay> twiml verb
	Pay interface {
		core.XMLer
		core.EmbedXMLer
	}

	pay struct {
		*core.XML
	}
)

// NewPay creates a <Pay> element
func NewPay(options ...attr.Option) Pay {
	t := core.NewCoreXML(verbPay)
	for _, o := range options {
		o(t)
	}
	return &pay{XML: t}
}

// GetEmbedXML returns embed xml
func (g *pay) GetEmbedXML() core.XMLer {
	return g.XML
}
