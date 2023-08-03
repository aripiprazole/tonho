package tonho

import "fmt"

// Token represents a token in the source code.
type Token struct {
	Kind int
	Text string

	// Represents the token with the ignored
	// tokens present. Like, if the token has
	// any spaces, or comments, or newlines,
	// it will appear here, but not on Text.
	FullText string
}

// Token kinds. This defines the tokens that
// are recognized by the lexer.
const (
	EOF = iota
	Error

	Identifier
	Decimal
	Int
	String

	Fun
	Val
	Var
	For
	While
	Loop
	If
	Else
	When

	Plus
	Minus
	Asterisk
	Slash
	Percent
	Equal
	NotEqual
	Less
	LessEqual
	Greater
	GreaterEqual
	And
	Or
	Not
	Assign

	LeftParen
	RightParen
	LeftBrace
	RightBrace
	LeftBracket
	RightBracket

	Comma
	Dot
	Colon
	Semi
	Arrow

	Comment
	Newline
)

// Token names. This is used for debugging
// purposes.
var names = map[int]string{
	EOF:   "EOF",
	Error: "Error",

	Identifier: "Identifier",
	Decimal:    "Decimal",
	Int:        "Int",
	String:     "String",

	Fun:   "fun",
	Val:   "val",
	Var:   "var",
	For:   "for",
	While: "while",
	Loop:  "loop",
	If:    "if",
	Else:  "else",
	When:  "when",

	Plus:         "+",
	Minus:        "-",
	Asterisk:     "*",
	Slash:        "/",
	Percent:      "%",
	Equal:        "==",
	NotEqual:     "!=",
	Less:         "<",
	LessEqual:    "<=",
	Greater:      ">",
	GreaterEqual: ">=",
	And:          "&&",
	Or:           "||",
	Not:          "!",
	Assign:       "=",

	LeftParen:    "(",
	RightParen:   ")",
	LeftBrace:    "[",
	RightBrace:   "]",
	LeftBracket:  "{",
	RightBracket: "}",

	Comma: ",",
	Dot:   ".",
	Colon: ":",
	Semi:  ";",
	Arrow: "->",

	Comment: "//",
	Newline: "\\n",
}

// NewToken creates a new token with the given
// kind, text and full text.
//
// The text is the text that is used to create
// the token. The full text is the text that
// is used to create the token, but with the
// ignored tokens present.
func NewToken(kind int, text, fullText string) *Token {
	return &Token{
		Kind:     kind,
		Text:     text,
		FullText: fullText,
	}
}

// String returns the string representation of
// the token.
func (t Token) String() string {
	return fmt.Sprintf("%s (%q)", names[t.Kind], t.Text)
}
