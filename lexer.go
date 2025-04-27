package main

import "strings"

// Token представляє окремий токен
type Token struct {
	Type    string
	Literal string
}

// Lex приймає код і повертає масив токенів
func Lex(input string) []Token {
	words := strings.Fields(input)
	var tokens []Token

	for _, word := range words {
		switch word {
		case "print":
			tokens = append(tokens, Token{Type: "PRINT", Literal: word})
		default:
			tokens = append(tokens, Token{Type: "STRING", Literal: word})
		}
	}
	return tokens
}
