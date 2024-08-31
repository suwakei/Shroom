package ast

import (
	"Shroom/token"
)

type Identifier struct {
	Token token.Token // token.IDENTIFIER トークン
	Value string
}

func (identifier *Identifier) expressionNode() {}

func (identifier *Identifier) TokenLiteral() string {
	return identifier.Token.Literal
}

func (iden *Identifier) String() string {
	return iden.Value
}
