package Token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	FUNCTION = "FUN"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	LET      = "LET"
	IDENT    = "IDENT"
	INT      = "INT"

	EQ       = "=="
	NEQ      = "!="
	BANG     = "!"
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	INCR     = "++"
	DECR     = "--"
	DIVIDE   = "/"

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
