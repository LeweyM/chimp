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

		p.NextToken()
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

func (p *Parser) ParseLetStatement() *Ast.LetStatement {
	letToken := p.curToken
	if p.NextToken(); p.curToken.Type != Token.IDENT {
		// err
	}
	return &Ast.LetStatement{
		Token: letToken,
		Name:  p.parseIdentExpression(),
		Value: &Ast.IntegerExpression{Token: letToken, Value: int64(999)},
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
	p.NextToken()
	return &p
}

func (p *Parser) NextToken() {
	p.curToken = p.l.NextToken()
}
