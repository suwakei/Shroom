package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: parser.currentToken, Left: left}

	parser.nextToken()
	exp.Index = parser.parseExpression(LOWEST)

	if !parser.expectPeek(token.RBRACKET) {
		return nil
	}

	return exp
}
