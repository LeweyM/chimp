package Lexer

import (
	"Chimp/Token"
)

var keywords = map[string]Token.TokenType{
	"monkeySay": Token.LET,
	"monkeyDo":  Token.FUNCTION,
	"if":        Token.IF,
	"else":      Token.ELSE,
	"return":    Token.RETURN,
}

type Lexer struct {
	input   string
	readPos int
	curPos  int
	ch      byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.NextToken()
	return l
}

func (l *Lexer) NextToken() Token.Token {

	l.skipWhiteSpaces()

	tok := Token.Token{}

	switch l.ch {
	case '=':
		if peekToken := l.peekToken(); peekToken == '=' {
			tok = newToken(Token.EQ, "==")
			l.readNextChar()
		} else {
			tok = newToken(Token.ASSIGN, "=")
		}
	case '+':
		if peekToken := l.peekToken(); peekToken == '+' {
			tok = newToken(Token.INCR, "++")
			l.readNextChar()
		} else {
			tok = newToken(Token.PLUS, "+")
		}
	case '*':
		tok = newToken(Token.MULTIPLY, "*")
	case '/':
		tok = newToken(Token.DIVIDE, "/")
	case '!':
		if peekToken := l.peekToken(); peekToken == '=' {
			tok = newToken(Token.NEQ, "!=")
			l.readNextChar()
		} else {
			tok = newToken(Token.BANG, "!")
		}
	case '-':
		if peekToken := l.peekToken(); peekToken == '-' {
			tok = newToken(Token.DECR, "--")
			l.readNextChar()
		} else {
			tok = newToken(Token.MINUS, "-")
		}
	case '>':
		if peekToken := l.peekToken(); peekToken == '=' {
			tok = newToken(Token.GTE, ">=")
			l.readNextChar()
		} else {
			tok = newToken(Token.GT, ">")
		}
	case '<':
		if peekToken := l.peekToken(); peekToken == '=' {
			tok = newToken(Token.LTE, "<=")
			l.readNextChar()
		} else {
			tok = newToken(Token.LT, "<")
		}
	case '{':
		tok = newToken(Token.LPAREN, "{")
	case '}':
		tok = newToken(Token.RPAREN, "}")
	case '(':
		tok = newToken(Token.LBRACE, "(")
	case ')':
		tok = newToken(Token.RBRACE, ")")
	case ',':
		tok = newToken(Token.COMMA, ",")
	case ';':
		tok = newToken(Token.SEMICOLON, ";")
	case 0:
		tok = newToken(Token.EOF, "EOF")
	default:
		if isLetter(l.ch) {
			word := l.getWord()
			if keyword := keywords[word]; keyword == "" {
				return newToken(Token.IDENT, word)
			} else {
				return newToken(keyword, word)
			}
		} else if isDigit(l.ch) {
			return newToken(Token.INT, getNumber(l))
		} else {
			tok = newToken(Token.ILLEGAL, "ILLEGAL")
		}
	}

	l.readNextChar()

	return tok
}

func (l *Lexer) peekToken() byte {
	return l.input[l.readPos]
}

func getNumber(l *Lexer) string {
	initialPosition := l.curPos
	for isDigit(l.ch) {
		l.readNextChar()
	}
	return l.input[initialPosition:l.curPos]
}

func isDigit(b byte) bool {
	return b <= '9' && b >= '0'
}

func (l *Lexer) getWord() string {
	initialPosition := l.curPos
	for isLetter(l.ch) {
		l.readNextChar()
	}
	return l.input[initialPosition:l.curPos]
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func newToken(Type Token.TokenType, Literal string) Token.Token {
	return Token.Token{Type: Type, Literal: Literal}
}

func (l *Lexer) readNextChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.curPos = l.readPos
	l.readPos++
}

func (l *Lexer) skipWhiteSpaces() {
	for charIsWhiteSpace(l.ch) {
		l.readNextChar()
	}
}

func charIsWhiteSpace(ch byte) bool {
	return ch == '\n' || ch == '\t' || ch == '\r' || ch == ' '
}
