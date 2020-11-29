package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/const/status"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = DescribeTable("generate <Connect> xml",
	func(c Client, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewClient(), "<Client></Client>"),
	Entry("set attributes", NewClient(
		attr.URL("http://example.com"),
		attr.Method(http.MethodGet),
		attr.StatusCallbackEvent(status.CallbackInitiated),
		attr.StatusCallback("http://callback.com"),
		attr.StatusCallbackMethod(http.MethodPost),
		), `<Client url="http://example.com" method="GET" statusCallbackEvent="initiated" statusCallback="http://callback.com" statusCallbackMethod="POST"></Client>`),
)
