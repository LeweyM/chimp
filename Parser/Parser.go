package Parser

import (
	"Chimp/Ast"
	"Chimp/Lexer"
	"Chimp/Token"
)

type Parser struct {
	l Lexer.Lexer
}

func (p *Parser)ParseProgramme() Ast.Programme {
	token := p.NextToken()

	switch token.Type {
	case Token.LET:
		ParseLetStatement(token)
	}

	return Ast.Programme{}
}

func ParseLetStatement(token Token.Token) {

}

func New(l Lexer.Lexer) *Parser {
	return &Parser{l}
}

func (p *Parser)NextToken() Token.Token {
	return p.l.NextToken()
}

