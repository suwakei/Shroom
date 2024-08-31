package ast

import (
	"Shroom/token"
	"bytes"
)

type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
}

func (lstmt *LetStatement) statementNode() {}

func (lstmt *LetStatement) TokenLiteral() string {
	return lstmt.Token.Literal
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
