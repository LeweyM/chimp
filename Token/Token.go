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
	NEQ    = "!="
	BANG   = "!"
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	GT  = ">"
	LT  = "<"
	GTE = ">="
	LTE = "<="

	LPAREN = "{"
	RPAREN = "}"
	LBRACE = "("
	RBRACE = ")"

	COMMA     = ","
	SEMICOLON = ";"
)
