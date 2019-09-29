package Parser

import (
	"Chimp/Ast"
	"Chimp/Lexer"
	"Chimp/Token"
	"fmt"
	"strconv"
)

type Parser struct {
	l         Lexer.Lexer
	curToken  Token.Token
	peekToken Token.Token
	errors    []string
}

func New(l Lexer.Lexer) *Parser {
	p := Parser{l: l}
	p.errors = []string{}
	p.AdvanceTokens()
	p.AdvanceTokens()
	return &p
}

func (p *Parser) AdvanceTokens() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) currentToken() Token.Token {
	return p.curToken
}

func (p *Parser) getPeekToken() Token.Token {
	return p.peekToken
}

func (p *Parser) ParseProgramme() Ast.Programme {
	programme := Ast.Programme{}
	programme.Statements = []Ast.Statement{}

	for p.currentToken().Type != Token.EOF {
		if statement := p.parseStatement(); statement != nil {
			programme.Statements = append(programme.Statements, statement)
		}

		p.AdvanceTokens()
	}

	return programme
}

func (p *Parser) parseStatement() Ast.Statement {
	switch p.currentToken().Type {
	case Token.LET:
		return p.ParseLetStatement()
	case Token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseExpression() Ast.Expression {
	p.AdvanceTokens()
	switch p.currentToken().Type {
	case Token.INT:
		i, err := strconv.Atoi(p.currentToken().Literal)
		if err != nil {
			p.errors = append(p.errors, "Non number in INT value")
			return nil
		}
		expression := &Ast.IntegerExpression{Token: p.currentToken(), Value: int64(i)}
		return expression
	default:
		return nil
	}
}

func (p *Parser) ParseLetStatement() *Ast.LetStatement {
	letToken := p.currentToken()

	if p.AdvanceTokens(); p.currentToken().Type != Token.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected IDENT, but received '%s'", p.currentToken().Literal))
		return nil
	}

	if p.getPeekToken().Type != Token.ASSIGN {
		p.errors = append(p.errors, fmt.Sprintf("expected '=', but received '%s'", p.currentToken().Literal))
		return nil
	}

	statement := &Ast.LetStatement{
		Token: letToken,
		Name:  *p.parseIdentExpression(),
		Value: p.parseExpression(),
	}
	return statement
}

func (p *Parser) parseIdentExpression() *Ast.IdentityExpression {
	token := p.currentToken()

	if p.AdvanceTokens(); p.currentToken().Type != Token.ASSIGN {
		p.errors = append(p.errors, fmt.Sprintf("expected '=', but received '%s'", p.currentToken().Literal))
		return nil
	}

	return &Ast.IdentityExpression{
		Token: token,
		Value: token.Literal,
	}
}

func (p *Parser) parseReturnStatement() *Ast.ReturnStatement {
	returnToken := p.currentToken()

	valueExpression := p.parseExpression()

	return &Ast.ReturnStatement{
		Token: returnToken,
		Value: valueExpression,
	}
}
