package parser

import (
	"Shroom/ast"
	"Shroom/token"
)

func (parser *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
		Left:     left,
	}

	precedence := parser.currentPrecedence()
	parser.nextToken()
	expression.Right = parser.parseExpression(precedence)

	return expression
}

// infixParseFuncにエントリを追加する
func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFunc) {
	parser.infixParseFns[tokenType] = fn
}
