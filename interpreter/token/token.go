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

	// Operators
	ASSIGN TokenType = 4
	PLUS   TokenType = 5

	// Delimiters
	COMMA     TokenType = 6
	SEMICOLON TokenType = 7

	LPAREN TokenType = 8
	RPAREN TokenType = 9
	LBRACE TokenType = 10
	RBRACE TokenType = 11

	// Keywords
	FUNCTION TokenType = 12
	LET      TokenType = 13
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupTokenType(str string) TokenType {
	if tok, ok := keywords[str]; ok {
		return tok
	}
	return IDENT
}
