package eval_test

import (
	"Shroom/eval"
	"Shroom/lexer"
	"Shroom/object"
	"Shroom/parser"
	"testing"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return eval.Eval(program, env)
}

// Bang == !
func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != eval.NULL {
		t.Errorf("object is not Null. got=%T(%+v)", obj, obj)
		return false
	}
	return true
}
