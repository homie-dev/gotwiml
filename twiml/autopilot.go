package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Autopilot is <Autopilot> twiml verb
	Autopilot interface {
		SetName(string) Autopilot
		core.XMLer
		core.EmbedXMLer
	}

	autopilot struct {
		*core.XML
	}
)

// NewAutopilot creates a <Autopilot> element
func NewAutopilot(name string) Autopilot {
	t := core.NewCoreXML(tagAutopilot)
	t.SetText(name)
	return &autopilot{XML: t}
}

// GetEmbedXML returns embed xml
func (e *autopilot) GetEmbedXML() core.XMLer {
	return e.XML
}

// SetName sets name
func (e *autopilot) SetName(name string) Autopilot {
	e.SetText(name)
	return e
}
