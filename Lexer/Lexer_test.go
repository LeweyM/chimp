package Lexer

import (
	"Chimp/Token"
	"testing"
)

func TestNextToken(t *testing.T) {
	var input = "=;{}(),+"

	var tests = []struct {
		tokenType    Token.TokenType
		tokenLiteral string
	}{
		{Token.EQUALS, "="},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if tt.tokenType != token.Type {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.tokenType, token.Type)
		}
	}
}
