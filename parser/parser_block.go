// parseBlockStatementみたいな関数を復元
package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: parser.currentToken}
	block.Statements = []ast.Statement{}

	parser.nextToken()

	// }かEOFが来るまでparseStatementを呼ぶ
	for !parser.currentTokenIs(token.RBRACE) && !parser.currentTokenIs(token.EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		parser.nextToken()
	}

	return block
}
