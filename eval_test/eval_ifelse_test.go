package eval_test

import (
	"testing"
)

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) {10}", 10},
		{"if (false) {10}", nil},
		{"if (1) {100}", 100},
		{"if (1 < 2) {19}", 19},
		{"if (1 > 2) {10}", nil},
		{"if (1 > 2) {10} else {30}", 30},
		{"if (1 < 2) {90} else {20}", 90},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
