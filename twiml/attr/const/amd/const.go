package amd

// official docs https://www.twilio.com/docs/voice/twiml/number#attributes-machine-detection
type Type string
type Timeout int            // 3-60 seconds. Default is 30 seconds.
type SpeechThreshold int    // 1000-6000 milliseconds. Default is 2400 milliseconds.
type SpeechEndThreshold int // 500-5000	milliseconds. Default is 1200 milliseconds.
type SilenceTimeout int     // 2000-10000 milliseconds. Default is 5000 milliseconds.

// machineDetection types
const (
	Enable           Type = "Enable"
	DetectMessageEnd Type = "DetectMessageEnd"
)
