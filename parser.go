package tonho

type Event interface {
}

// Parser is a struct that contains the state of the parser.
//
// It is used to parse a list of tokens to produce a list of events,
// which will be used to generate a concrete syntax tree.
type Parser struct {
	input  string
	index  int
	tokens []Token
	errors []Diagnostic
	events []Event

	// The fuel is the maximum number of tokens that the parser will
	// consume before stopping. This is used to prevent infinite loops
	// in the parser.
	//
	// If the maximum number of tokens is reached, the parser will
	// panic.
	fuel int
}

// NewParser creates a new parser with the given input.
func NewParser(filename, input string) Parser {
	tokens := Lex(filename, input)

	return Parser{input: input, tokens: tokens, fuel: 256}
}
