package tonho

import (
	"fmt"
	"unicode"
)

// Token represents a token in the source code.
type Token struct {
	Kind int
	Text string

	// Location represents the location of the
	// token in the source code.
	location Location

	// Represents the token with the ignored
	// tokens present. Like, if the token has
	// any spaces, or comments, or newlines,
	// it will appear here, but not on Text.
	FullText string
}

// lexerLocation represents a location in the
// source code.
type lexerLocation struct {
	start int
	end   int
	text  string
	file  string
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

// keywords is used to determine if an
// identifier is a keyword or not.
var keywords = map[string]int{
	"fun":   Fun,
	"val":   Val,
	"var":   Var,
	"for":   For,
	"while": While,
	"loop":  Loop,
	"if":    If,
	"else":  Else,
	"when":  When,
}

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

// lexer represents a scanner that will recognize
// tokens in the source code.
type lexer struct {
	filename, input string
	tokens          []Token
	position        int
	start           int
}

// Lex creates a new lexer with the given input.
func Lex(filename, input string) []Token {
	l := lexer{filename: filename, input: input}
	return l.lex()
}

// NewToken creates a new token with the given
// kind, text and full text.
//
// The text is the text that is used to create
// the token. The full text is the text that
// is used to create the token, but with the
// ignored tokens present.
func NewToken(kind int, text, fullText string) Token {
	return Token{
		Kind:     kind,
		Text:     text,
		FullText: fullText,
	}
}

// String returns the string representation of
// the token.
func (t Token) String() string {
	return fmt.Sprintf("Token (kind: '%s', text: '%s')", names[t.Kind], t.Text)
}

// Location returns the location of the token.
func (t Token) Location() Location {
	return t.location
}

// Start returns the start position of the
// token.
func (l lexerLocation) Start() int {
	return l.start
}

// End returns the end position of the
// token.
func (l lexerLocation) End() int {
	return l.end
}

// Text returns the text of the token.
func (l lexerLocation) Text() string {
	return l.text
}

// File returns the file name of the token.
func (l lexerLocation) File() string {
	return l.file
}

// lex scans the input and returns the tokens
// that were found.
func (l *lexer) lex() []Token {
	for {
		l.start = l.position

		if l.position >= len(l.input) {
			l.tokens = append(l.tokens, l.newToken(EOF))
			break
		}

		if !l.nextToken() {
			break
		}
	}
	return l.tokens
}

func (l *lexer) nextToken() bool {
	switch c := l.peek(); c {
	case ' ', '\t', '\r':
	case '%':
		l.tokens = append(l.tokens, l.newToken(Percent))
	case '.':
		l.tokens = append(l.tokens, l.newToken(Dot))
	case ',':
		l.tokens = append(l.tokens, l.newToken(Comma))
	case ';':
		l.tokens = append(l.tokens, l.newToken(Comma))
	case '{':
		l.tokens = append(l.tokens, l.newToken(LeftBrace))
	case '}':
		l.tokens = append(l.tokens, l.newToken(RightBrace))
	case '[':
		l.tokens = append(l.tokens, l.newToken(LeftBracket))
	case ']':
		l.tokens = append(l.tokens, l.newToken(RightBracket))
	case '(':
		l.tokens = append(l.tokens, l.newToken(LeftParen))
	case ')':
		l.tokens = append(l.tokens, l.newToken(RightParen))
	case '+':
		l.tokens = append(l.tokens, l.newToken(Plus))
	case '/':
		l.tokens = append(l.tokens, l.newToken(Slash))
	case '*':
		l.tokens = append(l.tokens, l.newToken(Asterisk))
	case '-':
		lookahead := l.lookahead(1)
		switch lookahead {
		case '>':
			l.tokens = append(l.tokens, l.newToken(Arrow))
			l.advance(2)
			return true
		default:
			l.tokens = append(l.tokens, l.newToken(Minus))
		}
		l.tokens = append(l.tokens, l.newToken(Minus))
	case '|':
		lookahead := l.lookahead(1)
		switch lookahead {
		case '|':
			l.tokens = append(l.tokens, l.newToken(Or))
			l.advance(2)
			return true
		}
	case '&':
		lookahead := l.lookahead(1)
		switch lookahead {
		case '&':
			l.tokens = append(l.tokens, l.newToken(And))
			l.advance(2)
			return true
		}
	case '!':
		lookahead := l.lookahead(1)
		switch lookahead {
		case '=':
			l.tokens = append(l.tokens, l.newToken(NotEqual))
			l.advance(2)
			return true
		default:
			l.tokens = append(l.tokens, l.newToken(Equal))
		}
	case '>':
		lookahead := l.lookahead(1)
		switch lookahead {
		case '=':
			l.tokens = append(l.tokens, l.newToken(GreaterEqual))
			l.advance(2)
			return true
		default:
			l.tokens = append(l.tokens, l.newToken(Greater))
		}
	case '<':
		lookahead := l.lookahead(1)
		switch lookahead {
		case '=':
			l.tokens = append(l.tokens, l.newToken(LessEqual))
			l.advance(2)
			return true
		default:
			l.tokens = append(l.tokens, l.newToken(Less))
		}
	case '=':
		lookahead := l.lookahead(1)
		switch lookahead {
		case '=':
			l.tokens = append(l.tokens, l.newToken(Equal))
			l.advance(2)
			return true
		default:
			l.tokens = append(l.tokens, l.newToken(Assign))
		}
	default:
		if unicode.IsLetter(c) {
			return l.lexIdentifier()
		} else if unicode.IsDigit(c) {
			return l.lexNumber()
		} else if c == '"' {
			return l.lexString()
		}
		l.tokens = append(l.tokens, l.newToken(Error))
	}
	l.position++
	return true
}

// newToken creates a new token with the given
// lexer state and kind.
func (l *lexer) newToken(kind int) Token {
	text := l.input[l.start:l.position]

	// build the token
	token := NewToken(kind, text, text)
	token.location = l.location()
	return token
}

// lexIdentifier scans the input and returns
// the identifier token.
func (l *lexer) lexIdentifier() bool {
	l.advance(1) // skip the first letter

	for !l.eof() && isIdentifierSegment(l.peek()) {
		l.advance(1)
	}

	identifier := l.input[l.start:l.position]

	// Check if the identifier is a keyword.
	if keyword, ok := keywords[identifier]; ok {
		l.tokens = append(l.tokens, l.newToken(keyword))
	} else {
		l.tokens = append(l.tokens, l.newToken(Identifier))
	}

	return true
}

// lexString scans the input and returns
// the string token.
func (l *lexer) lexString() bool {
	l.advance(1) // skip the first quote
	for !l.eof() && l.peek() != '"' {
		l.advance(1)
	}
	l.advance(1)

	// build the token of string
	text := l.input[l.start+1 : l.position-1]
	fullText := l.input[l.start:l.position]
	token := NewToken(String, text, fullText)
	token.location = l.location()

	l.tokens = append(l.tokens, token)
	return true
}

// lexNumber scans the input and returns
// the number token.
//
// The number token can be a decimal or
// an integer.
func (l *lexer) lexNumber() bool {
	l.advance(1) // skip the first digit
	for !l.eof() && unicode.IsDigit(l.peek()) {
		l.advance(1)

		if !l.eof() && l.peek() == '.' {
			l.advance(1)
			for !l.eof() && unicode.IsDigit(l.peek()) {
				l.advance(1)
			}
			l.tokens = append(l.tokens, l.newToken(Decimal))
			break
		}
	}
	l.tokens = append(l.tokens, l.newToken(Int))
	return true
}

// advance advances the lexer position.
func (l *lexer) advance(amount int) {
	l.position += amount
}

// peek returns the rune that is
// at the lexer position.
func (l *lexer) peek() rune {
	return rune(l.input[l.position])
}

// eof returns true if the lexer
// position is at the end of the
// input.
func (l *lexer) eof() bool {
	return l.position >= len(l.input)
}

// lookahead returns the rune that is
// ahead of the lexer position.
func (l *lexer) lookahead(amount int) rune {
	return rune(l.input[l.position+amount])
}

func (l *lexer) location() Location {
	return lexerLocation{
		start: l.start,
		end:   l.position,
		text:  l.input,
		file:  l.filename,
	}
}

// isIdentifierSegment returns true if the
// given rune is a valid identifier segment.
func isIdentifierSegment(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '.' || r == '\''
}
