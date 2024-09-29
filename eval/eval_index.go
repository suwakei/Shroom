package eval

import (
	"Shroom/object"
	"fmt"
)

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)

	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)
	fmt.Println(max)

	if idx < 0 || idx > max {
		return &object.Error{Message: "array index out of range"}
	}

	return arrayObject.Elements[idx]
}
