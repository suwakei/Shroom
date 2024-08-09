package token

type Tokentype string

type Token struct {
	Tokentype
	string
}

const (
	UNDEF = "UNDEF"
	EOF = "EOF"

	//識別子
	IDENT = "IDENT"
	INT = "INT"

	//演算子
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"

	//デリミタ
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//キーワード
	FUNCTION = "FUNCTION"
	LET = "LET"
)