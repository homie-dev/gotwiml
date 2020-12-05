package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Autopilot is <Autopilot> twiml verb
	Autopilot interface {
		core.XMLer
		core.EmbedXMLer
	}

	autopilot struct {
		*core.XML
	}
)

// NewAutopilot creates a <Autopilot> element
func NewAutopilot(name string) Autopilot {
	t := core.NewCoreXML(verbAutopilot)
	t.SetText(name)
	return &autopilot{XML: t}
}

// GetEmbedXML returns embed xml
func (c *autopilot) GetEmbedXML() core.XMLer {
	return c.XML
}
