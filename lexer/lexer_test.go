package lexer

import (
	"testing"
	"Shroom/token"
)


func TestNextToken(t *testing.T) {
	input := "=+-(){},;:"

	tests := []struct {
		expectedType token.Tokentype
		expectedLitetal string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COLON, ":"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ";"},
		{token.EOF, "EOF"},
	}

	l := New(input)

	for i, tt := range tests{
		tok := l.TestNextToken

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - literal wrong expected=%q, got=%q",
		i, tt.expectedLiteral, tok.Literal)
		}
	}
}