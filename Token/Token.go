package Token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	FUNCTION = "FUN"
	LET      = "LET"
	IDENT    = "IDENT"
	INT      = "INT"

	EQ     = "=="
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	GT = ">"
	LT = "<"

	LPAREN = "{"
	RPAREN = "}"
	LBRACE = "("
	RBRACE = ")"

	COMMA     = ","
	SEMICOLON = ";"
)
