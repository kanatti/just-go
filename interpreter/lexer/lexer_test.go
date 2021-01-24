package lexer

import (
	"testing"

	token "github.com/kanatti/just-go/interpreter/token"
)

func TestNextTokenBasic(t *testing.T) {
	input := `++(+){}`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.PLUS, "+"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
	}
	lexer := New(input)

	for i, test := range tests {
		tok := lexer.nextToken()

		if tok.Type != test.expectedType {
			t.Errorf("tests[%d] - Wrong Token Type expected: %s got: %s", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Errorf("tests[%d] - Wrong Token Literal expected: %s got: %s", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenFull(t *testing.T) {
	input := `let num1 = 10;
let num2 = 20;

fn add(x, y) {
	x + y;
};

let result = add(num1, num2);
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "num1"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "num2"},
		{token.ASSIGN, "="},
		{token.INT, "20"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "num1"},
		{token.COMMA, ","},
		{token.IDENT, "num2"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}
	lexer := New(input)

	for i, test := range tests {
		tok := lexer.nextToken()

		if tok.Type != test.expectedType {
			t.Errorf("tests[%d] - Wrong Token Type expected: %s got: %s", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Errorf("tests[%d] - Wrong Token Literal expected: %s got: %s", i, test.expectedLiteral, tok.Literal)
		}
	}
}
