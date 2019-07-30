package Lexer

import "Chimp/Token"

type Lexer struct {
	input string
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (*Lexer) NextToken() Token.Token {
	return Token.Token{
		Type:    Token.EQUALS,
		Literal: "=",
	}
}
