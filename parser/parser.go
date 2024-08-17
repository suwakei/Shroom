package parser

import (
	"Shroom/ast"
	"Shroom/lexer"
	"Shroom/token"
	"fmt"
)

type Parser struct {
	lex *lexer.Lexer // 字句解析機インスタンスへのポインタ
	currentToken token.Token // 現在のトークン
	peekToken token.Token //次のトークン
	errors []string
}

func New (lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex: lex,
		errors: []string{},
	}
	// 2つトークンを読み込んでcurrentTokenとpeekTokenの2つがセットされる
	parser.nextToken()
	parser.nextToken()
	return parser
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
t, parser.peekToken.Type)
parser.errors = append(parser.errors, msg)
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lex.NextToken()
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// :FIXME セミコロンに到達するまで式を読み飛ばしている
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}

func (parser *Parser) currentTokenIs(t token.TokenType) bool {
	return parser.currentToken.Type == t
}

func (parser *Parser) peekTokenIs(t token.TokenType) bool {
	return parser.peekToken.Type == t
}

func (parser *Parser) expectPeek(t token.TokenType) bool {
	if parser.peekTokenIs(t) {
		parser.nextToken()
		return true
	}else {
		parser.peekError(t)
		return false
	}
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.currentToken.Type != token.EOF {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.nextToken()
	}
	return program
}