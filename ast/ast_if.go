package ast

import (
	"Shroom/token"
	"bytes"
)

// if (<expression>) <consequence> else <alternative>
type IfExpression struct {
	Token token.Token // ifトークン
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ife *IfExpression) expressionNode() {}
func (ife *IfExpression) TokenLiteral() string {return ife.Token.Literal}
func (ife *IfExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("if")
	out.WriteString(ife.Condition.String())
	out.WriteString(" ")
	out.WriteString(ife.Consequence.String())

	if ife.Alternative != nil {
		out.WriteString("else")
		out.WriteString(ife.Alternative.String())
	}

	return out.String()
}



