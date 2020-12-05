package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Leave is <Leave> twiml verb
	Leave interface {
		core.XMLer
		core.EmbedXMLer
	}

	leave struct {
		*core.XML
	}
)

// NewLeave creates a <Leave> element
func NewLeave() Leave {
	t := core.NewCoreXML(verbLeave)
	return &leave{XML: t}
}

// GetEmbedXML returns embed xml
func (g *leave) GetEmbedXML() core.XMLer {
	return g.XML
}
