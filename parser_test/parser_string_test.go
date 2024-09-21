package parser_test

import (
	"Shroom/ast"
	"Shroom/lexer"
	"Shroom/parser"
	"testing"
)

func TestStringLiteralExpression(t *testing.T) {
	input := `Hello World;`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != "Hello World" {
		t.Fatalf("literal.Value is not %q. got=%q", "Hello World", literal.Value)
	}
}
