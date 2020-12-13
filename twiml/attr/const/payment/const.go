package payment

// Method is payment method
type Method string

// payment method types
const (
	ArcDebit   Method = "ach-debit"
	CreditCard Method = "credit-card"
)
