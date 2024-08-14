package token

type Tokentype string

type Token struct {
	Type Tokentype
	Literal string
}

const (
	UNDEF = "UNDEF"
	EOF = "EOF" // "End  of File"の略 ファイルの終端

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
	COLON = ":"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//キーワード
	FUNCTION = "FUNCTION"
	LET = "LET"
)