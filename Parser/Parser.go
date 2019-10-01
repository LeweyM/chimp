package Parser

import (
	"Chimp/Ast"
	"Chimp/Lexer"
	"Chimp/Token"
	"fmt"
	"strconv"
)

type Parser struct {
	l              Lexer.Lexer
	curToken       Token.Token
	peekToken      Token.Token
	errors         []string
	infixRegistry  map[Token.TokenType]infixFunc
	prefixRegistry map[Token.TokenType]prefixFunc
	precedence     map[string]int
}

type infixFunc = func(left Ast.Expression) *Ast.InfixExpression
type prefixFunc = func() *Ast.PrefixExpression

func New(l Lexer.Lexer) *Parser {
	p := Parser{l: l}
	p.errors = []string{}
	p.AdvanceTokens()
	p.AdvanceTokens()

	p.infixRegistry = make(map[Token.TokenType]infixFunc)
	p.infixRegistry[Token.PLUS] = p.parseInfixExpression
	p.infixRegistry[Token.MINUS] = p.parseInfixExpression
	p.infixRegistry[Token.MULTIPLY] = p.parseInfixExpression
	p.infixRegistry[Token.DIVIDE] = p.parseInfixExpression

	p.prefixRegistry = make(map[Token.TokenType]prefixFunc)
	p.prefixRegistry[Token.MINUS] = p.parsePrefixExpression

	p.precedence = make(map[string]int)
	p.precedence["/"] = MULTI
	p.precedence["*"] = MULTI
	p.precedence["+"] = SUM
	p.precedence["-"] = SUM
	p.precedence["LOWEST"] = LOWEST

	return &p
}

const (
	LOWEST = iota
	SUM
	MULTI
)

func (p *Parser) AdvanceTokens() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) getCurrentToken() Token.Token {
	return p.curToken
}

func (p *Parser) getPeekToken() Token.Token {
	return p.peekToken
}

func (p *Parser) ParseProgramme() Ast.Programme {
	programme := Ast.Programme{}
	programme.Statements = []Ast.Statement{}

	for p.getCurrentToken().Type != Token.EOF {
		if statement := p.parseStatement(); statement != nil {
			programme.Statements = append(programme.Statements, statement)
		}

		p.AdvanceTokens()
	}

	return programme
}

func (p *Parser) parseStatement() Ast.Statement {
	switch p.getCurrentToken().Type {
	case Token.LET:
		return p.ParseLetStatement()
	case Token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpression(contextPrecedence int) Ast.Expression {
	var leftExp Ast.Expression

	prefix, ok := p.prefixRegistry[p.getCurrentToken().Type]
	if !ok {
		leftExp = p.parseIntegerExpression()
	} else {
		leftExp = prefix()
	}

	operatorPres := p.getPeekPrecedence()

	for operatorPres > contextPrecedence {
		infix := p.infixRegistry[p.getPeekToken().Type]
		if infix == nil {
			return leftExp
		}
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) getPeekPrecedence() int {
	return p.precedence[p.getPeekToken().Literal]
}

func (p *Parser) ParseLetStatement() *Ast.LetStatement {
	letToken := p.getCurrentToken()

	if p.AdvanceTokens(); p.getCurrentToken().Type != Token.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected IDENT, but received '%s'", p.getCurrentToken().Literal))
		return nil
	}

	if p.getPeekToken().Type != Token.ASSIGN {
		p.errors = append(p.errors, fmt.Sprintf("expected '=', but received '%s'", p.getCurrentToken().Literal))
		return nil
	}

	statement := &Ast.LetStatement{
		Token: letToken,
		Name:  *p.parseIdentExpression(),
		Value: p.parseExpression(LOWEST),
	}

	p.ignoreUntilSemicolon()

	return statement
}

func (p *Parser) ignoreUntilSemicolon() {
	for p.getCurrentToken().Type != Token.SEMICOLON && p.getCurrentToken().Type != Token.EOF {
		p.AdvanceTokens()
	}
}

func (p *Parser) parseIdentExpression() *Ast.IdentityExpression {
	token := p.getCurrentToken()

	if p.getPeekToken().Type != Token.ASSIGN {
		p.errors = append(p.errors, fmt.Sprintf("expected '=', but received '%s'", p.getCurrentToken().Literal))
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
	returnToken := p.getCurrentToken()

	p.AdvanceTokens()

	valueExpression := p.parseExpression(LOWEST)

	p.ignoreUntilSemicolon()

	return &Ast.ReturnStatement{
		Token: returnToken,
		Value: valueExpression,
	}
}

func (p *Parser) parseIntegerExpression() *Ast.IntegerExpression {
	i, err := strconv.Atoi(p.getCurrentToken().Literal)
	if err != nil {
		p.errors = append(p.errors, "Non number in INT value")
		return nil
	}
	return &Ast.IntegerExpression{Token: p.getCurrentToken(), Value: int64(i)}
}

func (p *Parser) parseInfixExpression(left Ast.Expression) *Ast.InfixExpression {
	operator := p.getPeekToken().Literal
	precedence := p.getPeekPrecedence()

	p.AdvanceTokens()
	p.AdvanceTokens()

	right := p.parseExpression(precedence)

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
		Value: p.parseExpression(LOWEST),
	}

	p.ignoreUntilSemicolon()

	return statement
}

func (p *Parser) parsePrefixExpression() *Ast.PrefixExpression {
	token := p.getCurrentToken()
	p.AdvanceTokens()

	return &Ast.PrefixExpression{
		Token: Token.Token{
			Type:    token.Type,
			Literal: token.Literal,
		},
		Operator:   token.Literal,
		Expression: p.parseIntegerExpression(),
	}

	return nil
}
