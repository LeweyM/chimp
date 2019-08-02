package Lexer

import (
	"Chimp/Token"
	"testing"
)

func TestNextToken(t *testing.T) {
	var input = `
		=;{}(),+
		let foo =    5
		let myVar = 99
		let plus = fun(x, y) {
			x + y
		}
`

	var tests = []struct {
		tokenType    Token.TokenType
		tokenLiteral string
	}{
		{Token.EQUALS, "="},
		{Token.SEMICOLON, ";"},
		{Token.LPAREN, "{"},
		{Token.RPAREN, "}"},
		{Token.LBRACE, "("},
		{Token.RBRACE, ")"},
		{Token.COMMA, ","},
		{Token.PLUS, "+"},
		{Token.LET, "let"},
		{Token.IDENT, "foo"},
		{Token.EQUALS, "="},
		{Token.INT, "5"},
		{Token.LET, "let"},
		{Token.IDENT, "myVar"},
		{Token.EQUALS, "="},
		{Token.INT, "99"},
		{Token.LET, "let"},
		{Token.IDENT, "plus"},
		{Token.EQUALS, "="},
		{Token.FUNCTION, "fun"},
		{Token.LBRACE, "("},
		{Token.IDENT, "x"},
		{Token.COMMA, ","},
		{Token.IDENT, "y"},
		{Token.RBRACE, ")"},
		{Token.LPAREN, "{"},
		{Token.IDENT, "x"},
		{Token.PLUS, "+"},
		{Token.IDENT, "y"},
		{Token.RPAREN, "}"},
		{Token.EOF, "EOF"},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if tt.tokenType != token.Type {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.tokenType, token.Type)
		}

		if tt.tokenLiteral != token.Literal {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", i, tt.tokenLiteral, token.Literal)
		}
	}
}
