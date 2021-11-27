package main

import (
	"fmt"

	"github.com/homie-dev/gotwiml/twiml"
	"github.com/homie-dev/gotwiml/twiml/attr"
)

func main() {
	client := twiml.NewClient("hoge").
		Parameter(attr.Name("foo"), attr.Value("bar"))
	resp := twiml.NewVoiceResponse().
		AppendDial(twiml.NewDial().AppendClient(client))
	twiml, _ := resp.ToXMLPretty("  ")
	fmt.Println(twiml)
	// <?xml version="1.0" encoding="UTF-8"?>
	// <Response>
	//   <Dial>
	//     <Client>hoge
	//       <Parameter name="foo" value="bar"></Parameter>
	//     </Client>
	//   </Dial>
	// </Response>
}
