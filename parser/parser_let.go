package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	parser.nextToken()

	stmt.Value = parser.parseExpression(LOWEST)

	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	// // :FIXME セミコロンに到達するまで式を読み飛ばしている
	// for !parser.currentTokenIs(token.SEMICOLON) {
	// 	parser.nextToken()
	// }

	return stmt
}
