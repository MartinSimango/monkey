package lexer

import (
	"testing"

	"github.com/MartinSimango/monkey/token"
)

func TestLexer_NextToken(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedTokens []token.Token
	}{
		{
			name:  "Test subset of monkey characters",
			input: "=+(){},;",
			expectedTokens: []token.Token{
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},

		{
			name: "Test simple function definition",
			input: ` let five = 5;
				let ten = 10;

				let add = fn(x,y) {
				 x + y;
				}
				
				let result = add(five,ten);
			`,
			expectedTokens: []token.Token{
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)
			for i, et := range tt.expectedTokens {
				tok := l.NextToken()
				if tok.Type != et.Type {
					t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, et.Type, tok.Type)
				}

				if tok.Literal != et.Literal {
					t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q", i, et.Literal, tok.Literal)

				}
			}

		})

	}
}
