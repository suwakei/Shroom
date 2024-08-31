package ast

import (
	"Shroom/token"
)


type Boolean struct {
	Token token.Token
	Value bool
}


func (boolean *Boolean) expressionNode() {}
func (boolean *Boolean) TokenLiteral() string {return boolean.Token.Literal}
func (boolean *Boolean) String() string {return boolean.Token.Literal}