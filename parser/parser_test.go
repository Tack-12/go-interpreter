package parser

import (
	"go-interpreter/ast"
	"go-interpreter/lexer"
	"testing"
)

/*
func TestLetStatements(t *testing.T) {
	input := `
	let x = 5 ;
	let y = 10 ;
	let foobar = 838383 ;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s not *ast.Statement. got=%T", s)
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s is not *ast.Statement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("name is not same. got=%s , wanted=%s", letStmt.Name, name)
		return false
	}

	return true
}

*/

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors ", len(errors))
	for _, msg := range errors {
		t.Errorf("parser Errors: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
	return  5;
	return  10 ;
	return  838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("The return statement is not 'return' got =%q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {

	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program doesn't have enough statements. got=%T", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statement[0] is not astStatement. got=%T ", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("program.Statement[0] is not a Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value is not %s . got =%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.Value is not %s . got =%s", "foobar", ident.TokenLiteral())
	}

}

func testIntegerLiteral(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program doesn't have enough statements. got=%T", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statement[0] is not astStatement. got=%T ", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral) //Assert that the first expression we get is an Integer Literal Expression

	if !ok {
		t.Fatalf("program.Statement[0] is not a literalifier. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value is not %d . got =%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.Value is not %s . got =%s", "5", literal.TokenLiteral())
	}
}
