package lexer

import (
	"unicode"

	"github.com/arps18/andromex/token"
)

type Lexer struct {
	input        string
	pos, readPos int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) NextToken() token.Token {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

	var tok token.Token
	switch l.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case '+':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case '-':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case '*':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case '/':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case '(':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case ')':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case ';':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if unicode.IsLetter(rune(l.ch)) {
			literal := l.readIdentifier()
			if literal == "let" {
				tok.Type = token.LET
			} else if literal == "print" {
				tok.Type = token.PRINT
			} else {
				tok.Type = token.IDENT
			}
			tok.Literal = literal
			return tok
		} else if unicode.IsDigit(rune(l.ch)) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.Token{
				Type: token.ILLEGAL, Literal: string(l.ch)}
		}

	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	start := l.pos
	for unicode.IsLetter(rune(l.ch)) || unicode.IsDigit(rune(l.ch)) || l.ch == '_' {
		l.readChar()
	}
	return l.input[start:l.pos]
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for unicode.IsDigit(rune(l.ch)) {
		l.readChar()
	}
	return l.input[start:l.pos]
}
