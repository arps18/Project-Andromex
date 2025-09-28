package token

type TokenType string

type Token struct{
	Type TokenType // What kind of token (Number, Operator, etc)
	Literal string // The actual text from sourceCode
}

const(
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	IDENT = "IDENT" // add, foo, x
	INT = "INT" // 123
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"
	LPAREN = "("
	RPAREN = ")"
	SEMICOLON = ";"
	LET = "LET"
	PRINT = "PRINT"
)