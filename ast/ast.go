package ast
// 構文解析機

import (
	"Shroom/token"
)

type Node interface {
	TokenLiteral() string
}

// 文を表す
type Statement interface {
	Node
	statementNode()
}

// 式を表す
type Expression interface {
	Node
	expressionNode()
}


// 全ASTのルートノード
type Program struct {
	Statements []Statement
}


func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET トークン
	Name *Identifier
	Value Expression
}

func (lstm *LetStatement) statementNode() {}

func (lstm *LetStatement) TokenLiteral() string{
	return lstm.Token.Literal
}

type Identifier struct{
	Token token.Token // token.IDENTIFIER トークン
	value string
}

func (iden *Identifier) expressionNode() {}

func (iden *Identifier) TokenLiteral() string{
	return iden.Token.Literal
}
