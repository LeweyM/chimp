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

	case Ast.InfixExpression:
		return evalInfix(node)
	case *Ast.IntegerExpression:
		return Object.Integer{Value: node.Value}
	}

	return nil
}

func evalInfix(infix Ast.InfixExpression) Object.Object {
	left := Eval(infix.LeftExpression)
	right := Eval(infix.RightExpression)

	switch {
	case left.Type() == Object.INTEGER_OBJ && right.Type() == Object.INTEGER_OBJ:
		leftInteger := left.(Object.Integer)
		rightInteger := right.(Object.Integer)

		return evalInfixInteger(infix.Operator, leftInteger.Value, rightInteger.Value)
	}

	return nil
}

func evalInfixInteger(operator string, leftInteger, rightInteger int64) Object.Object {
	switch operator {
	case "+":
		return Object.Integer{Value: leftInteger + rightInteger}
	case "-":
		return Object.Integer{Value: leftInteger - rightInteger}
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
