package eval

import (
	"Shroom/ast"
	"Shroom/object"
)

func evalDictLiteral(node *ast.DictLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.DictKey]object.DictPair)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		dictKey, ok := key.(object.HashableDict)
		if !ok {
			return newError("unusable as dict key: %s", key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := dictKey.DictKey()
		pairs[hashed] = object.DictPair{Key: key, Value: value}
	}

	return &object.Dict{Pairs: pairs}
}
