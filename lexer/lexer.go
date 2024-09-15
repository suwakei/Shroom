package lexer

/*
lexer 字句解析機
*/

import (
	"Shroom/token"
)

type Lexer struct {
	input        string //入力
	position     int    // 入力における現在位置(現在の位置を示す)
	readPosition int    // これから読み込む位置(現在の次の文字)
	ch           byte   // 現在検査中の文字
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 数字かどうか判定
func isDisit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lex *Lexer) readNumber() string {
	position := lex.position
	for isDisit(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

// 空白、改行、タブを読み飛ばす
func (lex *Lexer) skipWhiteSpace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

// 英字かどうか識別
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '-'
}

func (lex *Lexer) readIdentifer() string {
	position := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

// 二文字以上の予約語かどうか判定
func (lex *Lexer) peekChar() byte {
	if lex.readPosition >= len(lex.input) {
		return 0
	} else {
		return lex.input[lex.readPosition]
	}
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhiteSpace()

	switch lex.ch {
	case '=':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			literal := string(ch) + string(lex.ch)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, lex.ch)
		}
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '-':
		tok = newToken(token.MINUS, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case ':':
		tok = newToken(token.COLON, lex.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	case '!':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			literal := string(ch) + string(lex.ch)
			tok = token.Token{Type: token.NOT_EQUAL, Literal: literal}
		} else {
			tok = newToken(token.BANG, lex.ch)
		}
	case '*':
		tok = newToken(token.ASTARISK, lex.ch)
	case '/':
		tok = newToken(token.SLASH, lex.ch)
	case '<':
		tok = newToken(token.LT, lex.ch)
	case '>':
		tok = newToken(token.GT, lex.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = lex.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		// 英字かどうか判定
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifer()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDisit(lex.ch) {
			tok.Type = token.INT
			tok.Literal = lex.readNumber()
			return tok
		} else {
			tok = newToken(token.UNDEF, lex.ch)
		}
	}

	lex.readChar()
	return tok
}


func (lex *Lexer) readString() string {
	position := lex.position + 1
	for {
		lex.readChar()
		if lex.ch == '"' || lex.ch == 0 {
			break
		}
	}
	return lex.input[position:lex.position]
}