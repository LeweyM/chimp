package Evaluator

import (
	"Chimp/Ast"
	"Chimp/Object"
)

func Eval(node Ast.Node, env Object.Environment) Object.Object {
	switch node := node.(type) {
	case Ast.Programme:
		return evalStatements(node.Statements, env)
	case Ast.ExpressionStatement:
		return Eval(node.Value, env)
	case *Ast.LetStatement:
		env.Set(node.Name.Value, Eval(node.Value, env))
		return Eval(node.Value, env)
	case *Ast.IdentityExpression:
		val, ok := env.Get(node.Value)
		if ok {
			return val
		}
	case *Ast.InfixExpression:
		return evalInfix(node, env)
	case *Ast.PrefixExpression:
		return evalPrefix(node, env)
	case *Ast.IntegerExpression:
		return Object.Integer{Value: node.Value}
	//case *Ast.FunctionExpression:
	//	return Object.Function{}
	}

	return nil
}

func evalPrefix(p *Ast.PrefixExpression, env Object.Environment) Object.Object {
	exp := Eval(p.Expression, env)

	switch {
	case exp.Type() == Object.INTEGER_OBJ:
		expInteger := exp.(Object.Integer)
		return evalPrefixInteger(p.Operator, expInteger.Value)
	}

	return nil
}

func evalPrefixInteger(operator string, value int64) Object.Object {
	switch operator {
	case "-":
		return Object.Integer{Value: -value}
	}
	return nil
}

func evalInfix(infix *Ast.InfixExpression, env Object.Environment) Object.Object {
	left := Eval(infix.LeftExpression, env)
	right := Eval(infix.RightExpression, env)

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
	case "*":
		return Object.Integer{Value: leftInteger * rightInteger}
	case "/":
		return Object.Integer{Value: leftInteger / rightInteger}
	}
	return nil
}

func evalStatements(statements []Ast.Statement, env Object.Environment) Object.Object {
	var eval Object.Object

	for _, statement := range statements {
		eval = Eval(statement, env)
	}

	return eval
}
