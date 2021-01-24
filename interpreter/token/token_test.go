package token

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	tests := []struct {
		token TokenType
		want  string
	}{
		{ILLEGAL, "ILLEGAL"},
		{IDENT, "IDENTIFIER"},
		{INT, "INTEGER"},
		{ASSIGN, "="},
		{PLUS, "+"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{FUNCTION, "fn"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.want)
		t.Run(testName, func(t *testing.T) {
			ans := fmt.Sprintf("%v", tt.token)
			if ans != tt.want {
				t.Errorf("Got %s, Want %s", ans, tt.want)
			}
		})
	}
}
