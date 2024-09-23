package parser

import (
	"Shroom/ast"
)

func (parser *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: parser.currentToken, Value: parser.currentToken.Literal}
}
