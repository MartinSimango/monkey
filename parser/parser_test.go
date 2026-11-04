package parser

import (
	"testing"

	"github.com/MartinSimango/monkey/ast"
	"github.com/MartinSimango/monkey/lexer"
)

func TestParser_ParseProgramLetStatements(t *testing.T) {
	tests := []struct {
		name                string
		input               string
		statementCount      int
		expectedIdentifiers []string
	}{
		{
			name: "Test let statements",
			input: `
		let x = 5;
		let y = 10;
		let foobar = 838383;`,
			statementCount:      3,
			expectedIdentifiers: []string{"x", "y", "foobar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
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
			if ln := len(program.Statements); ln != tt.statementCount {
				t.Fatalf("program.Statements does not contain %d statments got=%d", tt.statementCount, ln)
			}

			for i, stmt := range program.Statements {
				if !testLetStatement(t, stmt, tt.expectedIdentifiers[i]) {
					return

				}
			}

		})
	}

}

func TestParser_ParseProgramReturnStatements(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		statementCount int
	}{
		{
			name: "Test return statements",
			input: `
			return 5;
			return 10;
			return fooar;`,
			statementCount: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
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
			if ln := len(program.Statements); ln != tt.statementCount {
				t.Fatalf("program.Statements does not contain %d statments got=%d", tt.statementCount, ln)
			}

			for _, stmt := range program.Statements {
				if _, ok := stmt.(*ast.ReturnStatement); !ok {
					t.Fatalf("expected *ast.ReturnStatement got=%T", stmt)
				}
			}

		})
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

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	p := New(lexer.New(input))

	program := p.ParseProgram()

	if ln := len(program.Statements); ln != 1 {
		t.Fatalf("program.Statements does not contain %d statments got=%d", 1, ln)
	}

	if len(p.Errors()) > 0 {
		t.Errorf("program encounted %d errors", len(p.Errors()))
		for _, msg := range p.Errors() {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
	}

	s, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expected program.Statement[0] to be ast.ExpressionStatement and got %T", program.Statements[0])
	}

	ident, ok := s.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expected ident to be ast.Identifier and got %T", s.Expression)

	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not foobar got %q", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.Tokenliteral() not foobar got %q", ident.TokenLiteral())

	}

}

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	p := New(lexer.New(input))

	program := p.ParseProgram()

	if ln := len(program.Statements); ln != 1 {
		t.Fatalf("program.Statements does not contain %d statments got=%d", 1, ln)
	}

	if len(p.Errors()) > 0 {
		t.Errorf("program encounted %d errors", len(p.Errors()))
		for _, msg := range p.Errors() {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
	}

	s, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expected program.Statement[0] to be ast.ExpressionStatement and got %T", program.Statements[0])
	}

	ident, ok := s.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expected ident to be ast.IntegerLiteral and got %T", s.Expression)

	}

	if ident.Value != "5" {
		t.Errorf("ident.Value not 5 got %q", ident.Value)
	}

	if ident.TokenLiteral() != "5" {
		t.Errorf("ident.Tokenliteral() not 5 got %q", ident.TokenLiteral())

	}

}
