package lexer

import (
	"strings"
)

type TokenType string

const (
	TokenPrintln    TokenType = "PRINTLN"
	TokenPrint      TokenType = "PRINT"
	TokenLet        TokenType = "LET"
	TokenFunc       TokenType = "FUNC"
	TokenIf         TokenType = "IF"
	TokenElse       TokenType = "ELSE"
	TokenWhile      TokenType = "WHILE"
	TokenEnd        TokenType = "END"
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenString     TokenType = "STRING"
	TokenNumber     TokenType = "NUMBER"
	TokenUnknown    TokenType = "UNKNOWN"
)

type Token struct {
	Type    TokenType
	Literal string
}

func Lex(input string) []Token {
	lines := strings.Split(input, "\n")
	tokens := []Token{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "println"):
			tokens = append(tokens, Token{Type: TokenPrintln, Literal: line})
		case strings.HasPrefix(line, "print"):
			tokens = append(tokens, Token{Type: TokenPrint, Literal: line})
		case strings.HasPrefix(line, "let"):
			tokens = append(tokens, Token{Type: TokenLet, Literal: line})
		case strings.HasPrefix(line, "func"):
			tokens = append(tokens, Token{Type: TokenFunc, Literal: line})
		case strings.HasPrefix(line, "if"):
			tokens = append(tokens, Token{Type: TokenIf, Literal: line})
		case strings.HasPrefix(line, "else"):
			tokens = append(tokens, Token{Type: TokenElse, Literal: line})
		case strings.HasPrefix(line, "while"):
			tokens = append(tokens, Token{Type: TokenWhile, Literal: line})
		case strings.HasPrefix(line, "end"):
			tokens = append(tokens, Token{Type: TokenEnd, Literal: line})
		default:
			tokens = append(tokens, Token{Type: TokenIdentifier, Literal: line})
		}
	}

	return tokens
}
