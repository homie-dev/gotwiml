package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	Dial interface {
		Number(phoneNumber string, attrs ...attr.Option) Dial
		Client(identifier string, attrs ...attr.Option) Dial

		core.XMLer
		core.EmbedXMLer
	}

	dialImpl struct {
		*core.XML
	}
)

// NewDial creates a <Dial> element
func NewDial(options ...attr.Option) Dial {
	t := core.NewCoreXML(verbDial)
	for _, o := range options {
		o(t)
	}
	return &dialImpl{XML: t}
}

func (d *dialImpl) GetEmbedXML() core.XMLer {
	return d.XML
}

// Client appends <Client> element
func (d *dialImpl) Client(identifier string, attrs ...attr.Option) Dial {
	t := core.NewXML(nounClient).SetText(identifier)
	for _, a := range attrs {
		a(t)
	}
	d.Append(t)
	return d
}

// Number appends <Number> element
func (d *dialImpl) Number(phoneNumber string, attrs ...attr.Option) Dial {
	t := core.NewXML(nounNumber).SetText(phoneNumber)
	for _, a := range attrs {
		a(t)
	}
	d.Append(t)
	return d
}
