package Object

import (
	"Chimp/Ast"
	"bytes"
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ  = "INTEGER"
	BOOL_OBJ     = "BOOL"
	FUNCTION_OBJ = "FUNCTION"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Environment struct {
	parent *Environment
	store  map[string]Object
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		parent: parent,
		store:  map[string]Object{},
	}
}

func (e Environment) Set(key string, obj Object) {
	e.store[key] = obj
}

func (e Environment) Get(key string) (Object, bool) {
	object, ok := e.store[key]
	if !ok {
		if e.parent != nil {
			object, ok = e.parent.Get(key)
		} else {
			return nil, false
		}
	}
	return object, ok
}

type Integer struct {
	Value int64
}

func (i Integer) Type() ObjectType { return INTEGER_OBJ }
func (i Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b Boolean) Type() ObjectType { return BOOL_OBJ }
func (b Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type Function struct {
	Parameters []string
	Body       Ast.BlockStatement
	Env        *Environment
}

func (f Function) Type() ObjectType { return FUNCTION_OBJ }

func (f Function) Inspect() string {
	var params = bytes.Buffer{}
	for i, p := range f.Parameters {
		params.WriteString(p)
		if i+1 != len(f.Parameters) {
			params.WriteString(", ")
		}
	}

	return fmt.Sprintf("(%v) %s", params.String(), f.Body.ToString())
}
