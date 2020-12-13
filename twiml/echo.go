package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Echo is <Echo> twiml verb
	Echo interface {
		core.XMLer
		core.EmbedXMLer
	}

	echo struct {
		*core.XML
	}
)

// NewEcho creates a <Echo> element
func NewEcho() Echo {
	return &echo{XML: core.NewCoreXML(tagEcho)}
}

// GetEmbedXML returns embed xml
func (e *echo) GetEmbedXML() core.XMLer {
	return e.XML
}
