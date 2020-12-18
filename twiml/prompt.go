package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Prompt is <Prompt> twiml verb
	Prompt interface {
		Say(message string, options ...attr.Option) Prompt
		AppendSay(Say) Prompt
		Play(url string, options ...attr.Option) Prompt
		AppendPlay(Play) Prompt
		Pause(options ...attr.Option) Prompt
		AppendPause(Pause) Prompt
		core.XMLer
		core.EmbedXMLer
	}

	prompt struct {
		*core.XML
	}
)

// NewPrompt creates a <Prompt> element
func NewPrompt(options ...attr.Option) Prompt {
	t := core.NewCoreXML(tagPrompt)
	for _, o := range options {
		o(t)
	}
	return &prompt{XML: t}
}

// GetEmbedXML returns embed xml
func (e *prompt) GetEmbedXML() core.XMLer {
	return e.XML
}

// Say appends <Say> element with options
func (e *prompt) Say(message string, options ...attr.Option) Prompt {
	e.Append(NewSay(message, options...))
	return e
}

// AppendSay appends <Say> element
func (e *prompt) AppendSay(s Say) Prompt {
	e.Append(s)
	return e
}

// Play appends <Play> element with options
func (e *prompt) Play(url string, options ...attr.Option) Prompt {
	e.Append(NewPlay(url, options...))
	return e
}

// AppendPlay appends <Play> element
func (e *prompt) AppendPlay(s Play) Prompt {
	e.Append(s)
	return e
}

// Pause appends <Pause> element with options
func (e *prompt) Pause(options ...attr.Option) Prompt {
	e.Append(NewPause(options...))
	return e
}

// AppendPause appends <Pause> element
func (e *prompt) AppendPause(s Pause) Prompt {
	e.Append(s)
	return e
}
