package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: parser.currentToken}

	array.Elements = parser.parseExpressionList(token.RBRACKET)

	return array
}

func (parser *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if parser.peekTokenIs(end) {
		parser.nextToken()
		return list
	}

	parser.nextToken()
	list = append(list, parser.parseExpression(LOWEST))

	for parser.peekTokenIs(token.COMMA) {
		parser.nextToken()
		parser.nextToken()
		list = append(list, parser.parseExpression(LOWEST))
	}

	if !parser.expectPeek(end) {
		return nil
	}

	return list
}
