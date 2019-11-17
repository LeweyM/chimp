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
		l := Lexer.New(tt.input)
		p := Parser.New(*l)
		programme := p.ParseProgramme()
		evaluatedProgramme := Eval(programme)

		integerObject, ok := evaluatedProgramme.(Object.Integer)

		if !ok {
			t.Errorf("Object is not integer, is %s", integerObject.Type())
		}

		if integerObject.Value != tt.expected {
			t.Errorf("object has wrong value, expected %d, got %d", tt.expected, integerObject.Value)
		}

	}

}
