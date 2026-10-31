package parser

import (
	"fmt"

	"github.com/MartinSimango/monkey/ast"
	"github.com/MartinSimango/monkey/lexer"
	"github.com/MartinSimango/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	errors    []string
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: make([]ast.Statement, 0)}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		// This may seem unnceesray but if parseLetStatement returns nil
		// we can't return stmt as statement would be type *ast.LetStatement
		// so me need to
		stmt := p.parseLetStatement()
		if stmt == nil {
			return nil
		}
		return stmt
	}
	return nil
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// var letStmt *ast.LetStatement
	letStmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		// p.errors
		return nil
	}
	letStmt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
	if !p.expectPeek(token.ASSIGN) {

		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return letStmt
}

func (p *Parser) expectPeek(tokType token.TokenType) bool {
	if p.peekTokenIs(tokType) {
		p.nextToken()
		return true
	}
	p.peekError(tokType)
	return false
}

func (p *Parser) peekError(tokType token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("expected next token to be '%s' got '%s' instead", tokType, p.peekToken.Type))
}

func (p *Parser) curTokenIs(tokType token.TokenType) bool {
	return p.currToken.Type == tokType
}

func (p *Parser) peekTokenIs(tokType token.TokenType) bool {
	return p.peekToken.Type == tokType
}
