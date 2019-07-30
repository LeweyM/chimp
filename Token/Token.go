package Token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	EQUALS = "="
	PLUS   = "+"

	LPAREN = "{"
	RPAREN = "}"
	LBRACE = "("
	RBRACE = ")"

	COMMA     = ","
	SEMICOLON = ";"
)
