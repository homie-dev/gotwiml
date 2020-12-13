package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Record is <Record> twiml verb
	Record interface {
		core.XMLer
		core.EmbedXMLer
	}

	record struct {
		*core.XML
	}
)

// NewRecord creates a <Record> element
func NewRecord(options ...attr.Option) Record {
	t := core.NewCoreXML(verbRecord)
	for _, o := range options {
		o(t)
	}
	return &record{XML: t}
}

// GetEmbedXML returns embed xml
func (e *record) GetEmbedXML() core.XMLer {
	return e.XML
}
