package Evaluator

import (
	"fmt"
)

func wrongIdentifierErrorMsg(identifier string) string {
	return fmt.Sprintf("Cannot find indentifier '%s'.", identifier)
}
