package twiml

import (
	"net/http"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/homie-dev/gotwiml/twiml/attr"
)

var _ = DescribeTable("generate <Queue> xml",
	func(c Queue, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewQueue("support"), "<Queue>support</Queue>"),
	Entry("set attributes", NewQueue("support",
		attr.URL("http://example.url"),
		attr.Method(http.MethodPost),
		attr.ReservationSID("rsid"),
		attr.PostWorkActivitySID("pwasid"),
	), `<Queue url="http://example.url" method="POST" reservationSid="rsid" postWorkActivitySid="pwasid">support</Queue>`),
)
