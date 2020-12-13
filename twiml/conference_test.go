package twiml

import (
	"net/http"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/attr/const/beep"
	"github.com/homie-dev/gotwiml/twiml/attr/const/jitter"
	"github.com/homie-dev/gotwiml/twiml/attr/const/recording"
	"github.com/homie-dev/gotwiml/twiml/attr/const/region"
	"github.com/homie-dev/gotwiml/twiml/attr/const/status"
	"github.com/homie-dev/gotwiml/twiml/attr/const/trim"
)

var _ = DescribeTable("generate <Conference> xml",
	func(c Conference, xml string) {
		Expect(c.String()).To(Equal(xml))
	},
	Entry("empty connect", NewConference("roomName"), "<Conference>roomName</Conference>"),
	Entry("set attributes", NewConference("roomName",
		attr.Muted(true),
		attr.Beep(beep.OnEnter),
		attr.StartConferenceOnEnter(true),
		attr.EndConferenceOnExit(true),
		attr.ParticipantLabel("pLabel"),
		attr.JitterBufferSize(jitter.BufferSizeMedium),
		attr.WaitURL("http://wait.url"),
		attr.WaitMethod(http.MethodPost),
		attr.MaxParticipants(10),
		attr.Record(recording.FromStart),
		attr.Region(region.AU1),
		attr.Trim(trim.DoNot),
		attr.Coach("sid"),
		attr.StatusCallbackEvent(status.CallbackStart),
		attr.StatusCallback("http://callback.url"),
		attr.StatusCallbackMethod(http.MethodPost),
		attr.RecordingStatusCallback("http://recording.url"),
		attr.RecordingStatusCallbackMethod(http.MethodPost),
		attr.RecordingStatusCallbackEvent(recording.StatusCallbackEventAbsent),
	), `<Conference muted="true" beep="onEnter" startConferenceOnEnter="true" endConferenceOnExit="true" participantLabel="pLabel" jitterBufferSize="medium" waitUrl="http://wait.url" waitMethod="POST" maxParticipants="10" record="record-from-start" region="au1" trim="do-not-trim" coach="sid" statusCallbackEvent="start" statusCallback="http://callback.url" statusCallbackMethod="POST" recordingStatusCallback="http://recording.url" recordingStatusCallbackMethod="POST" recordingStatusCallbackEvent="absent">roomName</Conference>`),
)
