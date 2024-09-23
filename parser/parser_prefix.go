package parser

import (
	"Shroom/ast"
	"Shroom/token"
	"fmt"
)

func (parser *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
	}

	parser.nextToken()

	expression.Right = parser.parseExpression(PREFIX)

	return expression
}

func (parser *Parser) noPrefixParseFnError(tok token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", tok)
	parser.errors = append(parser.errors, msg)
}

// prefixParseFnマップにエントリを追加する
func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn
}
