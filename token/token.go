package token

// TokenType is the token type i.e integer, right bracket
type TokenType string

// TODO: try support row and column number
type Token struct {
	Type    TokenType
	Literal string // stores the actual token value
}

// Different token types below
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, x,y ...
	INT   = "INT"   // 123

	// Operators
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	BANG    = "!"
	LT      = "<"
	GT      = ">"
	SLASH   = "/"
	ASTERIX = "*"
	EQ      = "=="
	NOT_EQ  = "!="
	LT_EQ   = "<="
	GT_EQ   = ">="

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FALSE    = "FALSE"
	TRUE     = "TRUE"
)
