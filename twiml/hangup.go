package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Hangup is <Hangup> twiml verb
	Hangup interface {
		core.XMLer
		core.EmbedXMLer
	}

	hangup struct {
		*core.XML
	}
)

// NewHangup creates a <Hangup> element
func NewHangup() Hangup {
	t := core.NewCoreXML(verbHangup)
	return &hangup{XML: t}
}

// GetEmbedXML returns embed xml
func (g *hangup) GetEmbedXML() core.XMLer {
	return g.XML
}
