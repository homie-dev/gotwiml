package twiml

import (
	"github.com/homie-dev/gotwiml/twiml/attr"
	"github.com/homie-dev/gotwiml/twiml/const/beep"
	"github.com/homie-dev/gotwiml/twiml/const/jitter"
	"github.com/homie-dev/gotwiml/twiml/const/record"
	"github.com/homie-dev/gotwiml/twiml/const/region"
	"github.com/homie-dev/gotwiml/twiml/const/status"
	"github.com/homie-dev/gotwiml/twiml/const/trim"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"net/http"
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
		attr.Record(record.FromStart),
		attr.Region(region.AU1),
		attr.Trim(trim.DoNot),
		attr.Coach("sid"),
		attr.StatusCallbackEvent(status.CallbackStart),
		attr.StatusCallback("http://callback.url"),
		attr.StatusCallbackMethod(http.MethodPost),
		attr.RecordingStatusCallback("http://recording.url"),
		attr.RecordingStatusCallbackMethod(http.MethodPost),
		attr.RecordingStatusCallbackEvent(record.StatusCallbackEventAbsent),
	), `<Conference muted="true" beep="onEnter" startConferenceOnEnter="true" endConferenceOnExit="true" participantLabel="pLabel" jitterBufferSize="medium" waitUrl="http://wait.url" waitMethod="POST" maxParticipants="10" record="record-from-start" region="au1" trim="do-not-trim" coach="sid" statusCallbackEvent="start" statusCallback="http://callback.url" statusCallbackMethod="POST" recordingStatusCallback="http://recording.url" recordingStatusCallbackMethod="POST" recordingStatusCallbackEvent="absent">roomName</Conference>`),
)
