package Parser

import (
	"Chimp/Ast"
	"Chimp/Lexer"
	"Chimp/Token"
	"testing"
)

func TestParseLetStatement(t *testing.T) {
	input := `
		let foo = 67;
		let five = 5;
	`
	literals := []string{"foo", "five"}
	values := []int64{67, 5}

	l := Lexer.New(input)
	p := New(*l)

	programme := p.ParseProgramme()
	checkForErrors(p, t)

	if len(programme.Statements) != 2 {
		t.Fatalf("Expected 2 statements, got %d", len(programme.Statements))
	}

	for i, statement := range programme.Statements {

		letStatement, ok := statement.(*Ast.LetStatement)
		if !ok {
			t.Fatal("Not of type LetStatement")
		}

		if letStatement.Token.Type != Token.LET {
			t.Fatalf("Expected type of LET, got %s", letStatement.Token.Type)
		}

		if letStatement.Name.Value != literals[i] {
			t.Fatalf("Expected identifier name of %s, got %s", literals[i], letStatement.Name)
		}

		integerExpression, ok := letStatement.Value.(*Ast.IntegerExpression)
		if !ok {
			t.Fatal("not an integer expression")
		}

		if integerExpression.Value != values[i] {
			t.Fatalf("Expected value of %d, got %d", values[i], letStatement.Value)
		}

	}

}

func TestParseInfixExpression(t *testing.T) {
	input := `
		5 + 1;
		8 + 5
	`
	left := []string{"5", "8"}
	right := []string{"1", "5"}
	infix := []string{"+", "+"}

	l := Lexer.New(input)
	p := New(*l)

	programme := p.ParseProgramme()
	checkForErrors(p, t)

	if len(programme.Statements) != 2 {
		t.Fatalf("Expected 2 statements, got %d", len(programme.Statements))
	}

	for i, statement := range programme.Statements {

		expressionStatement, ok := statement.(Ast.ExpressionStatement)
		if !ok {
			t.Fatal("Not of type ExpressionStatement")
		}

		infixExpression, ok := expressionStatement.Value.(*Ast.InfixExpression)
		if !ok {
			t.Fatal("Not of type infixExpression")
		}

		if infixExpression.LeftExpression.TokenLiteral() != left[i] {
			t.Fatalf("Expected left to be %s, got %s", left[i], infixExpression.LeftExpression.TokenLiteral())
		}
		if infixExpression.RightExpression.TokenLiteral() != right[i] {
			t.Fatalf("Expected right to be %s, got %s", right[i], infixExpression.RightExpression.TokenLiteral())
		}
		if infixExpression.Operator != infix[i] {
			t.Fatalf("Expected operator to be %s, got %s", infix[i], infixExpression.Operator)
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
