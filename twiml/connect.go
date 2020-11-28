package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Connect is <Connect> twiml verb
	Connect interface {
		Autopilot(AssistantSID string) Connect
		Room(UniqueName string) Connect
		Stream(...attr.Option) Connect
		core.XMLer
		core.EmbedXMLer
	}

	connect struct {
		*core.XML
	}
)

// NewConnect creates a <Connect> element
func NewConnect(options ...attr.Option) Connect {
	t := core.NewCoreXML(verbConnect)
	for _, o := range options {
		o(t)
	}
	return &connect{XML: t}
}

// GetEmbedXML returns embed xml
func (c *connect) GetEmbedXML() core.XMLer {
	return c.XML
}

func (c *connect) Autopilot(AssistantSID string) Connect {
	t := core.NewXML(connectNounAutopilot)
	t.SetText(AssistantSID)
	c.Append(t)
	return c
}
func (c *connect) Room(UniqueName string) Connect {
	t := core.NewXML(connectNounRoom)
	t.SetText(UniqueName)
	c.Append(t)
	return c
}
func (c *connect) Stream(oo ...attr.Option) Connect {
	t := core.NewXML(connectNounStream)
	for _, o := range oo {
		o(t)
	}
	c.Append(t)
	return c
}
