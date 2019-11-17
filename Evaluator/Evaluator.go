package Evaluator

import (
	"Chimp/Ast"
	"Chimp/Object"
)

func Eval(node Ast.Node) Object.Object {
	switch node := node.(type) {
	case Ast.Programme:
		return evalStatements(node.Statements)
	case Ast.ExpressionStatement:
		return Eval(node.Value)

	case *Ast.IntegerExpression:
		return Object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(statements []Ast.Statement) Object.Object {
	var eval Object.Object

	for _, statement := range statements {
		eval = Eval(statement)
	}

	return eval
}
