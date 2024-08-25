package parser

import (
	"Shroom/ast"
	"Shroom/lexer"
	"Shroom/token"
	"fmt"
	"strconv"
)



type Parser struct {
	lex *lexer.Lexer // 字句解析機インスタンスへのポインタ
	currentToken token.Token // 現在のトークン
	peekToken token.Token //次のトークン
	errors []string

	prefixParseFns map[token.TokenType]prefixParseFunc
	infixParseFns map[token.TokenType]infixParseFunc
}


func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
}


func New (lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex: lex,
		errors: []string{},
	}

	parser.prefixParseFns = make(map[token.TokenType]prefixParseFunc)
	parser.registerPrefix(token.IDENTIFIER, parser.parseIdentifier)
	parser.registerPrefix(token.INT, parser.parseIntegerLiteral)

	parser.registerPrefix(token.BANG, parser.parsePrefixExpression)
	parser.registerPrefix(token.MINUS, parser.parsePrefixExpression)
	// 2つトークンを読み込んでcurrentTokenとpeekTokenの2つがセットされる
	parser.nextToken()
	parser.nextToken()
	return parser
}


func (parser *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token: parser.currentToken,
		Operator: parser.currentToken.Literal,
	}

	parser.nextToken()

	expression.Right = parser.parseExpression(PREFIX)

	return expression
}


func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
t, parser.peekToken.Type)
parser.errors = append(parser.errors, msg)
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lex.NextToken()
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// :FIXME セミコロンに到達するまで式を読み飛ばしている
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}

func (parser *Parser) currentTokenIs(t token.TokenType) bool {
	return parser.currentToken.Type == t
}

func (parser *Parser) peekTokenIs(t token.TokenType) bool {
	return parser.peekToken.Type == t
}

func (parser *Parser) expectPeek(t token.TokenType) bool {
	if parser.peekTokenIs(t) {
		parser.nextToken()
		return true
	}else {
		parser.peekError(t)
		return false
	}
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: parser.currentToken}

	parser.nextToken()

	// FIXME: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}

// 各識別子の優先順位
// iotaの部分が0であとから続く定数には1~7の数字が割り当てられている
const (
	_ int = iota
	LOWEST
	EQUALS // ==
	LESSGREATER // > or <
	SUM // +
	PRODUCT // *
	PREFIX // -x or !x
	CALL // myfunction()
)


func (parser *Parser) noPrefixParseFnError(tok token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", tok)
	parser.errors = append(parser.errors, msg)
}


func (parser *Parser) parseExpression(precedence int) ast.Expression {
	// parser.currentToken.Typeの前置に関連付けられた構文解析関数があるか確認
	// あれば構文解析関数の結果を返す なければnil
	prefix := parser.prefixParseFns[parser.currentToken.Type]
	if prefix == nil {
		parser.noPrefixParseFnError(parser.currentToken.Type)
		return nil
	}
	leftExp := prefix()

	return leftExp
}




func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: parser.currentToken}

	stmt.Expression = parser.parseExpression(LOWEST)

	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}


func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	case token.RETURN:
		return parser.parseReturnStatement()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.currentToken.Type != token.EOF {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.nextToken()
	}
	return program
}


type (
	prefixParseFunc func() ast.Expression // 前置構文解析関数 (++iなど)
	infixParseFunc func(ast.Expression) ast.Expression // 中置構文解析関数 (a + b) + c の()にあたるところ
)


// prefixParseFuncマップにエントリを追加する
func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFunc) {
	parser.prefixParseFns[tokenType] = fn
}

// infixParseFuncにエントリを追加する
func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFunc) {
	parser.infixParseFns[tokenType] = fn
}


func (parser *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: parser.currentToken}

	value, err := strconv.ParseInt(parser.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", parser.currentToken.Literal)
		parser.errors = append(parser.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}