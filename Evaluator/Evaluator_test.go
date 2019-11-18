package Evaluator

import (
	"Chimp/Lexer"
	"Chimp/Object"
	"Chimp/Parser"
	"testing"
)

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
	return Eval(programme)
}
