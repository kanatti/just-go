package token

type TokenType int8

type Token struct {
	Type    TokenType
	Literal string
}

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENTIFIER"
	case INT:
		return "INTEGER"
	case ASSIGN:
		return "="
	case PLUS:
		return "+"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case FUNCTION:
		return "fn"
	case LET:
		return "let"
	default:
		return "UNKNOWN TOKEN"
	}
}

const (
	ILLEGAL TokenType = 0
	EOF     TokenType = 1

	// Identifiers
	IDENT TokenType = 2
	INT   TokenType = 3

	// Delimiters
	COMMA     TokenType = 4
	SEMICOLON TokenType = 5

	LPAREN TokenType = 6
	RPAREN TokenType = 7
	LBRACE TokenType = 8
	RBRACE TokenType = 9

	// Operators
	ASSIGN       TokenType = 20
	PLUS         TokenType = 21
	EQUALS       TokenType = 22
	NOT          TokenType = 23
	NOT_EQUALS   TokenType = 24
	MINUS        TokenType = 25
	SLASH        TokenType = 26
	ASTERISK     TokenType = 27
	LESS_THAN    TokenType = 28
	GREATER_THAN TokenType = 29

	// Keywords
	FUNCTION TokenType = 40
	LET      TokenType = 41
	TRUE     TokenType = 42
	FALSE    TokenType = 43
	IF       TokenType = 44
	ELSE     TokenType = 45
	RETURN   TokenType = 46
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupTokenTypeFromString(str string) TokenType {
	if tok, ok := keywords[str]; ok {
		return tok
	}
	return IDENT
}
