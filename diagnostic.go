package tonho

// Diagnostic is an interface that represents a diagnostic message.
//
// It is used to report errors, warnings, and other messages.
type Diagnostic interface {
	Kind() int
	Error() []ErrorText
	Location() Location
}

// ErrorText is a struct that represents a diagnostic message.
//
// It is used to report errors, warnings, and other messages.
type ErrorText struct {
	Message string
	kind    int
}

const (
	TextKind = iota
	CodeKind
	NewlineKind
)

const (
	LexerError = iota
	ParserError
	ResolutionError
	TyperError
	UnknownError
	CompilerError
)

// NewText creates a new diagnostic message with the given text
// and returns it
func NewText(text string) ErrorText {
	return ErrorText{Message: text, kind: TextKind}
}

// NewCode creates a new diagnostic message with the given text
// and returns it
func NewCode(text string) ErrorText {
	return ErrorText{Message: text, kind: CodeKind}
}

// NewLine creates a new diagnostic message with the given text
// and returns it
func NewLine() ErrorText {
	return ErrorText{kind: NewlineKind}
}

func (e ErrorText) String() string {
	switch e.kind {
	case TextKind:
		return e.Message
	case CodeKind:
		return "`" + e.Message + "`"
	case NewlineKind:
		return "\n"
	}
	panic("Unknown ErrorText kind")
}
