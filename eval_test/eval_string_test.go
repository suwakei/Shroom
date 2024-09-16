package eval_test

import (
	"Shroom/object"
	"testing"
)


func TestStringLiteral(t *testing.T) {
	input := "Hello World!"

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T(%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=$q", str.Value)
	}
}