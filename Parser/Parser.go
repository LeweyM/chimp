package Parser

import (
	"Chimp/Ast"
	"Chimp/Lexer"
	"Chimp/Token"
)

type Parser struct {
	l        Lexer.Lexer
	curToken Token.Token
}

func (p *Parser) ParseProgramme() Ast.Programme {
	programme := Ast.Programme{}
	programme.Statements = []Ast.Statement{}

	for p.curToken.Type != Token.EOF {
		if statement := p.parseStatement(); statement != nil {
			programme.Statements = append(programme.Statements, statement)
		}

		p.AdvanceTokens()
	}

	return programme
}

func (p *Parser) parseStatement() Ast.Statement {
	switch p.curToken.Type {
	case Token.LET:
		return p.ParseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseExpression() Ast.Expression {
	switch p.curToken.Type {
	case Token.INT:
		return &Ast.IntegerExpression{Token: p.curToken, Value: int64(999)}
	default:
		return nil
	}
}

func (p *Parser) ParseLetStatement() *Ast.LetStatement {
	letToken := p.curToken
	if p.AdvanceTokens(); p.curToken.Type != Token.IDENT {
		// err
	}
	return &Ast.LetStatement{
		Token: letToken,
		Name:  p.parseIdentExpression(),
		Value: p.parseExpression(),
	}
}

func (p *Parser) parseIdentExpression() Ast.IdentityExpression {
	token := p.curToken
	return Ast.IdentityExpression{
		Token: token,
		Value: token.Literal,
	}
}

func New(l Lexer.Lexer) *Parser {
	p := Parser{l: l}
	p.AdvanceTokens()
	return &p
}

func (p *Parser) AdvanceTokens() {
	p.curToken = p.l.NextToken()
}
