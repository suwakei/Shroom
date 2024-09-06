package parser

import (
	"Shroom/ast"
	"Shroom/lexer"
	"Shroom/token"
	"fmt"
)

type Parser struct {
	lex          *lexer.Lexer // 字句解析機インスタンスへのポインタ
	currentToken token.Token  // 現在のトークン
	peekToken    token.Token  //次のトークン
	errors       []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression               // 前置構文解析関数 (++iなど)
	infixParseFn  func(ast.Expression) ast.Expression // 中置構文解析関数 (a + b) + c の()にあたるところ
)

func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex:    lex,
		errors: []string{},
	}

	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.registerPrefix(token.IDENTIFIER, parser.parseIdentifier)
	parser.registerPrefix(token.INT, parser.parseIntegerLiteral)

	parser.registerPrefix(token.BANG, parser.parsePrefixExpression)
	parser.registerPrefix(token.MINUS, parser.parsePrefixExpression)

	parser.registerPrefix(token.TRUE, parser.parseBoolean)
	parser.registerPrefix(token.FALSE, parser.parseBoolean)

	parser.registerPrefix(token.LPAREN, parser.parseGroupExpression)
	parser.registerPrefix(token.IF, parser.parseIfExpression)

	parser.registerPrefix(token.FUNCTION, parser.parseFunctionLiteral)

	// 中置構文解析関数を中置演算子に登録する
	// 中置演算子はparser.parseInfixExpressionに関連付けられる
	parser.infixParseFns = make(map[token.TokenType]infixParseFn)
	parser.registerInfix(token.PLUS, parser.parseInfixExpression)
	parser.registerInfix(token.MINUS, parser.parseInfixExpression)

	parser.registerInfix(token.SLASH, parser.parseInfixExpression)
	parser.registerInfix(token.ASTARISK, parser.parseInfixExpression)

	parser.registerInfix(token.EQUAL, parser.parseInfixExpression)
	parser.registerInfix(token.NOT_EQUAL, parser.parseInfixExpression)

	parser.registerInfix(token.LT, parser.parseInfixExpression)
	parser.registerInfix(token.GT, parser.parseInfixExpression)

	parser.registerInfix(token.LPAREN, parser.parseCallExpression)

	// 2つトークンを読み込んでcurrentTokenとpeekTokenの2つがセットされる
	parser.nextToken()
	parser.nextToken()
	return parser
}

func (parser *Parser) parseGroupExpression() ast.Expression {
	parser.nextToken()

	exp := parser.parseExpression(LOWEST)

	if !parser.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
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
	} else {
		parser.peekError(t)
		return false
	}
}

// 各識別子の優先順位
// iotaの部分が0であとから続く定数には1~7の数字が割り当てられている
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x or !x
	CALL        // myfunction()
)

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	// parser.currentToken.Typeの前置に関連付けられた構文解析関数があるか確認
	// あれば構文解析関数の結果を返す なければnil
	prefix := parser.prefixParseFns[parser.currentToken.Type]
	if prefix == nil {
		parser.noPrefixParseFnError(parser.currentToken.Type)
		return nil
	}
	leftExp := prefix()

	for !parser.peekTokenIs(token.SEMICOLON) && precedence < parser.peekPrecedence() {
		infix := parser.infixParseFns[parser.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		parser.nextToken()

		leftExp = infix(leftExp)
	}

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

// これが優先順位テーブルとなる
var precedences = map[token.TokenType]int{
	token.EQUAL:     EQUALS,
	token.NOT_EQUAL: EQUALS,
	token.LT:        LESSGREATER,
	token.GT:        LESSGREATER,
	token.PLUS:      SUM,
	token.MINUS:     SUM,
	token.SLASH:     PRODUCT,
	token.ASTARISK:  PRODUCT,
	token.LPAREN:    CALL,
}

func (parser *Parser) peekPrecedence() int {
	if parser, ok := precedences[parser.peekToken.Type]; ok {
		return parser
	}

	return LOWEST
}

func (parser *Parser) currentPrecedence() int {
	if parser, ok := precedences[parser.currentToken.Type]; ok {
		return parser
	}

	return LOWEST
}

func (parser *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: parser.currentToken, Value: parser.currentTokenIs(token.TRUE)}
}
