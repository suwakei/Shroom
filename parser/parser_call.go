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
		parser.peek
	}
}