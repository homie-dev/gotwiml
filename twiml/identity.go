package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Identity is <Identity> twiml verb
	Identity interface {
		SetIdentity(string) Identity
		core.XMLer
		core.EmbedXMLer
	}

	identity struct {
		*core.XML
	}
)

// NewIdentity creates a <Identity> element
func NewIdentity(clientIdentity string) Identity {
	t := core.NewCoreXML(nounIdentity)
	t.SetText(clientIdentity)
	return &identity{XML: t}
}

// GetEmbedXML returns embed xml
func (e *identity) GetEmbedXML() core.XMLer {
	return e.XML
}

func (e *identity) SetIdentity(clientIdentity string) Identity {
	e.SetText(clientIdentity)
	return e
}
