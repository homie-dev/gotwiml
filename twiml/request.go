package twiml

type (
	// VoiceRequest provides the request parameters for TwiML application
	VoiceRequest struct {
		CallSID        string
		AccountSID     string
		ApplicationSID string
		From           string
		To             string
		CallStatus     string
		APIVersion     string
		Direction      string
		ForwardedFrom  string
		CallerName     string
		ParentCallSID  string
		Called         string
		Caller         string
		FromCity       string
		FromState      string
		FromZip        string
		FromCountry    string
		ToCity         string
		ToState        string
		ToZip          string
		ToCountry      string
	}

	// IncomingRequest provides the request parameters for coming call
	IncomingRequest struct {
		AccountSID    string
		APIVersion    string
		CallSID       string
		CallStatus    string
		Called        string
		CalledCity    string
		CalledCountry string
		CalledState   string
		CalledVia     string
		CalledZip     string
		Caller        string
		CallerCity    string
		CallerCountry string
		CallerState   string
		CallerZip     string
		Direction     string
		ForwardedFrom string
		From          string
		FromCity      string
		FromCountry   string
		FromState     string
		FromZip       string
		To            string
		ToCity        string
		ToCountry     string
		ToState       string
		ToZip         string
	}

	// StatusCallbackRequest provides the request parameters for calling status callback
	StatusCallbackRequest struct {
		APIVersion     string
		Called         string
		ParentCallSID  string
		CallStatus     string
		From           string
		Direction      string
		TimeStamp      string
		AccountSID     string
		CallbackSource string
		Caller         string
		SequenceNumber string
		CallSID        string
		To             string
	}

	// RecordingStatusCallbackRequest provides the request parameters for recording status callback
	RecordingStatusCallbackRequest struct {
		CallSID            string
		ErrorCode          int
		RecordingChannels  int
		RecordingDuration  int
		RecordingSource    string
		RecordingSID       string
		RecordingStatus    string
		RecordingStartTime string
		RecordingURL       string
	}

	// DialActionRequest provides the request parameters of dialing result
	DialActionRequest struct {
		AccountSID     string
		APIVersion     string
		ApplicationSID string
		CallSID        string
		CallStatus     string
		Called         string
		Caller         string
		Direction      string
		From           string
		To             string
		DialCallSID    string
		DialCallStatus string
		RecordingSID   string
		RecordingURL   string
	}
)
