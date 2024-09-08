package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}


// object.Integer型
type Integer struct {
	Value int64
}

func (integer *Integer) Inspect() string {return fmt.Sprintf("%d", integer.Value)}
func (integer *Integer) Type() ObjectType {return INTEGER_OBJ}


// object.Boolean型
type Boolean struct {
	Value bool
}

func (boolean *Boolean) Type() ObjectType {return BOOLEAN_OBJ}
func (boolean *Boolean) Inspect() string {return fmt.Sprintf("%t", boolean.Value)}


// null型
type Null struct {}

func (n *Null) Type() ObjectType {return NULL_OBJ}
func (n *Null) Inspect() string {return "null"}