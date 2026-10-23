package lexer

import "github.com/MartinSimango/monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input string (points to current char)
	readPosition int  // current reading position in input string (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

var keywords = map[string]token.TokenType{
	"let":    token.LET,
	"fn":     token.FUNCTION,
	"if":     token.IF,
	"return": token.RETURN,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"else":   token.ELSE,
}

func getTokenType(ident string) token.TokenType {
	if keyword, ok := keywords[ident]; ok {
		return keyword
	}
	return token.IDENT
}

// TODO: idea - lexer only supports ASCII for - add suport for UNICODE and UTF-8
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()

	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) skipwhitespace() {
	for l.ch == '\t' || l.ch == '\r' || l.ch == ' ' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) peakChar(ch byte, ogTok, tokType token.TokenType) token.Token {
	if l.readPosition <= len(l.input)+1 {
		if l.input[l.readPosition] == ch {
			current := l.ch
			l.readChar()
			return token.Token{
				Type:    tokType,
				Literal: string(current) + string(ch),
			}
		}
	}
	return newToken(ogTok, l.ch)
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipwhitespace()
	switch l.ch {
	case '=':
		tok = l.peakChar('=', token.ASSIGN, token.EQ)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '!':
		tok = l.peakChar('=', token.BANG, token.NOT_EQ)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERIX, l.ch)
	case '<':
		tok = l.peakChar('=', token.LT, token.LT_EQ)
	case '>':
		tok = l.peakChar('=', token.GT, token.GT_EQ)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tok.Literal = ident
			tok.Type = getTokenType(ident)
			return tok
		} else if isDigit(l.ch) {
			number := l.readNumber()
			tok.Literal = number
			tok.Type = token.INT
			return tok
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = string(l.ch)
		}
	}
	l.readChar()
	return tok
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
func newStringToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}

}
