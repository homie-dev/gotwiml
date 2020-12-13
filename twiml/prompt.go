package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Prompt is <Prompt> twiml verb
	Prompt interface {
		Say(message string, options ...attr.Option) Prompt
		core.XMLer
		core.EmbedXMLer
	}

	prompt struct {
		*core.XML
	}
)

// NewPrompt creates a <Prompt> element
func NewPrompt(options ...attr.Option) Prompt {
	t := core.NewCoreXML(nounPrompt)
	for _, o := range options {
		o(t)
	}
	return &prompt{XML: t}
}

// GetEmbedXML returns embed xml
func (e *prompt) GetEmbedXML() core.XMLer {
	return e.XML
}

func (e *prompt) Say(message string, options ...attr.Option) Prompt {
	e.Append(NewSay(message, options...))
	return e
}
