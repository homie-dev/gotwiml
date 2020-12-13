package for_

// Type is Name of the payment source data element
type Type string

// for types
const (
	PaymentCardNumber Type = "payment-card-number"
	ExpirationDate    Type = "expiration-date"
	SecurityCode      Type = "security-code"
	PostalCode        Type = "postal-code"
	PaymentProcessing Type = "payment-processing"
	BankAccountNumber Type = "bank-account-number"
	BankRoutingNumber Type = "bank-routing-number"
)
