package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	UNDEF = "UNDEF"
	EOF   = "EOF" // "End  Of File"の略 ファイルの終端

	//識別子
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	STRING     = "STRING"

	//演算子
	ASSIGN    = "="
	EQUAL     = "=="
	NOT_EQUAL = "!="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTARISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"

	//デリミタ
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	RBRACKET  = "["
	LBRACKET  = "]"

	//キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELIF     = "ELIF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	RETURN   = "RETURN"
	PILEUS   = "PILEUS"
)

// 予約語定義
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"const":  CONST,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"elif":   ELIF,
	"else":   ELSE,
	"for":    FOR,
	"return": RETURN,
	"pileus": PILEUS,
}

// 予約語かどうか判定
func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENTIFIER
}
