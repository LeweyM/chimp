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
