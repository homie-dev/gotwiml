package attr

import (
	"strconv"
	"strings"

	"github.com/homie-dev/gotwiml/twiml/attr/const/bank"
	"github.com/homie-dev/gotwiml/twiml/attr/const/beep"
	"github.com/homie-dev/gotwiml/twiml/attr/const/card"
	"github.com/homie-dev/gotwiml/twiml/attr/const/for_"
	"github.com/homie-dev/gotwiml/twiml/attr/const/http"
	"github.com/homie-dev/gotwiml/twiml/attr/const/input"
	"github.com/homie-dev/gotwiml/twiml/attr/const/jitter"
	"github.com/homie-dev/gotwiml/twiml/attr/const/language"
	"github.com/homie-dev/gotwiml/twiml/attr/const/payment"
	"github.com/homie-dev/gotwiml/twiml/attr/const/reason"
	"github.com/homie-dev/gotwiml/twiml/attr/const/record"
	"github.com/homie-dev/gotwiml/twiml/attr/const/region"
	"github.com/homie-dev/gotwiml/twiml/attr/const/ring"
	"github.com/homie-dev/gotwiml/twiml/attr/const/speech"
	"github.com/homie-dev/gotwiml/twiml/attr/const/status"
	"github.com/homie-dev/gotwiml/twiml/attr/const/token"
	"github.com/homie-dev/gotwiml/twiml/attr/const/trim"
	"github.com/homie-dev/gotwiml/twiml/attr/const/voice"
	"github.com/homie-dev/gotwiml/twiml/core"
)

// Option is function to set XML attribute
type Option func(core.XMLer)

const (
	_action                        = "action"
	_actionOnEmptyResult           = "actionOnEmptyResult"
	_answerOnBridge                = "answerOnBridge"
	_attempt                       = "attempt"
	_bankAccountType               = "bankAccountType"
	_beep                          = "beep"
	_byoc                          = "byoc"
	_callerID                      = "callerId"
	_cardType                      = "cardType"
	_chargeAmount                  = "chargeAmount"
	_coach                         = "coach"
	_connectorName                 = "connectorName"
	_currency                      = "currency"
	_description                   = "description"
	_digits                        = "digits"
	_endConferenceOnExit           = "endConferenceOnExit"
	_enhanced                      = "enhanced"
	_errorType                     = "errorType"
	_finishOnKey                   = "finishOnKey"
	_for                           = "for"
	_hangupOnStar                  = "hangupOnStar"
	_hints                         = "hints"
	_input                         = "input"
	_jitterBufferSize              = "jitterBufferSize"
	_language                      = "language"
	_length                        = "length"
	_loop                          = "loop"
	_maxAttempts                   = "maxAttempts"
	_maxLength                     = "maxLength"
	_maxParticipants               = "maxParticipants"
	_method                        = "method"
	_minPostalCodeLength           = "minPostalCodeLength"
	_muted                         = "muted"
	_name                          = "name"
	_numDigits                     = "numDigits"
	_partialResultCallback         = "partialResultCallback"
	_partialResultCallbackMethod   = "partialResultCallbackMethod"
	_participantIdentity           = "participantIdentity"
	_participantLabel              = "participantLabel"
	_paymentConnector              = "paymentConnector"
	_paymentMethod                 = "paymentMethod"
	_password                      = "password"
	_playBeep                      = "playBeep"
	_postalCode                    = "postalCode"
	_postWorkActivitySID           = "postWorkActivitySid"
	_profanityFilter               = "profanityFilter"
	_reason                        = "reason"
	_record                        = "record"
	_recordingStatusCallback       = "recordingStatusCallback"
	_recordingStatusCallbackEvent  = "recordingStatusCallbackEvent"
	_recordingStatusCallbackMethod = "recordingStatusCallbackMethod"
	_region                        = "region"
	_reservationSID                = "reservationSid"
	_ringTone                      = "ringTone"
	_sendDigits                    = "sendDigits"
	_securityCode                  = "securityCode"
	_speechModel                   = "speechModel"
	_speechTimeout                 = "speechTimeout"
	_startConferenceOnEnter        = "startConferenceOnEnter"
	_statusCallback                = "statusCallback"
	_statusCallbackEvent           = "statusCallbackEvent"
	_statusCallbackMethod          = "statusCallbackMethod"
	_timeLimit                     = "timeLimit"
	_timeout                       = "timeout"
	_tokenType                     = "tokenType"
	_track                         = "track"
	_transcribe                    = "transcribe"
	_transcribeCallback            = "transcribeCallback"
	_trim                          = "trim"
	_url                           = "url"
	_username                      = "username"
	_validCardTypes                = "validCardTypes"
	_value                         = "value"
	_voice                         = "voice"
	_waitMethod                    = "waitMethod"
	_waitURL                       = "waitUrl"
	_waitURLMethod                 = "waitUrlMethod"
	_workflowSID                   = "workflowSid"
)

// Action sets action url
func Action(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_action, v)
	}
}

// ActionOnEmptyResult sets to force webhook to the action URL event if there is no input
func ActionOnEmptyResult(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_actionOnEmptyResult, strconv.FormatBool(v))
	}
}

// AnswerOnBridge preserve the ringing behavior of the inbound call until the Dialed call picks up
func AnswerOnBridge(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_answerOnBridge, strconv.FormatBool(v))
	}
}

// Attempt sets name of the payment source data element
func Attempt(vv ...int) Option {
	ss := make([]string, len(vv))
	for i, v := range vv {
		ss[i] = strconv.Itoa(v)
	}
	return func(t core.XMLer) {
		t.SetAttr(_attempt, strings.Join(ss, " "))
	}
}

// BankAccountType sets bank account type for ach transactions. If set, payment method attribute must be provided and value should be set to ach-debit. defaults to consumer-checking
func BankAccountType(v bank.AccountType) Option {
	return func(t core.XMLer) {
		t.SetAttr(_bankAccountType, string(v))
	}
}

// Beep sets to Play beep when joining
func Beep(v beep.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_beep, string(v))
	}
}

// BYOC sets byoc trunk SID
func BYOC(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_byoc, v)
	}
}

// CallerID sets Caller ID to display
func CallerID(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_callerID, v)
	}
}

// CardType sets type of the credit card
func CardType(vv ...card.Type) Option {
	ss := make([]string, len(vv))
	for i, v := range vv {
		ss[i] = string(v)
	}
	return func(t core.XMLer) {
		t.SetAttr(_cardType, strings.Join(ss, " "))
	}
}

// ChargeAmount sets amount to process. If value is greater than 0 then make the payment else create a payment token
func ChargeAmount(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_chargeAmount, v)
	}
}

// Coach sets to accept a call SID
func Coach(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_coach, v)
	}
}

// ConnectorName sets unique name for Connector
func ConnectorName(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_connectorName, v)
	}
}

// Currency sets currency of the amount attribute
func Currency(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_currency, v)
	}
}

// Description sets details regarding the payment
func Description(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_description, v)
	}
}

// Digits sets play DTMF tones for digits
func Digits(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_digits, v)
	}
}

// EndConferenceOnExit sets to accept a call SID
func EndConferenceOnExit(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_endConferenceOnExit, strconv.FormatBool(v))
	}
}

// Enhanced sets to use enhanced speech model
func Enhanced(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_enhanced, strconv.FormatBool(v))
	}
}

// ErrorType sets type of error
func ErrorType(vv ...payment.ErrorType) Option {
	ss := make([]string, len(vv))
	for i, v := range vv {
		ss[i] = string(v)
	}
	return func(t core.XMLer) {
		t.SetAttr(_errorType, strings.Join(ss, " "))
	}
}

// FinishOnKey sets finish gather on key
func FinishOnKey(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_finishOnKey, v)
	}
}

// For sets finish gather on key
func For(v for_.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_for, string(v))
	}
}

// HangupOnStar Hangup call on star press
func HangupOnStar(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_hangupOnStar, strconv.FormatBool(v))
	}
}

// Hints sets speech recognition hints
func Hints(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_hints, v)
	}
}

// Input sets input type Twilio should accept
func Input(vv ...input.Type) Option {
	ss := make([]string, len(vv))
	for i, v := range vv {
		ss[i] = string(v)
	}
	return func(t core.XMLer) {
		t.SetAttr(_input, strings.Join(ss, " "))
	}
}

// JitterBufferSize sets jitter buffer behavior for a conference participant
func JitterBufferSize(v jitter.BufferSize) Option {
	return func(t core.XMLer) {
		t.SetAttr(_jitterBufferSize, string(v))
	}
}

// Language sets message language
func Language(v language.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_language, string(v))
	}
}

// Length sets length in seconds to pause
func Length(v int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_length, strconv.Itoa(v))
	}
}

// Loop sets times to loop message
func Loop(v language.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_loop, string(v))
	}
}

// MaxAttempts sets maximum number of allowed retries when gathering input
func MaxAttempts(v int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_maxAttempts, strconv.Itoa(v))
	}
}

// MaxLength sets max time to record in seconds
func MaxLength(v int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_maxLength, strconv.Itoa(v))
	}
}

// MaxParticipants sets maximum number of participants
func MaxParticipants(v int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_maxParticipants, strconv.Itoa(v))
	}
}

// MinPostalCodeLength sets prompt for minimum postal code length
func MinPostalCodeLength(v int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_minPostalCodeLength, strconv.Itoa(v))
	}
}

// Method sets http method with request action url
func Method(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_method, string(v))
	}
}

// Muted sets to join the conference muted
func Muted(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_muted, strconv.FormatBool(v))
	}
}

// Name sets the name of the custom parameter
func Name(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_name, v)
	}
}

// NumDigits sets number of digits to collect
func NumDigits(v int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_numDigits, strconv.Itoa(v))
	}
}

// PartialResultCallback sets partial result callback URL
func PartialResultCallback(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_partialResultCallback, v)
	}
}

// PartialResultCallbackMethod sets partial result callback URL method
func PartialResultCallbackMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_partialResultCallbackMethod, string(v))
	}
}

// ParticipantIdentity sets a unique identity on the incoming caller
func ParticipantIdentity(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_participantIdentity, v)
	}
}

// ParticipantLabel sets a label for the conference participant
func ParticipantLabel(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_participantLabel, v)
	}
}

// PaymentConnector sets unique name for payment connector
func PaymentConnector(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_paymentConnector, v)
	}
}

// PaymentMethod sets payment method to be used. defaults to credit-card
func PaymentMethod(v payment.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_paymentMethod, string(v))
	}
}

// PlayBeep sets play beep
func PlayBeep(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_playBeep, strconv.FormatBool(v))
	}
}

// Password sets SIP password
func Password(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_password, v)
	}
}

// PostalCode sets prompt for postal code and it should be true/false or default postal code
func PostalCode(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_postalCode, v)
	}
}

// PostWorkActivitySID sets TaskRouter Activity SID
func PostWorkActivitySID(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_postWorkActivitySID, v)
	}
}

// ProfanityFilter sets profanity filter on speech
func ProfanityFilter(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_profanityFilter, strconv.FormatBool(v))
	}
}

// Reason sets rejection reason
func Reason(r reason.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_reason, string(r))
	}
}

// Record sets record type of the call
func Record(r record.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_record, string(r))
	}
}

// RecordingStatusCallback is Recording status callback URL
func RecordingStatusCallback(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_recordingStatusCallback, v)
	}
}

// RecordingStatusCallbackEvent is Recording status callback URL method
func RecordingStatusCallbackEvent(v record.StatusCallbackEventType) Option {
	return func(t core.XMLer) {
		t.SetAttr(_recordingStatusCallbackEvent, string(v))
	}
}

// RecordingStatusCallbackMethod is Recording status callback URL method
func RecordingStatusCallbackMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_recordingStatusCallbackMethod, string(v))
	}
}

// Region sets the region where Twilio should mix the conference
func Region(v region.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_region, string(v))
	}
}

// ReservationSID sets TaskRouter Reservation SID
func ReservationSID(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_reservationSID, v)
	}
}

// RingTone sets to allow you to override the ringback tone that Twilio will play back to the caller while executing the Dial
func RingTone(v ring.ToneType) Option {
	return func(t core.XMLer) {
		t.SetAttr(_ringTone, string(v))
	}
}

// SendDigits sets DTMF tones to play when the call is answered
func SendDigits(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_sendDigits, v)
	}
}

// SecurityCode sets prompt for security code
func SecurityCode(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_securityCode, strconv.FormatBool(v))
	}
}

// SpeechTimeout sets time to wait to gather speech input and it should be either auto or a positive integer
func SpeechTimeout(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_speechTimeout, v)
	}
}

// SpeechModel sets specify the model that is best suited for your use case
func SpeechModel(v speech.Model) Option {
	return func(t core.XMLer) {
		t.SetAttr(_speechModel, string(v))
	}
}

// StartConferenceOnEnter sets to start the conference on enter
func StartConferenceOnEnter(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_startConferenceOnEnter, strconv.FormatBool(v))
	}
}

// StatusCallback sets status call back url
func StatusCallback(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_statusCallback, v)
	}
}

// StatusCallbackEvent sets status call back url
func StatusCallbackEvent(vv ...status.CallbackEvent) Option {
	ss := make([]string, len(vv))
	for i, v := range vv {
		ss[i] = string(v)
	}
	return func(t core.XMLer) {
		t.SetAttr(_statusCallbackEvent, strings.Join(ss, " "))
	}
}

// StatusCallbackMethod sets status call back method
func StatusCallbackMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_statusCallbackMethod, string(v))
	}
}

// TimeLimit is Max time length
func TimeLimit(sec int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_timeLimit, strconv.Itoa(sec))
	}
}

// Timeout is Time to wait for answer
func Timeout(sec int) Option {
	return func(t core.XMLer) {
		t.SetAttr(_timeout, strconv.Itoa(sec))
	}
}

// TokenType sets type of token
func TokenType(v token.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_tokenType, string(v))
	}
}

// Track sets track to be streamed to remote service
func Track(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_track, v)
	}
}

// Transcribe sets transcribe the recording
func Transcribe(v bool) Option {
	return func(t core.XMLer) {
		t.SetAttr(_transcribe, strconv.FormatBool(v))
	}
}

// TranscribeCallback sets transcribe callback URL
func TranscribeCallbackURL(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_transcribeCallback, v)
	}
}

// Trim is Time to wait for answer
func Trim(v trim.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_trim, string(v))
	}
}

// URL sets url
func URL(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_url, v)
	}
}

// UserName sets SIP user name
func UserName(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_username, v)
	}
}

// ValidCardTypes sets comma separated accepted card types
func ValidCardTypes(v card.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_validCardTypes, string(v))
	}
}

// Value sets the value of the custom parameter
func Value(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_value, v)
	}
}

// Voice sets voice to use
func Voice(v voice.Type) Option {
	return func(t core.XMLer) {
		t.SetAttr(_voice, string(v))
	}
}

// WaitMethod sets wait URL method
func WaitMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_waitMethod, string(v))
	}
}

// WaitURL sets wait URL
func WaitURL(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_waitURL, string(v))
	}
}

// WaitURLMethod sets wait URL method
func WaitURLMethod(v http.Method) Option {
	return func(t core.XMLer) {
		t.SetAttr(_waitURLMethod, string(v))
	}
}

// WorkflowSID sets wait URL method
func WorkflowSID(v string) Option {
	return func(t core.XMLer) {
		t.SetAttr(_workflowSID, v)
	}
}
