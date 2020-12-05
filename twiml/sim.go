package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Sim is <Sim> twiml verb
	Sim interface {
		core.XMLer
		core.EmbedXMLer
	}

	sim struct {
		*core.XML
	}
)

// NewSim creates a <Sim> element
func NewSim(sid string) Sim {
	t := core.NewCoreXML(nounSim)
	t.SetText(sid)
	return &sim{XML: t}
}

// GetEmbedXML returns embed xml
func (c *sim) GetEmbedXML() core.XMLer {
	return c.XML
}
