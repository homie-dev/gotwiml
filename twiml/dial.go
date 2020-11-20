package twiml

type (
	Dial interface {
		Number(phoneNumber string, attrs ...Attr) Dial
		Client(identifier string, attrs ...Attr) Dial

		TwiML
		EmbedTwiML
	}

	dialImpl struct {
		*twiML
	}
)

// NewDial creates a <Dial> element
func NewDial(options ...Attr) Dial {
	t := new(verbDial)
	for _, o := range options {
		o(t)
	}
	return &dialImpl{twiML: t}
}

func (d *dialImpl) getTwiML() TwiML {
	return d.twiML
}

// Client appends <Client> element
func (d *dialImpl) Client(identifier string, attrs ...Attr) Dial {
	t := New(nounClient).SetText(identifier)
	for _, a := range attrs {
		a(t)
	}
	d.Append(t)
	return d
}

// Number appends <Number> element
func (d *dialImpl) Number(phoneNumber string, attrs ...Attr) Dial {
	t := New(nounNumber).SetText(phoneNumber)
	for _, a := range attrs {
		a(t)
	}
	d.Append(t)
	return d
}
