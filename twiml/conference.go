package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Conference is <Conference> twiml verb
	Conference interface {
		SetName(roomName string) Conference
		core.XMLer
		core.EmbedXMLer
	}

	conference struct {
		*core.XML
	}
)

// NewConference creates a <Conference> element
func NewConference(roomName string, options ...attr.Option) Conference {
	t := core.NewCoreXML(tagConference)
	t.SetText(roomName)
	for _, o := range options {
		o(t)
	}
	return &conference{XML: t}
}

// GetEmbedXML returns embed xml
func (c *conference) GetEmbedXML() core.XMLer {
	return c.XML
}

// SetName sets room name
func (c *conference) SetName(roomName string) Conference {
	c.SetText(roomName)
	return c
}
