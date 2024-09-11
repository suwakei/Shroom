package object

import (
	"fmt"
	"bytes"
	"Shroom/ast"
	"strings"
)

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	NULL_OBJ = "NULL"
	ERROR_OBJ = "ERROR"
	FUNCTION_OBJ = "FUNCTION"

)

type ObjectType string

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


type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {return RETURN_VALUE_OBJ}
func (rv *ReturnValue) Inspect() string {return rv.Value.Inspect()}


type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {return ERROR_OBJ}
func (e *Error) Inspect() string {return "ERROR: " + e.Message}


type Function struct {
	Parameters []*ast.Identifier
	Body *ast.BlockStatement
	Env *Environment
}

func (f *Function) Type() ObjectType {return FUNCTION_OBJ}
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}