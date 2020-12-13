package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Play is <Play> twiml verb
	Play interface {
		core.XMLer
		core.EmbedXMLer
	}

	play struct {
		*core.XML
	}
)

// NewPlay creates a <Play> element
func NewPlay(url string, options ...attr.Option) Play {
	t := core.NewCoreXML(tagPlay)
	t.SetText(url)
	for _, o := range options {
		o(t)
	}
	return &play{XML: t}
}

// GetEmbedXML returns embed xml
func (e *play) GetEmbedXML() core.XMLer {
	return e.XML
}
