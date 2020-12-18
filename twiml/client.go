package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Client is <Client> twiml verb
	Client interface {
		SetIdentity(string) Client
		Identity(string) Client
		AppendIdentity(Identity) Client
		Parameter(...attr.Option) Client
		AppendParameter(Parameter) Client
		core.XMLer
		core.EmbedXMLer
	}

	client struct {
		*core.XML
	}
)

// NewClient creates <Client> element
func NewClient(identifier string, options ...attr.Option) Client {
	c := core.NewCoreXML(tagClient)
	c.SetText(identifier)
	for _, o := range options {
		o(c)
	}
	return &client{XML: c}
}

// GetEmbedXML returns embed xml
func (e *client) GetEmbedXML() core.XMLer {
	return e.XML
}

// Identity appends <Identity> element
func (e *client) Identity(v string) Client {
	e.Append(NewIdentity(v))
	return e
}

// Parameter appends <Parameter> element with options
func (e *client) Parameter(attrs ...attr.Option) Client {
	e.Append(NewParameter(attrs...))
	return e
}

// SetIdentity sets client identity
func (e *client) SetIdentity(i string) Client {
	e.SetText(i)
	return e
}

// AppendIdentity appends <Identity> element
func (e *client) AppendIdentity(i Identity) Client {
	e.Append(i)
	return e
}

// AppendParameter appends <Parameter> element
func (e *client) AppendParameter(p Parameter) Client {
	e.Append(p)
	return e
}
