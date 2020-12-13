package bank

// AccountType is bank account type
type AccountType string

// bank account types
const (
	AccountConsumerChecking   AccountType = "consumer-checking"
	AccountConsumerSavings    AccountType = "consumer-savings"
	AccountCommercialChecking AccountType = "commercial-checking"
	AccountCommercialSavings  AccountType = "commercial-savings"
)
