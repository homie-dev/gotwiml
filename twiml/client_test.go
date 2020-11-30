package twiml

import (
	"net/http"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/attr/const/status"
)

var _ = DescribeTable("generate <Client> xml",
	func(c Client, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewClient(""), "<Client></Client>"),
	Entry("set attributes", NewClient("",
		attr.URL("http://example.com"),
		attr.Method(http.MethodGet),
		attr.StatusCallbackEvent(status.CallbackInitiated),
		attr.StatusCallback("http://callback.com"),
		attr.StatusCallbackMethod(http.MethodPost),
	), `<Client url="http://example.com" method="GET" statusCallbackEvent="initiated" statusCallback="http://callback.com" statusCallbackMethod="POST"></Client>`),
	Entry("append <Identity>", NewClient("").
		Identity("user-jane").
		Parameter(attr.Name("FirstName"), attr.Value("Jane")).
		Parameter(attr.Name("LastName"), attr.Value("Doe")),
		`<Client><Identity>user-jane</Identity><Parameter name="FirstName" value="Jane"></Parameter><Parameter name="LastName" value="Doe"></Parameter></Client>`),
)
