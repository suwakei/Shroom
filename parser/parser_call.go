package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: parser.currentToken, Function: function}
	exp.Arguments = parser.parseCallArguments()
	return exp
}

func (parser *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if parser.peekTokenIs(token.RPAREN) {
		parser.nextToken()
		return args
	}

	parser.nextToken()
	args = append(args, parser.parseExpression(LOWEST))

	for parser.peekTokenIs(token.COMMA) {
		parser.nextToken()
		parser.nextToken()
		args = append(args, parser.parseExpression(LOWEST))
	}

	if !parser.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}