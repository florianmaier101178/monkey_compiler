package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

// identifiers and literals
const (
	IDENT = "IDENT"
	INT   = "INT"
)

// operators
const (
	ASSIGN = "="
	PLUS   = "+"
)

// delimiters
const (
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)

// keywords
const (
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
