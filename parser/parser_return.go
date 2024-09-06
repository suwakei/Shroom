package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: parser.currentToken}

	parser.nextToken()

	stmt.ReturnValue = parser.parseExpression(LOWEST)

	// // FIXME: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	// for !parser.currentTokenIs(token.SEMICOLON) {
	// 	parser.nextToken()
	// }

	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}
