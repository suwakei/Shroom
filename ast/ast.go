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

type LetStatement struct {
	Token token.Token // token.LET トークン
	Name *Identifier
	Value Expression
}


func (lstmt *LetStatement) statementNode() {}

func (lstmt *LetStatement) TokenLiteral() string{
	return lstmt.Token.Literal
}


type ReturnStatement struct {
	Token token.Token // returnトークン
	ReturnValue Expression
}


func (rstmt *ReturnStatement) statementNode() {}

func (rstmt *ReturnStatement) TokenLiteral() string {
	return rstmt.Token.Literal
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


func (lstmt *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(lstmt.TokenLiteral() + " ")
	out.WriteString(lstmt.Name.String())
	out.WriteString(" = ")

	if lstmt.Value != nil {
		out.WriteString(lstmt.Value.String())
	}

	out.WriteString(";")

	return out.String()
}


func (rstmt *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rstmt.TokenLiteral() + " ")

	if rstmt.ReturnValue != nil {
		out.WriteString(rstmt.ReturnValue.String())
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
	Token token.TokenType
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {return il.Token.Literal}
func (il *IntegerLiteral) String() string {return il.Token.Literal}