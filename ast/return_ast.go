package ast

import (
	"Shroom/token"
	"bytes"
)


type ReturnStatement struct {
	Token token.Token // returnトークン
	ReturnValue Expression
}


func (rstmt *ReturnStatement) statementNode() {}

func (rstmt *ReturnStatement) TokenLiteral() string {
	return rstmt.Token.Literal
}


func (rstmt *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rstmt.TokenLiteral() + " ")

	if rstmt.ReturnValue != nil {
		out.WriteString(rstmt.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}


