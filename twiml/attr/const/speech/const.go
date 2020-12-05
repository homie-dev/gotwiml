package speech

// Model is speech model
type Model string

//  speech model types
const (
	ModelDefault           Model = "default"
	ModelNumberAndCommands Model = "numbers_and_commands"
	ModelPhoneCall         Model = "phone_call"
)
