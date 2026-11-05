package parser

import (
	"fmt"
	"strconv"

	"github.com/MartinSimango/monkey/ast"
	"github.com/MartinSimango/monkey/lexer"
	"github.com/MartinSimango/monkey/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
	// postfixParseFn - TODO: something for you to figure out
)

const (
	_ int = iota
	LOWEST
	EQUALS      //
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x
	CALL        // myFunction(x)

)

type Parser struct {
	l         *lexer.Lexer
	errors    []string
	currToken token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	// p.registerPrefix(token.BANG, p.)
	return p
}

func (p *Parser) registerPrefix(t token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

func (p *Parser) registerInfix(t token.TokenType, fn infixParseFn) {
	p.infixParseFns[t] = fn
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
	case token.RETURN:
		stmt := p.parseReturnStatement()
		if stmt == nil {
			return nil
		}
		return stmt
	default:
		stmt := p.parseExpressionStatement()
		if stmt == nil {
			return nil
		}
		return stmt
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}
func (p *Parser) parseIntegerLiteral() ast.Expression {
	val, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("failed to convert integer literal value of %s to int64 ", p.currToken.Literal))
		return nil
	}
	return &ast.IntegerLiteral{Token: p.currToken, Value: val}
}
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	retStmt := &ast.ReturnStatement{Token: p.currToken}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return retStmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	letStmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
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
