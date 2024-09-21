package eval

import (
	"Shroom/object"
	"reflect"
)



func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)

	case "-":
		return evalMinusPrefixOperatorExpression(right)

	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}


func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}


func evalInfixExpression(operator string,
	left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)

	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())

	case operator == "==":
		return nativeBooltoBooleanObject(left == right)

	case operator == "!=":
		return nativeBooltoBooleanObject(left != right)

	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)

	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	// if operator != "+" {
	// 	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	// }

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}

	case ">":
		return nativeBooltoBooleanObject(len(leftVal) > len(rightVal))

	case "<":
		return nativeBooltoBooleanObject(len(leftVal) < len(rightVal))

		//上手く機能していないので直す
	case "==":
		l := object.String{Value: leftVal}
		r := object.String{Value: rightVal}
		return nativeBooltoBooleanObject(reflect.DeepEqual(l, r))

	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())

	}

}

func evalIntegerInfixExpression(operator string,
	left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}

	case "-":
		return &object.Integer{Value: leftVal - rightVal}

	case "*":
		return &object.Integer{Value: leftVal * rightVal}

	case "/":
		return &object.Integer{Value: leftVal / rightVal}

	case "<":
		return nativeBooltoBooleanObject(leftVal < rightVal)

	case ">":
		return nativeBooltoBooleanObject(leftVal > rightVal)

	case "==":
		return nativeBooltoBooleanObject(leftVal == rightVal)

	case "!=":
		return nativeBooltoBooleanObject(leftVal != rightVal)

	default:
		return newError("unknown opeartor: %s %s %s", left.Type(), operator, right.Type())
	}
}