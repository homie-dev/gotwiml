package trim

// Type is record trim type
type Type string

// trim types
const (
	Silence = Type("trim-silence")
	DoNot   = Type("do-not-trim")
)
