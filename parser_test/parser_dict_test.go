package parser_test

import (
	"Shroom/ast"
	"Shroom/lexer"
	"Shroom/parser"
	"testing"
)

func TestParsingDictLiteralStringKeys(t *testing.T) {
	input := `{"one": 1, "two": 2, "three": 3}`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	dict, ok := stmt.Expression.(*ast.DictLiteral)
	if !ok {
		t.Fatalf("exp is not ast.DictLiteral. got=%T", stmt.Expression)
	}

	if len(dict.Pairs) != 3 {
		t.Errorf("dict.Pairs has wrong length. got=%d", len(dict.Pairs))
	}

	expected := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, value := range dict.Pairs {
		literal, ok := key.(*ast.StringLiteral)
		if !ok {
			t.Errorf("key is not ast.StringLitetal. got=%T", key)
		}

		expectedValue := expected[literal.String()]

		testIntegerLiteral(t, value, expectedValue)
	}
}

func TestParsingEmptyDictLiteral(t *testing.T) {
	input := "{}"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	dict, ok := stmt.Expression.(*ast.DictLiteral)
	if !ok {
		t.Fatalf("exp is not ast.DictLiteral. got=%T", stmt.Expression)
	}

	if len(dict.Pairs) != 0 {
		t.Errorf("dict.Pairs has wrong length. got=%d", len(dict.Pairs))
	}
}

func TestParsingDictLiteralsWithExpressions(t *testing.T) {
	input := `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	dict, ok := stmt.Expression.(*ast.DictLiteral)
	if !ok {
		t.Fatalf("exp is not ast.DictLiteral. got=%T", stmt.Expression)
	}

	if len(dict.Pairs) != 3 {
		t.Errorf("dict.Pairs has wrong length. got=%d", len(dict.Pairs))
	}

	tests := map[string]func(ast.Expression){
		"one": func(e ast.Expression) {
			testInfixExpression(t, e, 0, "+", 1)
		},
		"two": func(e ast.Expression) {
			testInfixExpression(t, e, 10, "-", 8)
		},
		"three": func(e ast.Expression) {
			testInfixExpression(t, e, 15, "/", 5)
		},
	}

	for key, value := range dict.Pairs {
		literal, ok := key.(*ast.StringLiteral)
		if !ok {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)
			continue
		}

		testFunc, ok := tests[literal.String()]
		if !ok {
			t.Errorf("No test function for key %q found", literal.String())
			continue
		}

		testFunc(value)
	}
}
