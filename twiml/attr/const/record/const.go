package record

// Type is record type
type Type string

// StatusCallbackEventType is recording status callback event type
type StatusCallbackEventType string

// record types
const (
	DoNot           Type = "do-not-record"
	FromAnswer      Type = "record-from-answer"
	FromRinging     Type = "record-from-ringing"
	FromAnswerDual  Type = "record-from-answer-dual"
	FromRingingDual Type = "record-from-ringing-dual"
	FromStart Type = "record-from-start"
)

// recording status callback event
const (
	StatusCallbackEventInProgress StatusCallbackEventType = "in-progress"
	StatusCallbackEventCompleted  StatusCallbackEventType = "completed"
	StatusCallbackEventAbsent     StatusCallbackEventType = "absent"
)
