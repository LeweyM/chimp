package Parser

import (
	"Chimp/Ast"
	"Chimp/Lexer"
	"Chimp/Token"
	"testing"
)

func TestParseIdentExpressions(t *testing.T) {
	input := `
		foo;
	`
	literals := []string{"foo"}

	l := Lexer.New(input)
	p := New(*l)

	programme := p.ParseProgramme()
	checkForErrors(p, t)

	if len(programme.Statements) != 1 {
		t.Fatalf("Expected 1 statements, got %d", len(programme.Statements))
	}

	for i, statement := range programme.Statements {

		expressionStatement, ok := statement.(Ast.ExpressionStatement)
		if !ok {
			t.Fatal("Not of type ExpressionStatement")
		}

		identExpression, ok := expressionStatement.Value.(*Ast.IdentityExpression)
		if !ok {
			t.Fatal("Not of type ExpressionStatement")
		}

		if identExpression.Token.Type != Token.IDENT {
			t.Fatalf("Expected type of Ident, got %s", identExpression.Token.Type)
		}

		if identExpression.Value != literals[i] {
			t.Fatalf("Expected identifier name of %s, got %s", literals[i], identExpression.Value)
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

func TestParseIfStatements(t *testing.T) {
	input := `
		if (1 < 2) { return 2; } else { return 1; };
		if (5 < 0) { return 8; };
	`
	output := []string{
		"if (1 < 2) { return 2 } else { return 1 }",
		"if (5 < 0) { return 8 }",
	}

	l := Lexer.New(input)
	p := New(*l)

	programme := p.ParseProgramme()
	checkForErrors(p, t)

	if len(programme.Statements) != len(output) {
		t.Fatalf("Expected %d statements, got %d", len(output), len(programme.Statements))
	}

	for i, statement := range programme.Statements {

		ifStatement, ok := statement.(Ast.IfStatement)
		if !ok {
			t.Fatalf("Statement %d Not of type IfStatement", i)
		}

		if ifStatement.ToString() != output[i] {
			t.Fatalf("Statement %d: Expected:\n%s \ngot: \n%s", i, output[i], ifStatement.ToString())
		}
	}
}

func TestParseLetStatements(t *testing.T) {
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

func TestParseInfixExpressions(t *testing.T) {
	input := `
		1 < 2;
		1 + 2 * 3;
		(1 * 2) + 3;
		foo + 5
	`
	output := []string{
		"(1 < 2)",
		"(1 + (2 * 3))",
		"((1 * 2) + 3)",
		"(foo + 5)",
	}

	l := Lexer.New(input)
	p := New(*l)

	programme := p.ParseProgramme()
	checkForErrors(p, t)

	if len(programme.Statements) != len(output) {
		t.Fatalf("Expected %d statements, got %d", len(output), len(programme.Statements))
	}

	for i, statement := range programme.Statements {

		expressionStatement, ok := statement.(Ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement %d Not of type ExpressionStatement", i )
		}

		infixExpression, ok := expressionStatement.Value.(*Ast.InfixExpression)
		if !ok {
			t.Fatalf("statement %d Not of type InfixExpression", i )
		}

		if infixExpression.ToString() != output[i] {
			t.Fatalf("Expected output to be %s, got %s", output[i], infixExpression.ToString())
		}
	}
}

func TestParsePrefixExpressions(t *testing.T) {
	input := `
		-3;
		!5;
		--4;
		++100;
	`
	output := []string{
		"(-3)",
		"(!5)",
		"(--4)",
		"(++100)",
	}

	l := Lexer.New(input)
	p := New(*l)

	programme := p.ParseProgramme()
	checkForErrors(p, t)

	if len(programme.Statements) != len(output) {
		t.Fatalf("Expected %d statements, got %d", len(output), len(programme.Statements))
	}

	for i, statement := range programme.Statements {

		expressionStatement, ok := statement.(Ast.ExpressionStatement)
		if !ok {
			t.Fatal("Not of type ExpressionStatement")
		}

		prefixExpression, ok := expressionStatement.Value.(*Ast.PrefixExpression)
		if !ok {
			t.Fatal("Not of type prefixExpression")
		}

		if prefixExpression.ToString() != output[i] {
			t.Fatalf("Expected output to be %s, got %s", output[i], prefixExpression.ToString())
		}
	}
}

func TestParseFunctionExpressions(t *testing.T) {
	input := `
		fun(x) {
			return x
		};
	`
	output := []string{
		"(x) { return x }",
	}

	l := Lexer.New(input)
	p := New(*l)

	programme := p.ParseProgramme()
	checkForErrors(p, t)

	if len(programme.Statements) != len(output) {
		t.Fatalf("Expected %d statements, got %d", len(output), len(programme.Statements))
	}

	for i, statement := range programme.Statements {

		expressionStatement, ok := statement.(Ast.ExpressionStatement)
		if !ok {
			t.Fatal("Not of type ExpressionStatement")
		}

		functionExpression, ok := expressionStatement.Value.(*Ast.FunctionExpression)
		if !ok {
			t.Fatal("Not of type functionExpression")
		}

		if functionExpression.ToString() != output[i] {
			t.Fatalf("Expected output to be %s, got %s", output[i], functionExpression.ToString())
		}
	}
}

func TestParseFunctionCallExpressions(t *testing.T) {
	input := `
		foo(5);
		foo(51);
	`
	output := []string{
		"foo(5)",
		"foo(51)",
	}

	l := Lexer.New(input)
	p := New(*l)

	programme := p.ParseProgramme()
	checkForErrors(p, t)

	if len(programme.Statements) != len(output) {
		t.Fatalf("Expected %d statements, got %d", len(output), len(programme.Statements))
	}

	for i, statement := range programme.Statements {

		expressionStatement, ok := statement.(Ast.ExpressionStatement)
		if !ok {
			t.Fatal("Not of type ExpressionStatement")
		}

		callExpression, ok := expressionStatement.Value.(*Ast.CallExpression)
		if !ok {
			t.Fatal("Not of type callExpression")
		}

		if callExpression.ToString() != output[i] {
			t.Fatalf("Expected output to be %s, got %s", output[i], callExpression.ToString())
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
