package Evaluator

import (
	"fmt"
)

func wrongIdentifierErrorMsg(identifier string) string {
	return fmt.Sprintf("Cannot find indentifier '%s'.", identifier)
}

func unknownFunctionErrorMsg(funcName string) string {
	return fmt.Sprintf("Could not find function '%s'", funcName)
}

func invalidInfixOperation(left string, right string, op string) string {
	return fmt.Sprintf("Invalid infix operation: Cannot use '%s' with '%s' and '%s'", op, left, right)
}