package twiml

import (
	"net/http"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/homie-dev/gotwiml/twiml/attr"
)

var _ = DescribeTable("generate <Enqueue> xml",
	func(c Enqueue, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewEnqueue("support"), "<Enqueue>support</Enqueue>"),
	Entry("set attributes", NewEnqueue("support",
		attr.Action("action"),
		attr.Method(http.MethodPost),
		attr.WaitURL("http://wait.url"),
		attr.WaitURLMethod(http.MethodPost),
		attr.WorkflowSID("wsid"),
	), `<Enqueue action="action" method="POST" waitUrl="http://wait.url" waitUrlMethod="POST" workflowSid="wsid">support</Enqueue>`),
)
