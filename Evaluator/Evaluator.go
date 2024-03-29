package Evaluator

import (
	"Chimp/Ast"
	"Chimp/Object"
	"errors"
)

func Eval(node Ast.Node, env *Object.Environment) (obj Object.Object, err error) {
	if err != nil {
		return nil, err
	}
	switch node := node.(type) {
	case Ast.Programme:
		return evalStatements(node.Statements, env)
	case Ast.ExpressionStatement:
		return Eval(node.Value, env)
	case *Ast.LetStatement:
		object, _ := Eval(node.Value, env)
		env.Set(node.Name.Value, object)
		return Eval(node.Value, env)
	case *Ast.IdentityExpression:
		val, ok := env.Get(node.Value)
		if ok {
			return val, nil
		} else {
			return nil, errors.New(wrongIdentifierErrorMsg(node.Value))
		}
	case *Ast.InfixExpression:
		return evalInfix(node, env)
	case Ast.BlockStatement:
		return evalStatements(node.Statements, env)
	case *Ast.ReturnStatement:
		return Eval(node.Value, env)
	case Ast.IfStatement:
		object, _ := Eval(node.Condition, env)
		boolExpression := object.(Object.Boolean)

		if boolExpression.Value {
			return Eval(node.Then, env)
		} else {
			return Eval(node.Else, env)
		}
	case *Ast.PrefixExpression:
		return evalPrefix(node, env), nil
	case *Ast.IntegerExpression:
		return Object.Integer{Value: node.Value}, nil
	case *Ast.BoolExpression:
		return Object.Boolean{Value: node.Value}, nil
	case *Ast.FunctionExpression:
		return evalFunction(node, env), nil
	case *Ast.CallExpression:
		return evalCall(node, env)
	}

	return nil, nil
}

func evalFunction(node *Ast.FunctionExpression, env *Object.Environment) Object.Object {
	var params []string
	for _, p := range node.Parameters {
		params = append(params, p.ToString())
	}
	return Object.Function{
		Parameters: params,
		Body:       node.Body,
		Env:        env,
	}
}

func evalCall(node *Ast.CallExpression, env *Object.Environment) (obj Object.Object, err error) {
	targetObject, _ := Eval(node.Target, env)
	function, ok := targetObject.(Object.Function)
	if !ok {
		return nil, errors.New(unknownFunctionErrorMsg(node.Target.ToString()))
	}

	extendedScope := Object.NewEnvironment(function.Env)
	for i, paramValue := range node.Parameters {
		paramObjectValue, err := Eval(paramValue, env)
		if err != nil {
			panic("this shouldn't happen")
		}
		extendedScope.Set(function.Parameters[i], paramObjectValue)
	}
	return Eval(function.Body, extendedScope)
}

func evalPrefix(p *Ast.PrefixExpression, env *Object.Environment) Object.Object {
	exp, _ := Eval(p.Expression, env)

	switch {
	case exp.Type() == Object.INTEGER_OBJ:
		expInteger := exp.(Object.Integer)
		return evalPrefixInteger(p.Operator, expInteger.Value)
	case exp.Type() == Object.BOOL_OBJ:
		expBool := exp.(Object.Boolean)
		return evalPrefixBool(p.Operator, expBool.Value)
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

func evalPrefixBool(operator string, value bool) Object.Object {
	switch operator {
	case "!":
		return Object.Boolean{Value: !value}
	}
	return nil
}

func evalInfix(infix *Ast.InfixExpression, env *Object.Environment) (Object.Object, error) {
	left, _ := Eval(infix.LeftExpression, env)
	right, _ := Eval(infix.RightExpression, env)

	switch {
	case left.Type() == Object.INTEGER_OBJ && right.Type() == Object.INTEGER_OBJ:
		leftInteger := left.(Object.Integer)
		rightInteger := right.(Object.Integer)
		return evalInfixInteger(infix.Operator, leftInteger.Value, rightInteger.Value), nil
	}

	return nil, errors.New(invalidInfixOperation(left.Inspect(), right.Inspect(), infix.Operator))
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
	case ">":
		return Object.Boolean{Value: leftInteger > rightInteger}
	case ">=":
		return Object.Boolean{Value: leftInteger >= rightInteger}
	case "<":
		return Object.Boolean{Value: leftInteger < rightInteger}
	case "<=":
		return Object.Boolean{Value: leftInteger <= rightInteger}
	case "==":
		return Object.Boolean{Value: leftInteger == rightInteger}
	case "!=":
		return Object.Boolean{Value: leftInteger != rightInteger}
	}
	return nil
}

func evalStatements(statements []Ast.Statement, env *Object.Environment) (Object.Object, error) {
	var (
		eval Object.Object
		err  error
	)

	for _, statement := range statements {
		eval, err = Eval(statement, env)
	}

	return eval, err
}
