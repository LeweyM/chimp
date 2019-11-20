package Object

import (
	"fmt"
)

type ObjectType string

const INTEGER_OBJ = "INTEGER"

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

func(e Environment)Set(key string, obj Object) {
	e.store[key] = obj
}

func(e Environment)Get(key string) (Object, bool) {
	object, ok := e.store[key]
	return object, ok
}

type Integer struct {
	Value int64
}

func (i Integer) Type() ObjectType { return INTEGER_OBJ }
func (i Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }