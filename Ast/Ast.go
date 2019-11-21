package Ast

import (
	"Chimp/Token"
	"bytes"
	"fmt"
	"strings"
)

type Programme struct {
	Statements []Statement
}

func (p Programme) ToString() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.ToString())
	}

	return out.String()
}
func (p Programme) TokenLiteral() string {
	return "" //TODO
}

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
	ToString() string
}
type Expression interface {
	Node
	expressionNode()
	ToString() string
}

type IntegerExpression struct {
	Token Token.Token
	Value int64
}

func (ie IntegerExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie IntegerExpression) expressionNode()      {}
func (ie IntegerExpression) ToString() string {
	return ie.Token.Literal
}

type InfixExpression struct {
	Token           Token.Token
	Operator        string
	LeftExpression  Expression
	RightExpression Expression
}

func (ie InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie InfixExpression) expressionNode()      {}
func (ie InfixExpression) ToString() string {
	return fmt.Sprintf("(%s %s %s)",
		ie.LeftExpression.ToString(),
		ie.Operator,
		ie.RightExpression.ToString())
}

type PrefixExpression struct {
	Token      Token.Token
	Operator   string
	Expression Expression
}

func (pe PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe PrefixExpression) expressionNode()      {}
func (pe PrefixExpression) ToString() string {
	return fmt.Sprintf("(%s%s)", pe.Operator, pe.Expression.ToString())
}

type IdentityExpression struct {
	Token Token.Token
	Value string
}

func (ie IdentityExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie IdentityExpression) expressionNode()      {}
func (ie IdentityExpression) ToString() string {
	return fmt.Sprintf("%s", ie.Value, )
}

type ExpressionStatement struct {
	Token Token.Token
	Value Expression
}

func (ls ExpressionStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls ExpressionStatement) statementNode()       {}
func (ls ExpressionStatement) ToString() string {
	return ls.Value.ToString()
}

type LetStatement struct {
	Token Token.Token
	Name  IdentityExpression
	Value Expression
}

func (ls LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls LetStatement) statementNode()       {}
func (ls LetStatement) ToString() string {
	return fmt.Sprintf("%v = %v", ls.Name.Value, ls.Value.ToString())
}

type ReturnStatement struct {
	Token Token.Token
	Value Expression
}

func (rs ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs ReturnStatement) statementNode()       {}
func (rs ReturnStatement) ToString() string {
	return fmt.Sprintf("return %s", rs.Value.ToString())
}

type IfStatement struct {
	Token     Token.Token
	Condition Expression
	Then      Statement
	Else      Statement
}

func (is IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is IfStatement) statementNode()       {}
func (is IfStatement) ToString() string {
	res := fmt.Sprintf("if %s %s", is.Condition.ToString(), is.Then.ToString())
	if is.Else.ToString() != "" {
		res += fmt.Sprintf(" else %s", is.Else.ToString())
	}
	return res
}

type BlockStatement struct {
	Token      Token.Token
	Statements []Statement
}

func (bs BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs BlockStatement) statementNode()       {}
func (bs BlockStatement) ToString() string {
	if len(bs.Statements) == 0 {
		return ""
	}
	builder := strings.Builder{}
	builder.WriteString("{ ")
	for _, s := range bs.Statements {
		builder.WriteString(s.ToString())
	}
	builder.WriteString(" }")
	return builder.String()
}

type FunctionExpression struct {
	Token      Token.Token
	Parameters []IdentityExpression
	Body       BlockStatement
}

func (f FunctionExpression) TokenLiteral() string {
	panic("implement me")
}
func (f FunctionExpression) expressionNode() {
	panic("implement me")
}
func (f FunctionExpression) ToString() string {
	return fmt.Sprintf("(%v) %v", f.Parameters[0].ToString(), f.Body.ToString())
}

type CallExpression struct {
	Token      Token.Token
	Target     Expression
	Parameters []Expression
}

func (c CallExpression) TokenLiteral() string {
	panic("implement me")
}
func (c CallExpression) expressionNode() {
	panic("implement me")
}
func (c CallExpression) ToString() string {
	return fmt.Sprintf("fun%s(%v)", c.Target.ToString(), c.Parameters[0].ToString())
}
