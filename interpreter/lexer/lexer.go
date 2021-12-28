package lexer

import (
	token "github.com/kanatti/just-go/interpreter/token"
)

type Lexer struct {
	input      string
	currentPos int  // Current position in input
	nextPos    int  // Next position to read in input
	ch         byte // current character under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) NextToken() *token.Token {
	l.ch = l.read()

	l.consumeWhiteSpace()

	var tokenType token.TokenType
	var literal = string(l.ch)

	switch l.ch {
	case '=':
		if l.peek() == '=' {
			l.consume()
			literal = "=="
			tokenType = token.EQUALS
		} else {
			tokenType = token.ASSIGN
		}
	case '+':
		tokenType = token.PLUS
	case ',':
		tokenType = token.COMMA
	case ';':
		tokenType = token.SEMICOLON
	case '(':
		tokenType = token.LPAREN
	case ')':
		tokenType = token.RPAREN
	case '{':
		tokenType = token.LBRACE
	case '}':
		tokenType = token.RBRACE
	case '!':
		if l.peek() == '=' {
			l.consume()
			literal = "!="
			tokenType = token.NOT_EQUALS
		} else {
			tokenType = token.NOT
		}
	case '-':
		tokenType = token.MINUS
	case '/':
		tokenType = token.SLASH
	case '*':
		tokenType = token.ASTERISK
	case '<':
		tokenType = token.LESS_THAN
	case '>':
		tokenType = token.GREATER_THAN
	case 0:
		tokenType = token.EOF
		literal = ""
	default:
		if isLetter(l.ch) {
			literal = l.readIdentifier()
			tokenType = token.LookupTokenTypeFromString(literal)
		} else if isNum(l.ch) {
			literal = l.readNum()
			tokenType = token.INT
		} else {
			tokenType = token.ILLEGAL
		}
	}

	return &token.Token{Type: tokenType, Literal: literal}
}

func (l *Lexer) read() byte {
	ch := l.peek()
	l.currentPos = l.nextPos
	l.nextPos += 1
	return ch
}

func (l *Lexer) consume() {
	l.currentPos = l.nextPos
	l.nextPos += 1
}

func (l *Lexer) peek() byte {
	if l.nextPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPos]
	}
}

func (l *Lexer) readIdentifier() string {
	startPos := l.currentPos
	for isLetterOrNum(l.peek()) {
		l.consume()
	}
	return l.input[startPos : l.currentPos+1]
}

func (l *Lexer) readNum() string {
	startPos := l.currentPos
	for isNum(l.peek()) {
		l.consume()
	}
	return l.input[startPos : l.currentPos+1]
}

func (l *Lexer) consumeWhiteSpace() {
	for isWhiteSpace(l.ch) {
		l.ch = l.read()
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isLetterOrNum(ch byte) bool {
	return isLetter(ch) || ('0' <= ch && ch <= '9')
}

func isNum(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch byte) bool {
	return ch == '\n' || ch == '\t' || ch == ' ' || ch == '\r'
}
