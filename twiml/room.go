package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Room is <Room> twiml verb
	Room interface {
		SetUniqueName(string) Room
		core.XMLer
		core.EmbedXMLer
	}

	room struct {
		*core.XML
	}
)

// NewRoom creates a <Room> element
func NewRoom(clientRoom string) Room {
	t := core.NewCoreXML(tagRoom)
	t.SetText(clientRoom)
	return &room{XML: t}
}

// GetEmbedXML returns embed xml
func (e *room) GetEmbedXML() core.XMLer {
	return e.XML
}

// SetUniqueName sets unique name
func (e *room) SetUniqueName(v string) Room {
	e.SetText(v)
	return e
}
