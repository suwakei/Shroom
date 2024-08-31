package ast

// 構文解析機

import (
	"Shroom/token"
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression
}

func (estmt *ExpressionStatement) statementNode() {}

func (estmt *ExpressionStatement) TokenLiteral() string {
	return estmt.Token.Literal
}

func (program *Program) String() string {
	var out bytes.Buffer

	for _, s := range program.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (estmt *ExpressionStatement) String() string {
	if estmt.Expression != nil {
		return estmt.Expression.String()
	}

	return ""
}

// ブロック構造を表現する
type BlockStatement struct {
	Token token.Token //トークン
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string {return bs.Token.Literal}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}