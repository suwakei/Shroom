package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if parser.peekTokenIs(token.RPAREN) {
		parser.nextToken()
		return identifiers
	}

	parser.nextToken()

	ident := &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
	identifiers = append(identifiers, ident)

	for parser.peekTokenIs(token.COMMA) {
		parser.nextToken()
		parser.nextToken()
		ident := &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !parser.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

func (parser *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: parser.currentToken}

	if !parser.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Parameters = parser.parseFunctionParameters()

	if !parser.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = parser.parseBlockStatement()

	return lit
}
