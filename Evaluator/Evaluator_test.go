package Evaluator

import (
	"Chimp/Lexer"
	"Chimp/Object"
	"Chimp/Parser"
	"fmt"
	"testing"
)

func TestObjectInspect(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"monkeyDo() { return 5; }", "() { return 5 }"},
		{"monkeyDo(x) { return 5; }", "(x) { return 5 }"},
		{"monkeyDo(x ,y) { return 5; }", "(x, y) { return 5 }"},
	}

	for _, tt := range tests {
		if tt.input != tt.expected {
			inspection := evaluateTest(tt.input).Inspect()

			if inspection != tt.expected {
				t.Fatalf(fmt.Sprintf("inspect didn't match, expected: %s, got: %s", tt.expected, inspection))
			}
		}
	}
}

func TestEvalInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"15", 15},
	}

	for _, tt := range tests {
		evaluatedProgramme := evaluateTest(tt.input)

		testInteger(evaluatedProgramme, t, tt.expected)
	}

}

func TestEvalFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"(monkeyDo() { return 5; })()", 5},
		{"(monkeyDo(y) { return y * 2; })(5)", 10},
		{"(monkeyDo(x, y) { return y * x; })(5, 3)", 15},
		//{"monkeySay x = 5; (monkeyDo(x) { return x; })(10); x; ", 5},
	}

	for _, tt := range tests {
		evaluatedProgramme := evaluateTest(tt.input)

		testInteger(evaluatedProgramme, t, tt.expected)
	}

}

func TestInfixInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5 + 1", 6},
		{"5 - 2", 3},
		{"5 * 2 + 2", 12},
		{"10 / 2 + 2", 7},
	}

	for _, tt := range tests {
		evaluatedProgramme := evaluateTest(tt.input)

		testInteger(evaluatedProgramme, t, tt.expected)
	}

}

func TestPrefixInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"-10", -10},
	}

	for _, tt := range tests {
		evaluatedProgramme := evaluateTest(tt.input)

		testInteger(evaluatedProgramme, t, tt.expected)
	}

}

func TestAssignment(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"monkeySay foo = 5; foo", 5},
		{"monkeySay foo = 5;foo + 1", 6},
	}

	for _, tt := range tests {
		evaluatedProgramme := evaluateTest(tt.input)

		testInteger(evaluatedProgramme, t, tt.expected)
	}

}

func testInteger(evaluatedProgramme Object.Object, t *testing.T, expected int64) {
	integerObject, ok := evaluatedProgramme.(Object.Integer)
	if !ok {
		t.Errorf("Object is not integer, is %s", integerObject.Type())
	}
	if integerObject.Value != expected {
		t.Errorf("object has wrong value, expected %d, got %d", expected, integerObject.Value)
	}
}

func evaluateTest(input string) Object.Object {
	l := Lexer.New(input)
	p := Parser.New(*l)
	programme := p.ParseProgramme()
	env := Object.NewEnvironment()

	return Eval(programme, *env)
}
