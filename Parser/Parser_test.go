package Parser

import (
	"Chimp/Ast"
	"Chimp/Lexer"
	"Chimp/Token"
	"testing"
)

func TestParseLetStatement(t *testing.T) {
	var tests = []struct {
		input         string
		expectedToken Token.Token
		expectedValue Ast.Expression
		expectedName  Ast.Expression
	}{
		{
			input:         "let five = 5",
			expectedToken: Token.Token{Type: Token.LET, Literal: "let"},
			expectedName: Ast.IdentityExpression{
				Token: Token.Token{Type: Token.IDENT, Literal: "five"},
				Value: "five",
			},
		}}

	for _, tt := range tests {
		l := Lexer.New(tt.input)
		p := New(*l)

		programme := p.ParseProgramme()

		if len(programme.Statements) != 1 {
			t.Fatalf("wrong number of statements in programme; wanted 1, got %d", len(programme.Statements))
		}

		statement := programme.Statements[0]

		letStatement, ok := statement.(*Ast.LetStatement)
		if !ok {
			t.Fatal("Not of type LetStatement")
		}

		if tt.expectedToken.Type != letStatement.Token.Type {
			t.Fatalf("Expected type of %s, got %s", tt.expectedToken.Type, letStatement.Token.Type)
		}

		if tt.expectedToken.Literal != letStatement.Token.Literal {
			t.Fatalf("Expected literal of %s, got %s", tt.expectedToken.Literal, letStatement.TokenLiteral())
		}

		identityExpression, ok := tt.expectedName.(Ast.IdentityExpression)
		if !ok {
			t.Fatal("Identifier not of type IdentityExpression")
		}

		if identityExpression.Value != letStatement.Name.Value {
			t.Fatalf("Expected identifier name of %s, got %s", tt.expectedName, letStatement.Name)
		}

	}
}
