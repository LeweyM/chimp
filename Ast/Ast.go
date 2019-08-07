package Ast

import "Chimp/Token"

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

type LetStatement struct {
	Token Token.Token
	Name string
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}
func (ls *LetStatement) statementNode()       {}