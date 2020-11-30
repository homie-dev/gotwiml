package twiml

import (
	"net/http"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/attr/const/status"
)

var _ = DescribeTable("generate <Number> xml",
	func(c Number, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewNumber("+12345678"), "<Number>+12345678</Number>"),
	Entry("set attributes", NewNumber("+12345678",
		attr.SendDigits("wwww1928"),
		attr.URL("http://example.url"),
		attr.Method(http.MethodPost),
		attr.BYOC("sid"),
		attr.StatusCallbackEvent(status.CallbackInitiated, status.CallbackCompleted),
		attr.StatusCallback("http://callback.url"),
		attr.StatusCallbackMethod(http.MethodPost),
	), `<Number sendDigits="wwww1928" url="http://example.url" method="POST" byoc="sid" statusCallbackEvent="initiated completed" statusCallback="http://callback.url" statusCallbackMethod="POST">+12345678</Number>`),
)
