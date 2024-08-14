package lexer

import (
	"Shroom/token"
)

type Lexer struct {
	input string
	position int // 入力における現在位置(現在の位置を示す)
	readPosition int // これから読み込む位置(現在の次の文字)
	ch byte // 現在検査中の文字 
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	}else {
		lex.ch = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}

func NewToken(tokenType token.Tokentype, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}


func (lex *Lexer) NextToken() token.Token{
	var tok token.Token

	switch lex.ch {
	case '=':
		tok = NewToken(token.ASSIGN, lex.ch)
	case '+':
		tok = NewToken(token.PLUS, lex.ch)
	case '-':
		tok = NewToken(token.MINUS, lex.ch)
	case ',':
		tok = NewToken(token.COMMA, lex.ch)
	case ':':
		tok = NewToken(token.COLON, lex.ch)
	case ';':
		tok = NewToken(token.SEMICOLON, lex.ch)
	case '(':
		tok = NewToken(token.LPAREN, lex.ch)
	case ')':
		tok = NewToken(token.RPAREN, lex.ch)
	case '{':
		tok = NewToken(token.LBRACE, lex.ch)
	case '}':
		tok = NewToken(token.RBRACE, lex.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	lex.readChar()
	return tok
}