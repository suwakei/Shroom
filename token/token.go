package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	UNDEF = "UNDEF"
	EOF = "EOF" // "End  of File"の略 ファイルの終端

	//識別子
	IDENTIFIER = "IDENTIFIER"
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

// 予約語定義
var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

// 予約語かどうか判定
func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENTIFIER
}