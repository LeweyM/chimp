package Parser

import (
	"Chimp/Ast"
	"Chimp/Lexer"
	"Chimp/Token"
	"testing"
)

func TestParseLetStatement(t *testing.T) {
	var tests = []struct {
		input          string
		expectedToken  Token.Token
		expectedName   Ast.Expression
		expectedValue  Ast.Expression
		expectedErrors []string
	}{
		{
			input:         "let five = 5;",
			expectedToken: Token.Token{Type: Token.LET, Literal: "let"},
			expectedName: Ast.IdentityExpression{
				Token: Token.Token{Type: Token.IDENT, Literal: "five"},
				Value: "five",
			},
			expectedValue: Ast.IntegerExpression{
				Token: Token.Token{Type: Token.INT, Literal: "5"},
				Value: 5,
			},
			expectedErrors: []string{},
		}, {
			input:         "let foo = 67",
			expectedToken: Token.Token{Type: Token.LET, Literal: "let"},
			expectedName: Ast.IdentityExpression{
				Token: Token.Token{Type: Token.IDENT, Literal: "foo"},
				Value: "foo",
			},
			expectedValue: Ast.IntegerExpression{
				Token: Token.Token{Type: Token.INT, Literal: "67"},
				Value: 67,
			},
			expectedErrors: []string{},
		},
	}

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

		checkForErrors(p, t)

		if tt.expectedToken.Type != letStatement.Token.Type {
			t.Fatalf("Expected type of %s, got %s", tt.expectedToken.Type, letStatement.Token.Type)
		}

		if tt.expectedToken.Literal != letStatement.Token.Literal {
			t.Fatalf("Expected let literal of %s, got %s", tt.expectedToken.Literal, letStatement.TokenLiteral())
		}

		identityExpression, ok := tt.expectedName.(Ast.IdentityExpression)
		if !ok {
			t.Fatal("Identifier not of type IdentityExpression")
		}

		if identityExpression.Value != letStatement.Name.Value {
			t.Fatalf("Expected identifier name of %s, got %s", tt.expectedName, letStatement.Name)
		}

		expectedIntegerExpression, ok := tt.expectedValue.(Ast.IntegerExpression)
		if !ok {
			t.Fatal("Expected value not of type IntegerExpression")
		}

		integerExpression, ok := letStatement.Value.(*Ast.IntegerExpression)
		if !ok {
			t.Fatal("Value not of type IntegerExpression")
		}

		if expectedIntegerExpression.Value != integerExpression.Value {
			t.Fatalf("Expected value of %s, got %s", tt.expectedValue, letStatement.Value)
		}

	}
}

func TestParseReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 500;
	`

	values := []int64{5, 500}

	l := *Lexer.New(input)
	p := New(l)

	programme := p.ParseProgramme()
	checkForErrors(p, t)

	if len(programme.Statements) != 2 {
		t.Fatalf("expected 2 statments, found %d", len(programme.Statements))
	}

	for i, statement := range programme.Statements {

		returnStatement, ok := statement.(*Ast.ReturnStatement)
		if !ok {
			t.Fatalf("statement %d is not a return statement", i)
		}

		intExpression, ok := returnStatement.Value.(*Ast.IntegerExpression)
		if !ok {
			t.Fatalf("statement %d value is not an integer expression", i)
		}

		if intExpression.Value != values[i] {
			t.Fatalf("statement %d expected value is %d, got %d", i, values[i], intExpression.Value)
		}
	}
}

func checkForErrors(p *Parser, t *testing.T) {
	if len(p.errors) > 0 {
		t.Errorf("%d errors found.\n", len(p.errors))
		for _, msg := range p.errors {
			t.Errorf("Error: %s", msg)
		}
		t.FailNow()
	}
}
