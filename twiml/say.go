package twiml

import (
	"strconv"

	"gotwiml/twiml/const/say"
)

func Voice(v say.Voice) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrVoice, string(v))
	}
}

func Language(l say.Language) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrLanguage, string(l))
	}
}

func Loop(n int) Attr {
	return func(t TwiML) {
		t.SetAttr(AttrLoop, strconv.Itoa(n))
	}
}
