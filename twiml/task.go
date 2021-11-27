package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Task is <Task> twiml verb
	Task interface {
		SetBody(body string) Task
		core.XMLer
		core.EmbedXMLer
	}

	task struct {
		*core.XML
	}
)

// NewTask creates a <Task> element
func NewTask(body string, options ...attr.Option) Task {
	t := core.NewCoreXML(tagTask)
	t.SetText(body)
	for _, o := range options {
		o(t)
	}
	return &task{XML: t}
}

// GetEmbedXML returns embed xml
func (c *task) GetEmbedXML() core.XMLer {
	return c.XML
}

// SetBody sets task body
func (c *task) SetBody(body string) Task {
	c.SetText(body)
	return c
}
