package dial

type (
	// RecordType is record type
	RecordType string

	// RecordingStatusCallbackEventType is recording status callback event type
	RecordingStatusCallbackEventType string

	// RingToneType is ring tone type
	RingToneType string

	// TrimType is record trim type
	TrimType string
)

// <Dial> attribution keys
const (
	DoNotRecord           RecordType = "do-not-record"
	RecordFromAnswer      RecordType = "record-from-answer"
	RecordFromRinging     RecordType = "record-from-ringing"
	RecordFromAnswerDual  RecordType = "record-from-answer-dual"
	RecordFromRingingDual RecordType = "record-from-ringing-dual"

	RecordingStatusCallbackEventInProgress RecordingStatusCallbackEventType = "in-progress"
	RecordingStatusCallbackEventCompleted  RecordingStatusCallbackEventType = "completed"
	RecordingStatusCallbackEventAbsent     RecordingStatusCallbackEventType = "absent"

	RingToneAt    = RingToneType("at")
	RingToneAu    = RingToneType("au")
	RingToneBg    = RingToneType("bg")
	RingToneBr    = RingToneType("br")
	RingToneBe    = RingToneType("be")
	RingToneCh    = RingToneType("ch")
	RingToneCl    = RingToneType("cl")
	RingToneCn    = RingToneType("cn")
	RingToneCz    = RingToneType("cz")
	RingToneDe    = RingToneType("de")
	RingToneDk    = RingToneType("dk")
	RingToneEe    = RingToneType("ee")
	RingToneEs    = RingToneType("es")
	RingToneFi    = RingToneType("fi")
	RingToneFr    = RingToneType("fr")
	RingToneGr    = RingToneType("gr")
	RingToneHu    = RingToneType("hu")
	RingToneIl    = RingToneType("il")
	RingToneIn    = RingToneType("in")
	RingToneIt    = RingToneType("it")
	RingToneLt    = RingToneType("lt")
	RingToneJp    = RingToneType("jp")
	RingToneMx    = RingToneType("mx")
	RingToneMy    = RingToneType("my")
	RingToneNl    = RingToneType("nl")
	RingToneNo    = RingToneType("no")
	RingToneNz    = RingToneType("nz")
	RingTonePh    = RingToneType("ph")
	RingTonePl    = RingToneType("pl")
	RingTonePt    = RingToneType("pt")
	RingToneRu    = RingToneType("ru")
	RingToneSe    = RingToneType("se")
	RingToneSg    = RingToneType("sg")
	RingToneTh    = RingToneType("th")
	RingToneUk    = RingToneType("uk")
	RingToneUs    = RingToneType("us")
	RingToneUsOld = RingToneType("us-old")
	RingToneTw    = RingToneType("tw")
	RingToneVe    = RingToneType("ve")
	RingToneZa    = RingToneType("za")

	TrimSilence    = TrimType("trim-silence")
	TrimNotSilence = TrimType("do-not-trim")
)
