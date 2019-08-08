package Parser

import (
	"github.com/LeweyM/chimp/Ast"
	"github.com/LeweyM/chimp/Lexer"
	"github.com/LeweyM/chimp/Token"
)

type Parser struct {
	l        Lexer.Lexer
	curToken Token.Token
}

func (p *Parser) ParseProgramme() Ast.Programme {
	switch p.curToken.Type {
	case Token.LET:
		ParseLetStatement(p.curToken)
	}

	return Ast.Programme{}
}

func ParseLetStatement(token Token.Token) Ast.LetStatement {
	return Ast.LetStatement{
		Token: token,
		Name:  "hello",
		Value: &Ast.IntegerExpression{Token: token, Value: int64(999)},
	}
}

func New(l Lexer.Lexer) *Parser {
	p := Parser{l: l}
	p.NextToken()
	return &p
}

func (p *Parser) NextToken() {
	p.curToken = p.l.NextToken()
}
