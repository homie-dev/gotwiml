package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/const/http"
	"github.com/homie-dev/gotwiml/twiml/const/status"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	Client interface {
		SetURL(url string) Client
		SetMethod(method http.Method) Client
		SetStatusCallbackEvent(event status.CallbackEvent) Client
		SetStatusCallback(url string) Client
		SetStatusCallbackMethod(method http.Method) Client
		core.XMLer
		core.EmbedXMLer
	}
	clientImpl struct {
		*core.XML
	}
)

// NewClient creates <Client> element
func NewClient(options ...attr.Option) Client {
	c := core.NewCoreXML(nounClient)
	for _, o := range options {
		o(c)
	}
	return &clientImpl{XML: c}
}

func (c *clientImpl) GetEmbedXML() core.XMLer {
	return c.XML
}

// SetURL sets client url
func (c *clientImpl) SetURL(url string) Client {
	attr.URL(url)(c)
	return c
}

// SetMethod sets client url method
func (c *clientImpl) SetMethod(method http.Method) Client {
	attr.Method(method)(c)
	return c
}

// SetStatusCallbackEvent sets event to trigger status callback
func (c *clientImpl) SetStatusCallbackEvent(event status.CallbackEvent) Client {
	attr.StatusCallbackEvent(event)(c)
	return c
}

// SetStatusCallback sets status callback url
func (c *clientImpl) SetStatusCallback(url string) Client {
	attr.StatusCallback(url)(c)
	return c
}

// SetStatusCallbackMethod sets status callback method
func (c *clientImpl) SetStatusCallbackMethod(method http.Method) Client {
	attr.RecordingStatusCallbackMethod(method)(c)
	return c
}
