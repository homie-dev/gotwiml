package twiml

import (
	"twilio-golang/twiml/const/client"
	"twilio-golang/twiml/const/http"
)

type (
	ClientAttr func(t TwiML)

	Client interface {
		SetURL(url string) Client
		SetMethod(method http.Method) Client
		SetStatusCallbackEvent(event client.StatusCallbackEvent) Client
		SetStatusCallback(url string) Client
		SetStatusCallbackMethod(method http.Method) Client
		TwiML
		EmbedTwiML
	}
	clientImpl struct {
		*twiML
	}
)

// NewClient creates <Client> element
func NewClient(options ...ClientAttr) Client {
	c := new(nounClient)
	for _, o := range options {
		o(c)
	}
	return &clientImpl{twiML: c}
}

func (c *clientImpl) getTwiML() TwiML {
	return c.twiML
}

// SetURL sets client url
func (c *clientImpl) SetURL(url string) Client {
	c.SetAttr(AttrURL, url)
	return c
}

// SetMethod sets client url method
func (c *clientImpl) SetMethod(method http.Method) Client {
	c.SetAttr(AttrMethod, string(method))
	return c
}

// SetStatusCallbackEvent sets event to trigger status callback
func (c *clientImpl) SetStatusCallbackEvent(event client.StatusCallbackEvent) Client {
	c.SetAttr(AttrStatusCallbackEvent, string(event))
	return c
}

// SetStatusCallback sets status callback url
func (c *clientImpl) SetStatusCallback(url string) Client {
	c.SetAttr(AttrStatusCallback, url)
	return c
}

// SetStatusCallbackMethod sets status callback method
func (c *clientImpl) SetStatusCallbackMethod(method http.Method) Client {
	c.SetAttr(AttrStatusCallbackMethod, string(method))
	return c
}
