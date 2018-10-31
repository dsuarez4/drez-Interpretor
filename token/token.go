package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x ,y
	INT = "INT" // 2, 3, 8564

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"

)

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string] TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// Why return this IDENT ?? HUH?!?!
	return IDENT
}
