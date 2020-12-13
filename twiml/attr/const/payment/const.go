package payment

// Method is payment method
type Method string

// payment method types
const (
	ArcDebit   Method = "ach-debit"
	CreditCard Method = "credit-card"
)

// ErrorType is error type
type ErrorType string

// error types
const (
	Timeout             ErrorType = "timeout"
	InvalidCardNumber   ErrorType = "invalid-card-number"
	InvalidCardType     ErrorType = "invalid-card-type"
	InvalidDate         ErrorType = "invalid-date"
	InvalidSecurityCode ErrorType = "invalid-security-code"
	InternalError       ErrorType = "internal-error"
)
