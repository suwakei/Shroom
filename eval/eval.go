package eval

import (
	"Shroom/ast"
	"Shroom/object"
	"fmt"
)

// バリエーションが少ない語に遭遇した時ここから参照する
var(
	NULL = &object.Null{}
	TRUE = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)


func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	
	case *ast.Boolean:
		return nativeBooltoBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatement:
		return evalBlockStatement(node)

	case *ast.IfExpression:
		return evalIfExpression(node)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		return &object.ReturnValue{Value: val}
	}

	return nil
}


func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}

		if result != nil && result.Type() == object.RETURN_VALUE_OBJ {
			return result
		}
	}

	return result
}


func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement)

		switch result := result.(type){
		case *object.ReturnValue:
			return result.Value
		
		case *object.Error:
			return result
		}
	

	if returnValue, ok := result.(*object.ReturnValue); ok {
		return returnValue.Value
	}
}
return result
}


func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)

	if isTruThy(condition) {
		return Eval(ie.Consequence)
	}else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	}else {
		return NULL
	}
}

func isTruThy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false

	case TRUE:
		return true

	case FALSE:
		return false

	default:
		return true
	}
}

func nativeBooltoBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

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


func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	
	case FALSE:
		return TRUE

	case NULL:
		return TRUE

	default:
		return FALSE
	}
}


func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}


func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}


func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}


