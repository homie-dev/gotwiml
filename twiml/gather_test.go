package twiml

import (
	"net/http"

	"github.com/homie-dev/gotwiml/twiml/attr/const/language"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/attr/const/input"
	"github.com/homie-dev/gotwiml/twiml/attr/const/speech"
)

var _ = DescribeTable("generate <Gather> xml",
	func(c Gather, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewGather(), "<Gather></Gather>"),
	Entry("set attributes", NewGather(
		attr.Action("action"),
		attr.FinishOnKey("*"),
		attr.Hints("hints"),
		attr.Input(input.DTMF, input.Speech),
		attr.Language(language.JaJp),
		attr.Method(http.MethodPost),
		attr.NumDigits(1),
		attr.PartialResultCallback("http://partial_callback.url"),
		attr.PartialResultCallbackMethod(http.MethodPost),
		attr.ProfanityFilter(true),
		attr.SpeechTimeout("auto"),
		attr.Timeout(10),
		attr.SpeechModel(speech.ModelDefault),
		attr.Enhanced(true),
		attr.ActionOnEmptyResult(true),
	), `<Gather action="action" finishOnKey="*" hints="hints" input="dtmf speech" language="ja-JP" method="POST" numDigits="1" partialResultCallback="http://partial_callback.url" partialResultCallbackMethod="POST" profanityFilter="true" speechTimeout="auto" timeout="10" speechModel="default" enhanced="true" actionOnEmptyResult="true"></Gather>`),
)
