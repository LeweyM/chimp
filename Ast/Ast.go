package Ast

import "github.com/LeweyM/chimp/Token"

type Programme struct {
	Statements []Statement
}

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type IntegerExpression struct {
	Token Token.Token
	Value int64
}

func (ie *IntegerExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IntegerExpression) expressionNode()      {}

type LetStatement struct {
	Token Token.Token
	Name  string
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) statementNode()       {}
