package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("generate <Connect> xml",
	func(c Connect, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewConnect(), "<Connect></Connect>"),
	Entry("append <Autopilot>", NewConnect().Autopilot("asid"), `<Connect><Autopilot>asid</Autopilot></Connect>`),
	Entry("append <Room>", NewConnect().Room("unique_name"), `<Connect><Room>unique_name</Room></Connect>`),
	Entry("append <Stream>", NewConnect().Stream(attr.URL("url")), `<Connect><Stream url="url"></Stream></Connect>`),
)
