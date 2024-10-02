package eval

import (
	"Shroom/object"
	"fmt"
)

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)

	case left.Type() == object.DICT_OBJ:
		return evalDictIndexExpression(left, index)

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

func evalDictIndexExpression(dict, index object.Object) object.Object {
	dictObject := dict.(*object.Dict)

	key, ok := index.(object.HashableDict)
	if !ok {
		return newError("unusable as dict key: %s", index.Type())
	}

	pair, ok := dictObject.Pairs[key.DictKey()]
	if !ok {
		return NULL
	}

	return pair.Value
}
