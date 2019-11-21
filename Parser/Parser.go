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

type infixFunc = func(left Ast.Expression) Ast.Expression
type prefixFunc = func() Ast.Expression

func New(l Lexer.Lexer) *Parser {
	p := Parser{l: l}
	p.errors = []string{}
	p.advanceTokens()
	p.advanceTokens()
	p.infixRegistry = make(map[Token.TokenType]infixFunc)
	p.infixRegistry[Token.LBRACE] = p.parseCallExpression
	p.infixRegistry[Token.GT] = p.parseInfixExpression
	p.infixRegistry[Token.GTE] = p.parseInfixExpression
	p.infixRegistry[Token.LT] = p.parseInfixExpression
	p.infixRegistry[Token.LTE] = p.parseInfixExpression
	p.infixRegistry[Token.EQ] = p.parseInfixExpression
	p.infixRegistry[Token.NEQ] = p.parseInfixExpression
	p.infixRegistry[Token.PLUS] = p.parseInfixExpression
	p.infixRegistry[Token.MINUS] = p.parseInfixExpression
	p.infixRegistry[Token.MULTIPLY] = p.parseInfixExpression
	p.infixRegistry[Token.DIVIDE] = p.parseInfixExpression

	p.prefixRegistry = make(map[Token.TokenType]prefixFunc)
	p.prefixRegistry[Token.FUNCTION] = p.parseFunctionExpression
	p.prefixRegistry[Token.IDENT] = p.parseIdentExpression
	p.prefixRegistry[Token.BANG] = p.parsePrefixExpression
	p.prefixRegistry[Token.MINUS] = p.parsePrefixExpression
	p.prefixRegistry[Token.INCR] = p.parsePrefixExpression
	p.prefixRegistry[Token.DECR] = p.parsePrefixExpression
	p.prefixRegistry[Token.LBRACE] = p.parseBracePrefixExpression

	p.precedence = make(map[string]int)
	p.precedence["/"] = MULTI
	p.precedence["*"] = MULTI
	p.precedence["+"] = SUM
	p.precedence["-"] = SUM
	p.precedence[">"] = EQUALS
	p.precedence[">="] = EQUALS
	p.precedence["<"] = EQUALS
	p.precedence["<="] = EQUALS
	p.precedence["=="] = EQUALS
	p.precedence["!="] = EQUALS
	p.precedence["LOWEST"] = LOWEST

	return &p
}

const (
	LOWEST = iota
	EQUALS
	SUM
	MULTI
)

func (p *Parser) ParseProgramme() Ast.Programme {
	programme := Ast.Programme{}
	programme.Statements = []Ast.Statement{}

	for p.getCurrentToken().Type != Token.EOF {
		if statement := p.parseStatement(); statement != nil {
			programme.Statements = append(programme.Statements, statement)
		}

		p.advanceTokens()
	}

	return programme
}

func (p *Parser) parseStatement() Ast.Statement {
	switch p.getCurrentToken().Type {
	case Token.LET:
		return p.parseLetStatement()
	case Token.LPAREN:
		return p.parseBlockStatement()
	case Token.RETURN:
		return p.parseReturnStatement()
	case Token.IF:
		return p.parseIfStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpression(contextPrecedence int) Ast.Expression {
	var leftExp Ast.Expression

	prefix, ok := p.prefixRegistry[p.getCurrentToken().Type]
	if !ok {
		leftExp = p.parseLiteral()
	} else {
		leftExp = prefix()
	}

	p.advanceTokens()

	operatorPres := p.getCurrentPrecedence()

	for operatorPres >= contextPrecedence {
		infix := p.infixRegistry[p.getCurrentToken().Type]
		if infix == nil {
			return leftExp
		}
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseLetStatement() *Ast.LetStatement {
	letToken := p.getCurrentToken()

	if p.advanceTokens(); p.getCurrentToken().Type != Token.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected IDENT, but received '%s'", p.getCurrentToken().Literal))
		return nil
	}

	identityExpression := *p.parseIdentExpression().(*Ast.IdentityExpression)

	if p.getPeekToken().Type != Token.ASSIGN {
		p.errors = append(p.errors, fmt.Sprintf("expected '=', but received '%s'", p.getCurrentToken().Literal))
		return nil
	}
	p.advanceTokens()
	p.advanceTokens()
	valueExpression := p.parseExpression(LOWEST)

	statement := &Ast.LetStatement{
		Token: letToken,
		Name:  identityExpression,
		Value: valueExpression,
	}

	p.ignoreUntilSemicolon()

	return statement
}

func (p *Parser) parseReturnStatement() *Ast.ReturnStatement {
	returnToken := p.getCurrentToken()

	p.advanceTokens()

	valueExpression := p.parseExpression(LOWEST)

	p.ignoreUntilSemicolon()

	return &Ast.ReturnStatement{
		Token: returnToken,
		Value: valueExpression,
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

func (p *Parser) parseBlockStatement() Ast.BlockStatement {
	token := p.getCurrentToken()
	p.advanceTokens()

	var statements []Ast.Statement
	for p.getCurrentToken().Type != Token.RPAREN && p.getCurrentToken().Type != Token.EOF {
		statements = append(statements, p.parseStatement())
		p.advanceTokens()
	}

	p.advanceTokens()

	return Ast.BlockStatement{
		Token:      token,
		Statements: statements,
	}
}

func (p *Parser) parseIfStatement() Ast.IfStatement {
	token := p.getCurrentToken()
	p.advanceTokens()

	condition := p.parseExpression(LOWEST)

	thenStatement := p.parseBlockStatement()

	var elseStatement Ast.BlockStatement

	if p.getCurrentToken().Type == Token.ELSE {
		p.advanceTokens()
		elseStatement = p.parseBlockStatement()
	}

	return Ast.IfStatement{
		Token:     token,
		Condition: condition,
		Then:      thenStatement,
		Else:      elseStatement,
	}
}

func (p *Parser) parseIdentExpression() Ast.Expression {
	token := p.getCurrentToken()

	return &Ast.IdentityExpression{
		Token: token,
		Value: token.Literal,
	}
}

func (p *Parser) parseLiteral() Ast.Expression {
	switch p.getCurrentToken().Type {
	case Token.INT:
		return p.parseIntegerExpression()
	}
	panic("cannot parse literal")
}

func (p *Parser) parseIntegerExpression() *Ast.IntegerExpression {
	i, err := strconv.Atoi(p.getCurrentToken().Literal)
	if err != nil {
		p.errors = append(p.errors, "Non number in INT value")
		return nil
	}
	return &Ast.IntegerExpression{Token: p.getCurrentToken(), Value: int64(i)}
}

func (p *Parser) parseFunctionExpression() Ast.Expression {
	token := p.getCurrentToken()

	p.advanceTokens()
	p.advanceTokens()

	//TODO multi parameters
	parameters := p.parseIdentExpression().(*Ast.IdentityExpression)

	p.advanceTokens()
	p.advanceTokens()

	body := p.parseBlockStatement()

	functionExpression := Ast.FunctionExpression{
		Token:      token,
		Parameters: []Ast.IdentityExpression{*parameters},
		Body:       body,
	}
	return &functionExpression
}

func (p *Parser) parseCallExpression(left Ast.Expression) Ast.Expression {
	p.advanceTokens()

	//TODO multi parameters
	parameters := p.parseExpression(LOWEST)

	callExpression := Ast.CallExpression{
		Token:      Token.Token{},
		Target:     left,
		Parameters: []Ast.Expression{parameters},
	}
	return &callExpression
}

func (p *Parser) parseInfixExpression(left Ast.Expression) Ast.Expression {
	operator := p.getCurrentToken().Literal
	precedence := p.getCurrentPrecedence()

	p.advanceTokens()

	right := p.parseExpression(precedence)

	// last token ends at right-expression
	return &Ast.InfixExpression{
		Token:           Token.Token{},
		Operator:        operator,
		LeftExpression:  left,
		RightExpression: right,
	}
}

func (p *Parser) parsePrefixExpression() Ast.Expression {
	token := p.getCurrentToken()
	p.advanceTokens()

	return &Ast.PrefixExpression{
		Token: Token.Token{
			Type:    token.Type,
			Literal: token.Literal,
		},
		Operator:   token.Literal,
		Expression: p.parseIntegerExpression(),
	}
	// last token handled pos by parseIntegerExpression
}

func (p *Parser) parseBracePrefixExpression() Ast.Expression {
	p.advanceTokens()

	expression := p.parseExpression(LOWEST)

	if p.getCurrentToken().Type != Token.RBRACE {
		return nil
	}

	return expression
	//last token pos at right brace
}

func (p *Parser) ignoreUntilSemicolon() {
	for p.getCurrentToken().Type != Token.SEMICOLON && p.getCurrentToken().Type != Token.EOF {
		p.advanceTokens()
	}
}

func (p *Parser) advanceTokens() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) getCurrentToken() Token.Token {
	return p.curToken
}

func (p *Parser) getPeekToken() Token.Token {
	return p.peekToken
}

func (p *Parser) getPeekPrecedence() int {
	return p.precedence[p.getPeekToken().Literal]
}

func (p *Parser) getCurrentPrecedence() int {
	return p.precedence[p.getCurrentToken().Literal]
}
