package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: parser.currentToken}

	parser.nextToken()

	// FIXME: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}
