package card

// Type is card type
type Type string

// card types
const (
	Visa       Type = "visa"
	Mastercard Type = "mastercard"
	Amex       Type = "amex"
	Maestro    Type = "maestro"
	Discover   Type = "discover"
	Optima     Type = "optima"
	JCB        Type = "jcb"
	DinersClub Type = "diners-club"
	EnRoute    Type = "enroute"
)
