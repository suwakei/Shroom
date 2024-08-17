package parser

import (
	"Shroom/ast"
	"Shroom/lexer"
	"Shroom/token"
)

type Parser struct {
	lex *lexer.Lexer
	currentToken token.Token
	peekToken token.Token
}

func New (lex *lexer.Lexer) *Parser {
	
}