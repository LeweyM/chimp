package Lexer

import (
	"Chimp/Token"
	"testing"
)

func TestNextToken(t *testing.T) {
	var input = `
		true false
		=;{}(),+!*/
		monkeySay myVar =   99
		monkeySay plus = monkeyDo(x, y) {
			return x + y
		}
		6 - 5; 10 > 2 < 3
		-- ++
`

	var tests = []struct {
		tokenType    Token.TokenType
		tokenLiteral string
	}{
		{Token.TRUE, "true"},
		{Token.FALSE, "false"},
		{Token.ASSIGN, "="},
		{Token.SEMICOLON, ";"},
		{Token.LPAREN, "{"},
		{Token.RPAREN, "}"},
		{Token.LBRACE, "("},
		{Token.RBRACE, ")"},
		{Token.COMMA, ","},
		{Token.PLUS, "+"},
		{Token.BANG, "!"},
		{Token.MULTIPLY, "*"},
		{Token.DIVIDE, "/"},
		{Token.LET, "monkeySay"},
		{Token.IDENT, "myVar"},
		{Token.ASSIGN, "="},
		{Token.INT, "99"},
		{Token.LET, "monkeySay"},
		{Token.IDENT, "plus"},
		{Token.ASSIGN, "="},
		{Token.FUNCTION, "monkeyDo"},
		{Token.LBRACE, "("},
		{Token.IDENT, "x"},
		{Token.COMMA, ","},
		{Token.IDENT, "y"},
		{Token.RBRACE, ")"},
		{Token.LPAREN, "{"},
		{Token.RETURN, "return"},
		{Token.IDENT, "x"},
		{Token.PLUS, "+"},
		{Token.IDENT, "y"},
		{Token.RPAREN, "}"},
		{Token.INT, "6"},
		{Token.MINUS, "-"},
		{Token.INT, "5"},
		{Token.SEMICOLON, ";"},
		{Token.INT, "10"},
		{Token.GT, ">"},
		{Token.INT, "2"},
		{Token.LT, "<"},
		{Token.INT, "3"},
		{Token.DECR, "--"},
		{Token.INCR, "++"},
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

func TestNumberAndWordLexing(t *testing.T) {
	var input = `
		5;
		100;
		fifty;
`

	var tests = []struct {
		tokenType    Token.TokenType
		tokenLiteral string
	}{
		{Token.INT, "5"},
		{Token.SEMICOLON, ";"},
		{Token.INT, "100"},
		{Token.SEMICOLON, ";"},
		{Token.IDENT, "fifty"},
		{Token.SEMICOLON, ";"},
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

func TestPeekingTokens(t *testing.T) {
	var input = `
		== >= <= !=
`
	var tests = []struct {
		tokenType    Token.TokenType
		tokenLiteral string
	}{
		{Token.EQ, "=="},
		{Token.GTE, ">="},
		{Token.LTE, "<="},
		{Token.NEQ, "!="},
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
