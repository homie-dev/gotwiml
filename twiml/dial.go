package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/core"
)

type (
	// Dial is <Dial> twiml verb
	Dial interface {
		Number(phoneNumber string, options ...attr.Option) Dial
		Client(identifier string, options ...attr.Option) Dial
		Conference(roomName string, options ...attr.Option) Dial
		Queue(queueName string, options ...attr.Option) Dial
		Enqueue(queueName string, options ...attr.Option) Dial
		Sim(sid string) Dial
		Sip(uri string, options ...attr.Option) Dial

		AppendNumber(Number) Dial
		AppendClient(Client) Dial
		AppendConference(Client) Dial
		AppendQueue(Queue) Dial
		AppendEnqueue(Enqueue) Dial
		AppendSim(Sim) Dial
		AppendSip(Sip) Dial
		core.XMLer
		core.EmbedXMLer
	}

	dial struct {
		*core.XML
	}
)

// NewDial creates a <Dial> element with options
func NewDial(options ...attr.Option) Dial {
	t := core.NewCoreXML(tagDial)
	for _, o := range options {
		o(t)
	}
	return &dial{XML: t}
}

// GetEmbedXML returns embed xml
func (d *dial) GetEmbedXML() core.XMLer {
	return d.XML
}

// Number appends <Number> element with options
func (d *dial) Number(phoneNumber string, attrs ...attr.Option) Dial {
	d.Append(NewNumber(phoneNumber, attrs...))
	return d
}

// AppendNumber appends <Number> element
func (d *dial) AppendNumber(n Number) Dial {
	d.Append(n)
	return d
}

// Client appends <Client> element with options
func (d *dial) Client(identifier string, attrs ...attr.Option) Dial {
	d.Append(NewClient(identifier, attrs...))
	return d
}

func (d *dial) AppendClient(c Client) Dial {
	d.Append(c)
	return d
}

// Conference appends <Conference> element with options
func (d *dial) Conference(roomName string, attrs ...attr.Option) Dial {
	d.Append(NewConference(roomName, attrs...))
	return d
}

// Conference appends <Conference> element
func (d *dial) AppendConference(c Client) Dial {
	d.Append(c)
	return d
}

// Queue appends <Conference> element with options
func (d *dial) Queue(queueName string, options ...attr.Option) Dial {
	d.Append(NewQueue(queueName, options...))
	return d
}

// Enqueue appends <Conference> element with options
func (d *dial) Enqueue(queueName string, options ...attr.Option) Dial {
	d.Append(NewQueue(queueName, options...))
	return d
}

// Sim appends <Conference> element with options
func (d *dial) Sim(sid string) Dial {
	d.Append(NewSim(sid))
	return d
}

// Sip appends <Conference> element with options
func (d *dial) Sip(uri string, options ...attr.Option) Dial {
	d.Append(NewSip(uri, options...))
	return d
}

// AppendQueue appends <Conference> element
func (d *dial) AppendQueue(e Queue) Dial {
	d.Append(e)
	return d
}

// AppendEnqueue appends <Conference> element
func (d *dial) AppendEnqueue(e Enqueue) Dial {
	d.Append(e)
	return d
}

// AppendSim appends <Conference> element
func (d *dial) AppendSim(e Sim) Dial {
	d.Append(e)
	return d
}

// AppendSip appends <Conference> element
func (d *dial) AppendSip(e Sip) Dial {
	d.Append(e)
	return d
}
