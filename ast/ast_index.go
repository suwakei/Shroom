package ast

import (
	"Shroom/token"
	"bytes"
)

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (idxe *IndexExpression) expressionNode()      {}
func (idxe *IndexExpression) TokenLiteral() string { return idxe.Token.Literal }
func (idxe *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(idxe.Left.String())
	out.WriteString("[")
	out.WriteString(idxe.Index.String())
	out.WriteString("]")

	return out.String()
}
