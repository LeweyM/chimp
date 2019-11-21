package Object

import (
	"Chimp/Ast"
	"fmt"
)

type ObjectType string

const INTEGER_OBJ = "INTEGER"
const FUNCTION_OBJ = "FUNCTION"

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Environment struct {
	store map[string]Object
}

func NewEnvironment() *Environment {
	return &Environment{
		store: map[string]Object{},
	}
}

func (e Environment) Set(key string, obj Object) {
	e.store[key] = obj
}

func (e Environment) Get(key string) (Object, bool) {
	object, ok := e.store[key]
	return object, ok
}

type Integer struct {
	Value int64
}

func (i Integer) Type() ObjectType { return INTEGER_OBJ }
func (i Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Function struct {
	Parameters []string
	Body       Ast.BlockStatement
}

func (f Function) Type() ObjectType { return FUNCTION_OBJ }

func (f Function) Inspect() string { return fmt.Sprintf("(%v) %s", f.Parameters[0], f.Body.ToString())}
