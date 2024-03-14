package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte // current character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // 调用一次，初始化position: 0, readPosition: 1
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.skipWhiteSpace() // white space
	switch string(l.ch) {

	case token.ASSIGN:
		if l.peekChar() == '=' {
			t = token.Token{Type: token.EQ, Literal: "=="}
			l.readChar()
		} else {
			t = NewToken(token.ASSIGN, l.ch)
		}

	case token.PLUS:
		t = NewToken(token.PLUS, l.ch)

	case token.MINUS:
		t = NewToken(token.MINUS, l.ch)

	case token.ASTERISK:
		t = NewToken(token.ASTERISK, l.ch)

	case token.BANG:
		if l.peekChar() == '=' {
			t = token.Token{Type: token.NOT_EQ, Literal: "!="}
			l.readChar()
		} else {
			t = NewToken(token.BANG, l.ch)
		}

	case token.SLASH:
		t = NewToken(token.SLASH, l.ch)

	case token.LT:
		t = NewToken(token.LT, l.ch)

	case token.GT:
		t = NewToken(token.GT, l.ch)

	case token.COMMA:
		t = NewToken(token.COMMA, l.ch)

	case token.SEMICOLON:
		t = NewToken(token.SEMICOLON, l.ch)

	case token.LPAREN:
		t = NewToken(token.LPAREN, l.ch)

	case token.RPAREN:
		t = NewToken(token.RPAREN, l.ch)

	case token.LBRACE:
		t = NewToken(token.LBRACE, l.ch)

	case token.RBRACE:
		t = NewToken(token.RBRACE, l.ch)

	case string(byte(0)):
		t = token.Token{Type: token.EOF, Literal: ""}

	default:
		if isLetter(l.ch) {
			str := l.readIdentifier()
			tokenType := token.LookupIdent(str)
			return token.Token{Type: tokenType, Literal: str} // position 已经指向下一个字符
		} else if isNumber(l.ch) {
			str := l.readInt()
			return token.Token{Type: token.INT, Literal: str}
		} else {
			t = NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readInt() string {
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for isWhiteSpace(l.ch) {
		l.readChar()
	}
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch == '_'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
