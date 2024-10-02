package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseDictLiteral() ast.Expression {
	dict := &ast.DictLiteral{Token: parser.currentToken}
	dict.Pairs = make(map[ast.Expression]ast.Expression)

	for !parser.peekTokenIs(token.RBRACE) {
		parser.nextToken()
		key := parser.parseExpression(LOWEST)

		if !parser.expectPeek(token.COLON) {
			return nil
		}

		parser.nextToken()
		value := parser.parseExpression(LOWEST)

		dict.Pairs[key] = value

		if !parser.peekTokenIs(token.RBRACE) && !parser.expectPeek(token.COMMA) {
			return nil
		}
	}

	if !parser.expectPeek(token.RBRACE) {
		return nil
	}

	return dict
}
