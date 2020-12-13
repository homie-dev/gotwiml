package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Connect is <Connect> twiml verb
	Connect interface {
		Autopilot(AssistantSID string) Connect
		AppendAutopilot(Autopilot) Connect
		Room(UniqueName string) Connect
		AppendRoom(Room) Connect
		Stream(...attr.Option) Connect
		AppendStream(Stream) Connect
		core.XMLer
		core.EmbedXMLer
	}

	connect struct {
		*core.XML
	}
)

// NewConnect creates a <Connect> element
func NewConnect(options ...attr.Option) Connect {
	t := core.NewCoreXML(tagConnect)
	for _, o := range options {
		o(t)
	}
	return &connect{XML: t}
}

// GetEmbedXML returns embed xml
func (c *connect) GetEmbedXML() core.XMLer {
	return c.XML
}

// Autopilot appends <Autopilot> element
func (c *connect) Autopilot(name string) Connect {
	c.Append(NewAutopilot(name))
	return c
}

// Room appends <Room> element
func (c *connect) Room(UniqueName string) Connect {
	c.Append(NewRoom(UniqueName))
	return c
}

// Stream appends <Stream> element
func (c *connect) Stream(oo ...attr.Option) Connect {
	c.Append(NewStream(oo...))
	return c
}

// AppendAutopilot appends <Autopilot> element
func (c *connect) AppendAutopilot(e Autopilot) Connect {
	c.Append(e)
	return c
}

// AppendRoom appends <Room> element
func (c *connect) AppendRoom(e Room) Connect {
	c.Append(e)
	return c
}

// AppendStream appends <Stream> element
func (c *connect) AppendStream(e Stream) Connect {
	c.Append(e)
	return c
}
