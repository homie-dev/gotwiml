package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/const/http"
	"github.com/homie-dev/gotwiml/twiml/const/status"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Client is <Client> twiml verb
	Client interface {
		Identity(string) Client
		Parameter(...attr.Option) Client
		SetURL(url string) Client
		SetMethod(method http.Method) Client
		SetStatusCallbackEvent(event status.CallbackEvent) Client
		SetStatusCallback(url string) Client
		SetStatusCallbackMethod(method http.Method) Client
		core.XMLer
		core.EmbedXMLer
	}

	client struct {
		*core.XML
	}
)

// NewClient creates <Client> element
func NewClient(options ...attr.Option) Client {
	c := core.NewCoreXML(nounClient)
	for _, o := range options {
		o(c)
	}
	return &client{XML: c}
}

// GetEmbedXML returns embed xml
func (c *client) GetEmbedXML() core.XMLer {
	return c.XML
}

func (c *client) Identity(v string) Client {
	x := core.NewCoreXML(nounIdentity).SetText(v)
	c.Append(x)
	return c
}
func (c *client) Parameter(attrs ...attr.Option) Client {
	x := core.NewCoreXML(nounParameter)
	for _, a := range attrs {
		a(x)
	}
	c.Append(x)
	return c
}

// SetURL sets client url
func (c *client) SetURL(url string) Client {
	attr.URL(url)(c)
	return c
}

// SetMethod sets client url method
func (c *client) SetMethod(method http.Method) Client {
	attr.Method(method)(c)
	return c
}

// SetStatusCallbackEvent sets event to trigger status callback
func (c *client) SetStatusCallbackEvent(event status.CallbackEvent) Client {
	attr.StatusCallbackEvent(event)(c)
	return c
}

// SetStatusCallback sets status callback url
func (c *client) SetStatusCallback(url string) Client {
	attr.StatusCallback(url)(c)
	return c
}

// SetStatusCallbackMethod sets status callback method
func (c *client) SetStatusCallbackMethod(method http.Method) Client {
	attr.RecordingStatusCallbackMethod(method)(c)
	return c
}
