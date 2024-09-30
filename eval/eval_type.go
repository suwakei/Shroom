package eval

import (
	"Shroom/object"
	"reflect"
)

func getType(value interface{}) *object.ObjType {
	var target reflect.Type = reflect.TypeOf(value)

	if target.String() == "*object.Integer" {
		return &object.ObjType{Value: "int"}
	}

	if target.String() == "*object.String" {
		return &object.ObjType{Value: "string"}
	}

	if target.String() == "*object.Boolean" {
		return &object.ObjType{Value: "boolean"}
	}

	if target.String() == "*object.Array" {
		return &object.ObjType{Value: "array"}
	}

	if target.String() == "*object.Function" {
		return &object.ObjType{Value: "function"}
	}

	if target.String() == "*object.Null" {
		return &object.ObjType{Value: "null"}
	}

	return &object.ObjType{Value: target.String()}
}
