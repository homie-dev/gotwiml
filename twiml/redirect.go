package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Redirect is <Redirect> twiml verb
	Redirect interface {
		core.XMLer
		core.EmbedXMLer
	}

	redirect struct {
		*core.XML
	}
)

// NewRedirect creates a <Redirect> element
func NewRedirect(url string, options ...attr.Option) Redirect {
	t := core.NewCoreXML(tagRedirect)
	t.SetText(url)
	for _, o := range options {
		o(t)
	}
	return &redirect{XML: t}
}

// GetEmbedXML returns embed xml
func (e *redirect) GetEmbedXML() core.XMLer {
	return e.XML
}
