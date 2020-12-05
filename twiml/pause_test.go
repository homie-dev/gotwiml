package twiml

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/homie-dev/gotwiml/twiml/attr"
)

var _ = DescribeTable("generate <Pause> xml",
	func(c Pause, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewPause(), "<Pause></Pause>"),
	Entry("set attributes", NewPause(
		attr.Length(10),
	), `<Pause length="10"></Pause>`),
)
