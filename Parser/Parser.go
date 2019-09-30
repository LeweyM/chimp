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
		return p.parseExpressionStatement()
		//parse expression statement
		return nil
	}
}

func (p *Parser) parseExpression() Ast.Expression {
	switch p.currentToken().Type {
	case Token.INT:
		if p.getPeekToken().Type == Token.PLUS {
			return p.parseInfixExpression()
		} else {
			return p.parseIntegerExpression()
		}
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

	p.ignoreUntilSemicolon()

	return statement
}

func (p *Parser) ignoreUntilSemicolon() {
	for p.currentToken().Type != Token.SEMICOLON && p.currentToken().Type != Token.EOF {
		p.AdvanceTokens()
	}
}

func (p *Parser) parseIdentExpression() *Ast.IdentityExpression {
	token := p.currentToken()

	if p.getPeekToken().Type != Token.ASSIGN {
		p.errors = append(p.errors, fmt.Sprintf("expected '=', but received '%s'", p.currentToken().Literal))
		return nil
	}

	p.AdvanceTokens()
	p.AdvanceTokens()

	return &Ast.IdentityExpression{
		Token: token,
		Value: token.Literal,
	}
}

func (p *Parser) parseReturnStatement() *Ast.ReturnStatement {
	returnToken := p.currentToken()

	p.AdvanceTokens()

	valueExpression := p.parseExpression()

	p.ignoreUntilSemicolon()

	return &Ast.ReturnStatement{
		Token: returnToken,
		Value: valueExpression,
	}
}

func (p *Parser) parseIntegerExpression() *Ast.IntegerExpression {
	i, err := strconv.Atoi(p.currentToken().Literal)
	if err != nil {
		p.errors = append(p.errors, "Non number in INT value")
		return nil
	}
	return &Ast.IntegerExpression{Token: p.currentToken(), Value: int64(i)}
}

func (p *Parser) parseInfixExpression() *Ast.InfixExpression {
	left := p.parseIntegerExpression()
	p.AdvanceTokens()
	operator := p.currentToken().Literal
	p.AdvanceTokens()
	right := p.parseIntegerExpression()

	return &Ast.InfixExpression{
		Token:           Token.Token{},
		Operator:        operator,
		LeftExpression:  left,
		RightExpression: right,
	}
}

func (p *Parser) parseExpressionStatement() Ast.ExpressionStatement {

	statement := Ast.ExpressionStatement{
		Token: Token.Token{},
		Value: p.parseExpression(),
	}

	p.ignoreUntilSemicolon()

	return statement
}
