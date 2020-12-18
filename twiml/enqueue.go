package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Enqueue is <Enqueue> twiml verb
	Enqueue interface {
		SetQueueName(queueName string) Enqueue
		Task(body string, options ...attr.Option) Enqueue
		AppendTask(Task) Enqueue
		core.XMLer
		core.EmbedXMLer
	}

	enqueue struct {
		*core.XML
	}
)

// NewEnqueue creates a <Enqueue> element
func NewEnqueue(queueName string, options ...attr.Option) Enqueue {
	t := core.NewCoreXML(tagEnqueue)
	t.SetText(queueName)
	for _, o := range options {
		o(t)
	}
	return &enqueue{XML: t}
}

// GetEmbedXML returns embed xml
func (c *enqueue) GetEmbedXML() core.XMLer {
	return c.XML
}

// SetQueueName sets queue name
func (c *enqueue) SetQueueName(queueName string) Enqueue {
	c.SetText(queueName)
	return c
}

// Task appends <Task> element with options
func (c *enqueue) Task(body string, attrs ...attr.Option) Enqueue {
	c.Append(NewTask(body, attrs...))
	return c
}

// AppendTask appends <Task> element
func (c *enqueue) AppendTask(n Task) Enqueue {
	c.Append(n)
	return c
}
