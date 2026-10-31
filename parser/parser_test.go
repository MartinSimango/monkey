package parser

import (
	"testing"

	"github.com/MartinSimango/monkey/ast"
	"github.com/MartinSimango/monkey/lexer"
)

func TestParser_LetStatements(t *testing.T) {

	input := `
		let x  5;
		let  10;
		let  = 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if ln := len(p.Errors()); ln > 0 {
		t.Errorf("parser encountered  %d errors", ln)
		for _, msg := range p.Errors() {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
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

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("statment not let *ast.LetStatement got=%T", stmt)
		return false
	}

	if letStmt.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiteral() not '%s', got=%s", name, letStmt.TokenLiteral())
		return false
	}
	if letStmt.Name == nil {
		t.Errorf("statement.Name is nil ")

		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("statement.Name.Value not '%s' got='%s'", name, letStmt.Name.Value)
		return false

	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("statement.Name.TokenLiteral() not '%s' got='%s'", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
