package translator

import (
	"fmt"

	"github.com/4am-png/Nexa-Next-Generation-Systems-Architecture/lexer"
)

func Translate(tokens []lexer.Token) {
	for _, token := range tokens {
		switch token.Type {
		case lexer.TokenPrintln:
			fmt.Println("Executing:", token.Literal)
		case lexer.TokenPrint:
			fmt.Print("Executing:", token.Literal)
		default:
			fmt.Println("Unknown execution for:", token.Literal)
		}
	}
}
