package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
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