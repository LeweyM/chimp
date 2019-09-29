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

type IntegerExpression struct {
	Token Token.Token
	Value int64
}

func (ie IntegerExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie IntegerExpression) expressionNode()      {}

type IdentityExpression struct {
	Token Token.Token
	Value string
}

func (ie IdentityExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie IdentityExpression) expressionNode()      {}

type LetStatement struct {
	Token Token.Token
	Name  IdentityExpression
	Value Expression
}

func (ls LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls LetStatement) statementNode()       {}

type ReturnStatement struct {
	Token Token.Token
	Value Expression
}

func (rs ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs ReturnStatement) statementNode()          {}
