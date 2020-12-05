package twiml

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/homie-dev/gotwiml/twiml/attr"
)

var _ = DescribeTable("generate <Sip> xml",
	func(c Sip, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewSip("sip:jack@example.com"), "<Sip>sip:jack@example.com</Sip>"),
	Entry("set attributes", NewSip("sip:kate@example.com",
		attr.UserName("admin"),
		attr.Password("1234"),
	), `<Sip username="admin" password="1234">sip:kate@example.com</Sip>`),
)
