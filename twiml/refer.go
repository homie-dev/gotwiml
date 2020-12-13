package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Refer is <Refer> twiml verb
	Refer interface {
		SIP(uri string, options ...attr.Option) Refer
		AppendSIP(Sip) Refer
		core.XMLer
		core.EmbedXMLer
	}

	refer struct {
		*core.XML
	}
)

// NewRefer creates a <Refer> element
func NewRefer(options ...attr.Option) Refer {
	t := core.NewCoreXML(verbRefer)
	for _, o := range options {
		o(t)
	}
	return &refer{XML: t}
}

// GetEmbedXML returns embed xml
func (e *refer) GetEmbedXML() core.XMLer {
	return e.XML
}

func (e *refer) SIP(uri string, options ...attr.Option) Refer {
	e.Append(NewSip(uri, options...))
	return e
}
func (e *refer) AppendSIP(s Sip) Refer {
	e.Append(s)
	return e
}
