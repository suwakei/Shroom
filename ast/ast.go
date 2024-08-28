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


func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}else {
		return ""
	}
}


type ExpressionStatement struct {
	Token token.Token // 式の最初のトークン
	Expression Expression
}


func (estmt *ExpressionStatement) statementNode() {}

func (estmt *ExpressionStatement) TokenLiteral() string {
	return estmt.Token.Literal
}

type Identifier struct{
	Token token.Token // token.IDENTIFIER トークン
	Value string
}

func (identifier *Identifier) expressionNode() {}

func (identifier *Identifier) TokenLiteral() string {
	return identifier.Token.Literal
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


func (iden *Identifier) String() string {
	return iden.Value
}


type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {return il.Token.Literal}
func (il *IntegerLiteral) String() string {return il.Token.Literal}



type PrefixExpression struct {
	Token token.Token // 前置トークン 例えば ! や ++
	Operator string
	Right Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string {return pe.Token.Literal}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}


type InfixExpression struct {
	Token token.Token // 演算子トークン例えば +
	Left Expression
	Operator string
	Right Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string {return ie.Token.Literal}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}