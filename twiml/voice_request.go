package twiml

type (
	// VoiceRequest provides the request parameters for TwiML application
	VoiceRequest struct {
		CallSID        string `form:"callSid"`
		AccountSID     string `form:"accountSid"`
		ApplicationSID string `form:"applicationSid"`
		From           string `form:"from"`
		To             string `form:"to"`
		CallStatus     string `form:"callStatus"`
		APIVersion     string `form:"apiVersion"`
		Direction      string `form:"direction"`
		ForwardedFrom  string `form:"forwardedFrom"`
		CallerName     string `form:"callerName"`
		ParentCallSID  string `form:"parentCallSid"`
		Called         string `form:"called"`
		Caller         string `form:"caller"`
		FromCity       string `form:"fromCity"`
		FromState      string `form:"fromState"`
		FromZip        string `form:"fromZip"`
		FromCountry    string `form:"fromCountry"`
		ToCity         string `form:"toCity"`
		ToState        string `form:"toState"`
		ToZip          string `form:"toZip"`
		ToCountry      string `form:"toCountry"`
	}

	// IncomingRequest provides the request parameters for coming call
	IncomingRequest struct {
		AccountSID    string `form:"accountSid"`
		APIVersion    string `form:"apiVersion"`
		CallSID       string `form:"callSid"`
		CallStatus    string `form:"callStatus"`
		Called        string `form:"called"`
		CalledCity    string `form:"calledCity"`
		CalledCountry string `form:"calledCountry"`
		CalledState   string `form:"calledState"`
		CalledVia     string `form:"calledVia"`
		CalledZip     string `form:"calledZip"`
		Caller        string `form:"caller"`
		CallerCity    string `form:"callerCity"`
		CallerCountry string `form:"callerCountry"`
		CallerState   string `form:"callerState"`
		CallerZip     string `form:"callerZip"`
		Direction     string `form:"direction"`
		ForwardedFrom string `form:"forwardedFrom"`
		From          string `form:"from"`
		FromCity      string `form:"fromCity"`
		FromCountry   string `form:"fromCountry"`
		FromState     string `form:"fromState"`
		FromZip       string `form:"fromZip"`
		To            string `form:"to"`
		ToCity        string `form:"toCity"`
		ToCountry     string `form:"toCountry"`
		ToState       string `form:"toState"`
		ToZip         string `form:"toZip"`
	}

	// StatusCallbackRequest provides the request parameters for calling status callback
	StatusCallbackRequest struct {
		APIVersion     string `form:"apiVersion"`
		Called         string `form:"called"`
		ParentCallSID  string `form:"parentCallSid"`
		CallStatus     string `form:"callStatus"`
		From           string `form:"from"`
		Direction      string `form:"direction"`
		TimeStamp      string `form:"timeStamp"`
		AccountSID     string `form:"accountSid"`
		CallbackSource string `form:"callbackSource"`
		Caller         string `form:"caller"`
		SequenceNumber string `form:"sequenceNumber"`
		CallSID        string `form:"callSid"`
		To             string `form:"to"`
	}

	// RecordingStatusCallbackRequest provides the request parameters for recording status callback
	RecordingStatusCallbackRequest struct {
		CallSID            string `form:"callSid"`
		ErrorCode          int    `form:"errorCode"`
		RecordingChannels  int    `form:"recordingChannels"`
		RecordingDuration  int    `form:"recordingDuration"`
		RecordingSource    string `form:"recordingSource"`
		RecordingSID       string `form:"recordingSid"`
		RecordingStatus    string `form:"recordingStatus"`
		RecordingStartTime string `form:"recordingStartTime"`
		RecordingURL       string `form:"recordingUrl"`
	}

	// DialActionRequest provides the request parameters of dialing result
	DialActionRequest struct {
		AccountSID       string `form:"accountSid"`
		APIVersion       string `form:"apiVersion"`
		ApplicationSID   string `form:"applicationSid"`
		CallSID          string `form:"callSid"`
		CallStatus       string `form:"callStatus"`
		Called           string `form:"called"`
		Caller           string `form:"caller"`
		Direction        string `form:"direction"`
		From             string `form:"from"`
		To               string `form:"to"`
		DialCallSID      string `form:"dialCallSid"`
		DialCallDuration int    `form:"dialCallDuration"`
		DialCallStatus   string `form:"dialCallStatus"`
		RecordingSID     string `form:"recordingSid"`
		RecordingURL     string `form:"recordingUrl"`
	}

	// TODO
	// EnqueueActionRequest provides
	EnqueueActionRequest struct {
	}

	// TODO
	// WaitURLRequest provides
	WaitURLRequest struct {
	}
)
